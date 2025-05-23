package config

import (
	"fmt"
	"time"
)

type Config struct {
	AppName  string         `json:"appName"`
	IsTest   bool           `json:"isTest"`
	Version  string         `json:"version"`
	ExpireAt string         `json:"expireDate"`
	Secrets  []Secret       `json:"secrets"`
	Postgres PostgresConfig `json:"postgresql"`
	Mongo    Mongo          `json:"mongo"`
	Admin    AdminUser      `json:"admin"`
	Host     Server         `json:"server"`
	GRPCs    []Grpc         `json:"grpc"`
}

type Server struct {
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	ReadTimeOut  time.Duration `json:"readTimeOut"`
	WriteTimeOut time.Duration `json:"writeTimeOut"`
	IdleTimeOut  time.Duration `json:"idleTimeOut"`
}

type AdminUser struct {
	UserName string `json:"user"`
	Password string `json:"pass"`
}
type PostgresConfig struct {
	Host           string  `yaml:"host" json:"host"`
	Port           int     `yaml:"port" json:"port"`
	UserName       string  `yaml:"username" json:"username"`
	Password       string  `yaml:"password" json:"password"`
	SSLModeEnabled bool    `yaml:"sslEnabled" json:"sslEnabled"`
	DatabaseName   string  `yaml:"database" json:"database"`
	Schema         string  `yaml:"schema" json:"schema"`
	SSLMode        SSLMode `yaml:"sslmode" json:"sslmode"`
}
type Mongo struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
type Secret struct {
	Name      string `json:"name"`
	SecretKey string `json:"secretKey"`
}
type Grpc struct {
	Srv   string `json:"srv"`
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Count int    `json:"poolSize"`
}

type SSLMode string

const (
	_          SSLMode = ""
	Disabled           = "disable"
	Require            = "require"
	VerifyCA           = "verify-ca"
	VerifyFull         = "verify-full"
)

func IsValid(v interface{}) (SSLMode, error) {
	var s SSLMode
	switch v {
	case Disabled:
		s = Disabled
	case Require:
		s = Require
	case VerifyCA:
		s = VerifyCA
	case VerifyFull:
		s = VerifyFull
	default:
		return "", fmt.Errorf("invalid ssl mode type")
	}
	return s, nil
}
