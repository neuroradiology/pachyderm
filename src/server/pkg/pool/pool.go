package pool

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc"
	"k8s.io/kubernetes/pkg/api"
	kube "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/watch"
)

type connCount struct {
	cc    *grpc.ClientConn
	count int64
}

// Pool stores a pool of grpc connections to a k8s service, it's useful in
// places where you would otherwise need to keep recreating connections.
type Pool struct {
	conns          map[string]*connCount
	connsLock      sync.Mutex
	endpointsWatch watch.Interface
	opts           []grpc.DialOption
	done           chan struct{}
}

// NewPool creates a new connection pool with connections to pods in the
// given service.
func NewPool(kubeClient *kube.Client, namespace string, serviceName string, numWorkers int, opts ...grpc.DialOption) (*Pool, error) {
	endpointsInterface := kubeClient.Endpoints(namespace)

	watch, err := endpointsInterface.Watch(api.ListOptions{
		LabelSelector: labels.SelectorFromSet(
			map[string]string{"app": serviceName},
		),
		Watch: true,
	})
	if err != nil {
		return nil, err
	}

	pool := &Pool{
		endpointsWatch: watch,
		opts:           opts,
		done:           make(chan struct{}),
	}
	go pool.watchEndpoints()
	return pool, nil
}

func (p *Pool) watchEndpoints() {
	for {
		select {
		case event, ok := <-p.endpointsWatch.ResultChan():
			if !ok {
				return
			}
			endpoints := event.Object.(*api.Endpoints)
			p.updateAddresses(endpoints)
		case <-p.done:
			return
		}
	}
}

func (p *Pool) updateAddresses(endpoints *api.Endpoints) {
	addresses := make(map[string]*connCount)
	p.connsLock.Lock()
	defer p.connsLock.Unlock()
	for _, subset := range endpoints.Subsets {
		// According the k8s docs, the full set of endpoints is the cross
		// product of (addresses x ports).
		for _, address := range subset.Addresses {
			for _, port := range subset.Ports {
				addr := fmt.Sprintf("%s:%d", address.IP, port.Port)
				if cc := p.conns[addr]; cc != nil {
					addresses[addr] = cc
				} else {
					addresses[addr] = nil
				}
			}
		}
	}
	p.conns = addresses
}

func (p *Pool) Do(ctx context.Context, f func(cc *grpc.ClientConn) error) error {
	var conn *connCount
	if err := func() error {
		p.connsLock.Lock()
		defer p.connsLock.Unlock()
		for addr, mapConn := range p.conns {
			if mapConn == nil {
				cc, err := grpc.DialContext(ctx, addr, p.opts...)
				if err != nil {
					return err
				}
				conn = &connCount{cc: cc}
				p.conns[addr] = conn
				// We break because this conn has a count of 0 which we know
				// we're not beating
				break
			} else {
				if conn == nil || mapConn.count < conn.count {
					conn = mapConn
				}
			}
		}
		atomic.AddInt64(&conn.count, 1)
		return nil
	}(); err != nil {
		return err
	}
	defer atomic.AddInt64(&conn.count, -1)
	return f(conn.cc)
}

// Close closes all connections stored in the pool, it returns an error if any
// of the calls to Close error.
func (p *Pool) Close() error {
	close(p.done)
	var retErr error
	for _, conn := range p.conns {
		if err := conn.cc.Close(); err != nil {
			retErr = err
		}
	}
	return retErr
}
