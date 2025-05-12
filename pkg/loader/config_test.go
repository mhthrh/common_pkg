package config_test

import (
	"github.com/mhthrh/common_pkg/pkg/loader"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	"log"
	"testing"
)

const (
	path   = "src/common_pkg/config/file"
	secret = "kiripiri"
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
