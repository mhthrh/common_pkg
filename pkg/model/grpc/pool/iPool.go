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
	Get() (*grpc.ClientConn, error)
	Release() error
}
