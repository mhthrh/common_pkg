package pool

import (
	"google.golang.org/grpc"
	"sync"
)

type GRPC struct {
	Conns []*grpc.ClientConn
	Mu    sync.Mutex
	Next  int
}

type IGrpcPool interface {
	Get(key string) (*grpc.ClientConn, error)
	Release(keys []string) error
}
