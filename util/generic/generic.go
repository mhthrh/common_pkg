package generic

import (
	"github.com/mhthrh/common_pkg/pkg/model/config"
)

func Filter[T config.Secret | config.Grpc](slice []T, i string, predicate func(T, string) bool) T {
	var result T
	for _, v := range slice {
		if predicate(v, i) {
			return v
		}
	}
	return result
}
