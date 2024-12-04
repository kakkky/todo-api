package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server server
	MySQL  mysql
	Redis  redis
}

type server struct {
	Address string `envconfig:"ADDRESS"`
	Port    string `envconfig:"PORT"`
}

type mysql struct {
	Name     string `envconfig:"TODO_DB_NAME"`
	User     string `envconfig:"TODO_DB_USER"`
	Password string `envconfig:"TODO_DB_PASSWORD"`
	Port     string `envconfig:"TODO_DB_PORT"`
	Host     string `envconfig:"TODO_DB_HOST"`
}

type redis struct {
	Host string `envconfig:"KVS_HOST"`
	Port string `envconfig:"KVS_PORT"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
