package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

var (
	once   = &sync.Once{}
	config = &ImmutableConfig{}
)

type (
	ServiceConfig struct {
		Name string `env:"name"`
	}

	ImmutableConfig struct {
		Service ServiceConfig `env:"service"`
	}

	IImmutableConfig interface {
		GetServiceName() string
	}
)

func (c *ImmutableConfig) GetServiceName() string {
	return c.Service.Name
}

func NewImmutableConfig() IImmutableConfig {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}

		path := "configs/config.local.yml"
		stage := os.Getenv("GO_STAGE") // Whether local, staging, or production

		if stage != "" {
			path = fmt.Sprintf("configs/config.%s.yml", stage)
		}

		if err := configor.New(&configor.Config{AutoReload: true, Debug: true, Verbose: true}).Load(config, path); err != nil {
			panic("cannot load configuration!")
		}
	})

	return config
}
