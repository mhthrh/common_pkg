package config_test

import (
	"github.com/mhthrh/common_pkg/pkg/loader"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"log"
	"testing"
)

const (
	path   = "src/common_pkg/config/file"
	secret = "AnKoloft@~delNazok!12345"
)

var (
	cnfg config.IConfig
	err  *xErrors.Error
)

type test struct {
	name   string
	input  any
	output any
	hasErr bool
}

func init() {
	cnfg, err = config.New("", path, "", "", secret, true)
	if err != nil {
		log.Fatalf("failed to initialize configuration")
	}

}

func TestConfig_init(t *testing.T) {
	for _, v := range []bool{true, false} {
		cnfg, err = config.New("", path, "", "", secret, v)
		if err != nil {
			t.Errorf("failed to initialize configuration")
		}
		if err = cnfg.Read(); err != nil {
			t.Errorf("failed to initialize configuration")
		}
	}
}

func TestConfig_GetServer(t *testing.T) {
	for _, v := range []bool{true, false} {
		cnfg, err = config.New("", path, "", "", secret, v)
		if err != nil {
			t.Errorf("failed to initialize configuration")
		}
		if err = cnfg.Read(); err != nil {
			t.Errorf("failed to initialize configuration")
		}
		srv, err := cnfg.GetServer()
		if err != nil {
			t.Errorf("failed to initialize configuration, %v", err)
		}
		if srv.Host != "0.0.0.0" {
			t.Errorf("failed to initialize host")
		}
	}
}

func TestConfig_GetDB(t *testing.T) {
	for _, v := range []bool{true, false} {
		cnfg, err = config.New("", path, "", "", secret, v)
		if err != nil {
			t.Errorf("failed to initialize configuration")
		}
		if err = cnfg.Read(); err != nil {
			t.Errorf("failed to initialize configuration")
		}
		srv, err := cnfg.GetDbConfig()
		if err != nil {
			t.Errorf("failed to initialize configuration, %v", err)
		}
		if srv.Host != "localhost" {
			t.Errorf("failed to initialize host")
		}
	}
}

func TestConfig_GetMongo(t *testing.T) {
	for _, v := range []bool{true, false} {
		cnfg, err = config.New("", path, "", "", secret, v)
		if err != nil {
			t.Errorf("failed to initialize configuration")
		}
		if err = cnfg.Read(); err != nil {
			t.Errorf("failed to initialize configuration")
		}
		srv, err := cnfg.GetMongo()
		if err != nil {
			t.Errorf("failed to initialize configuration, %v", err)
		}
		if srv.Host != "localhost" {
			t.Errorf("failed to initialize host")
		}
	}
}
