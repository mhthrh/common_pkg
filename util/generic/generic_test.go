package generic_test

import (
	"github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/util/generic"
	"testing"
)

func TestFilter(t *testing.T) {
	grpcs := []config.Grpc{
		{
			Srv:  "user",
			Host: "192.168.21.1",
			Port: 1501,
		},
		{
			Srv:  "notification",
			Host: "192.168.21.2",
			Port: 1502,
		},
	}
	secrets := []config.Secret{
		{
			Name:      "secret-1",
			SecretKey: "P@$$w0rd@",
		},
		{
			Name:      "secret-2",
			SecretKey: "nimdA@",
		},
	}

	item1 := generic.Filter(grpcs, "user", func(grpc config.Grpc, item string) bool {
		if grpc.Srv == item {
			return true
		}
		return false
	})
	if item1.Host != "192.168.21.1" {
		t.Error("got a wrong value")
	}

	item2 := generic.Filter(secrets, "secret-2", func(grpc config.Secret, item string) bool {
		if grpc.Name == item {
			return true
		}
		return false
	})
	if item2.SecretKey != "nimdA@" {
		t.Error("got a wrong value")
	}
}
