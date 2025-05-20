package config

import (
	. "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
)

type IConfig interface {
	Read() *xErrors.Error
	GetServer() (Server, *xErrors.Error)
	GetAdminUser() (AdminUser, *xErrors.Error)
	GetDbConfig() (PostgresConfig, *xErrors.Error)
	GetMongo() (Mongo, *xErrors.Error)
	GetSecrets() ([]Secret, *xErrors.Error)
	GetGrpcs() ([]Grpc, *xErrors.Error)
}
