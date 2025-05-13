package config

import (
	"encoding/json"
	errs "errors"
	"fmt"
	. "github.com/mhthrh/common_pkg/pkg/model/config"
	"github.com/mhthrh/common_pkg/pkg/xErrors"
	cryptx "github.com/mhthrh/common_pkg/util/cryptox"
	env "github.com/mhthrh/common_pkg/util/environment"
	"github.com/mhthrh/common_pkg/util/file/text"
	"log"
)

const (
	environment = "environment"
	file        = "config.%s"
)

var (
	config      *Config
	isEncrypted = false
)

type Local struct {
	path   string
	secret string
}

type Remote struct {
	url  string
	user string
	pass string
}

func New(url, path, user, pass, secret string, enc bool) (IConfig, *xErrors.Error) {
	isEncrypted = enc
	if e := env.GetEnv(environment, ""); e != "remote" {
		return Local{
			path:   path,
			secret: secret,
		}, nil
	}
	return Remote{
		url:  url,
		user: user,
		pass: pass,
	}, nil

}

func (l Local) Read() *xErrors.Error {
	if config != nil {
		return nil
	}
	if !isEncrypted {
		txt := text.New(l.path, fmt.Sprintf(file, "json"), false)
		byts, err := txt.Read()
		if err != nil {
			log.Fatalf("read file failed: %v", err)
		}
		err = json.Unmarshal(byts, &config)
		if err != nil {
			log.Fatalf("json convert to struct has been failed: %v", err)
		}
	}

	c, err := cryptx.New(l.secret)
	if err != nil {
		log.Fatalf("crypto failed: %v", err)
	}
	f := text.New(l.path, fmt.Sprintf(file, "enc"), false)
	byts, err := f.Read()
	if err != nil {
		log.Fatalf("read file failed: %v", err)
	}
	content, err := c.Decrypt(string(byts))
	if err != nil {
		log.Fatalf("decription failed: %v", err)
	}

	err = json.Unmarshal([]byte(content), &config)
	if err != nil {
		log.Fatalf("json convert to struct has been failed: %v", err)
	}
	return nil
}

func (l Local) GetServer() (Server, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (l Local) GetAdminUser() (AdminUser, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (l Local) GetDbConfig() (PostgresConfig, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (l Local) GetMongo() (Mongo, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (l Local) GetSecrets() ([]Secret, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}
func (l Local) DbConfig() (PostgresConfig, *xErrors.Error) {
	config := PostgresConfig{}

	if config.Host == "" {
		return PostgresConfig{}, xErrors.FailedResource(errs.New("host is required"), nil)
	}
	if config.Port == 0 {
		return PostgresConfig{}, xErrors.FailedResource(errs.New("port is required"), nil)
	}
	if config.UserName == "" {
		return PostgresConfig{}, xErrors.FailedResource(errs.New("username is required"), nil)
	}
	if config.Password == "" {
		return PostgresConfig{}, xErrors.FailedResource(errs.New("password is required"), nil)
	}

	return config, nil
}

func (l Local) GetRootAdmin() (AdminUser, *xErrors.Error) {
	admin := AdminUser{}
	if admin.UserName == "" {
		return AdminUser{}, xErrors.FailedResource(errs.New("username is required"), nil)
	}
	if admin.Password == "" {
		return AdminUser{}, xErrors.FailedResource(errs.New("password is required"), nil)
	}
	return admin, nil
}

func (r Remote) Read() *xErrors.Error {
	//TODO implement me
	panic("implement me")
}

func (r Remote) GetServer() (Server, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (r Remote) GetAdminUser() (AdminUser, *xErrors.Error) {
	return AdminUser{}, xErrors.NotImplemented("service")

}

func (r Remote) GetDbConfig() (PostgresConfig, *xErrors.Error) {
	return PostgresConfig{}, xErrors.NotImplemented("service")
}

func (r Remote) GetMongo() (Mongo, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}

func (r Remote) GetSecrets() ([]Secret, *xErrors.Error) {
	//TODO implement me
	panic("implement me")
}
