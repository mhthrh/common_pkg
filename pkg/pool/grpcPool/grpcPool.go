package grpcPool

import (
	"errors"
	"fmt"
	"github.com/mhthrh/common_pkg/pkg/model/grpc/pool"
	"google.golang.org/grpc"
	"sync"
)

var (
	poolMap = sync.Map{}
	poolKey = "pool"
)

type GrpcPool struct {
}

func NewPool(address string, size int) (*GrpcPool, error) {
	if size < 1 {
		return nil, errors.New("pool size is less than zero")
	}
	if len(address) == 0 {
		return nil, errors.New("address is empty")
	}

	conns := make([]*grpc.ClientConn, size)
	for i := range conns {
		conn, err := grpc.NewClient(address, grpc.WithInsecure())
		if err != nil {
			return nil, fmt.Errorf("dial failed: %v", err)
		}
		conns[i] = conn
	}
	result := pool.GRPC{
		Conns: conns,
		Mu:    sync.Mutex{},
		Next:  0,
	}

	poolMap.Store(poolKey, &result)

	return &GrpcPool{}, nil
}

func (g *GrpcPool) Get() (*grpc.ClientConn, error) {
	p1, ok := poolMap.Load(poolKey)
	if !ok {
		return nil, errors.New("pool is empty")
	}

	p := p1.(*pool.GRPC)

	p.Mu.Lock()
	defer p.Mu.Unlock()
	conn := p.Conns[p.Next]
	p.Next = (p.Next + 1) % len(p.Conns)
	return conn, nil
}

func (g *GrpcPool) Release() error {
	pp, _ := poolMap.Load(poolKey)
	p := pp.(*pool.GRPC)

	for _, conn := range p.Conns {
		_ = conn.Close()
	}

	return nil
}
