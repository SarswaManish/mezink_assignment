package common

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server struct {
		Port uint64 `toml:"port"`
	} `toml:"server"`

	Cors struct {
		AllowedHeaders     []string `toml:"allowed_headers"`
		AllowedMethods     []string `toml:"allowed_methods"`
		AllowedCredentials bool     `toml:"allowed_credentials"`
		AllowedOrigin      []string `toml:"allowed_origin"`
	} `toml:"cors"`
	Database struct {
		Postgres struct {
			Host     string `toml:"pg_host"`
			Password string `toml:"pg_password"`
			User     string `toml:"pg_user"`
			Db       string `toml:"pg_db"`
		} `toml:"postgres"`
	} `toml:"database"`
	InternalAuth struct {
		Username string `toml:"auth_username"`
		Password string `toml:"auth_password"`
	} `toml:"internalAuth"`
}

func NewConfig(filepath string) Config {
	var conf Config

	if _, err := toml.DecodeFile(filepath, &conf); err != nil {
		panic(err)
	}
	return conf
}
