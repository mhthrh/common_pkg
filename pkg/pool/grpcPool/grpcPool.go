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

func NewPool(address string, size int, keys []string) (*GrpcPool, error) {
	if len(keys) == 0 {
		return nil, errors.New("key is empty")
	}
	if size < 1 {
		return nil, errors.New("pool size is less than zero")
	}
	if len(address) == 0 {
		return nil, errors.New("address is empty")
	}
	result := make(map[string]*pool.GRPC)
	for _, v := range keys {
		conns := make([]*grpc.ClientConn, size)
		for i := range conns {
			conn, err := grpc.NewClient(address, grpc.WithInsecure())
			if err != nil {
				return nil, fmt.Errorf("dial failed: %v", err)
			}
			conns[i] = conn
		}
		result[v].Conns = conns
		result[v].Mu = sync.Mutex{}
		result[v].Next = 0

	}
	poolMap.Store(poolKey, &result)

	return &GrpcPool{}, nil
}

func (g *GrpcPool) Get(key string) (*grpc.ClientConn, error) {
	p1, ok := poolMap.Load(poolKey)
	if !ok {
		return nil, errors.New("pool is empty")
	}

	p := p1.(*map[string]*pool.GRPC)

	cPool, ok := (*p)[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s not exist", key))
	}

	cPool.Mu.Lock()
	defer cPool.Mu.Unlock()
	conn := cPool.Conns[cPool.Next]
	cPool.Next = (cPool.Next + 1) % len(cPool.Conns)
	return conn, nil
}

func (g *GrpcPool) Release(keys []string) error {
	pp, _ := poolMap.Load(poolKey)
	p := pp.(map[string]*pool.GRPC)
	for _, v := range keys {
		for _, conn := range p[v].Conns {
			_ = conn.Close()
		}
	}
	return nil
}
