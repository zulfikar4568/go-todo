package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
	"github.com/zulfikar4568/go-todo/pkg/log"
)

var (
	once   = &sync.Once{}
	config = &ImmutableConfig{}
	logger = log.NewAppLogger("[config]")
)

type (
	ServiceConfig struct {
		Name string `env:"name"`
	}

	JwtConfig struct {
		Secret        string `env:"secret"`
		SecretDefault string `default:"secret_default"`
	}

	HTTPConfig struct {
		Port uint16 `env:"port"`
	}

	DatabaseConfig struct {
		ID     string `env:"id"`
		Driver string `env:"driver"`
		URL    string `env:"url"`
	}

	ImmutableConfig struct {
		Jwt       JwtConfig        `env:"jwt"`
		HTTP      HTTPConfig       `env:"http"`
		Databases []DatabaseConfig `env:"databases"`
		Service   ServiceConfig    `env:"service"`
	}

	IImmutableConfig interface {
		GetServiceName() string
		GetServiceHTTPPort() uint16
		GetDatabaseConfig() []DatabaseConfig
		GetJwtSecret() string
	}
)

func (c *ImmutableConfig) GetServiceName() string {
	return c.Service.Name
}

func (c *ImmutableConfig) GetServiceHTTPPort() uint16 {
	return c.HTTP.Port
}

func (c *ImmutableConfig) GetDatabaseConfig() []DatabaseConfig {
	return c.Databases
}

func (c *ImmutableConfig) GetJwtSecret() string {
	return c.Jwt.Secret
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

		if err := configor.New(&configor.Config{AutoReload: true, Debug: false, Verbose: false}).Load(config, path); err != nil {
			logger.Error("error while load configuration!", err)
			panic("cannot load configuration!")
		}

		// Config Secret of JWT
		jwtSecret := os.Getenv(config.Jwt.Secret)

		if jwtSecret == "" {
			config.Jwt.Secret = config.Jwt.SecretDefault
		} else {
			config.Jwt.Secret = jwtSecret
		}

		// Config Database
		for i, db := range config.Databases {
			id := os.Getenv(db.ID)
			url := os.Getenv(db.URL)

			if id == "" || url == "" {
				logger.Error("cannot define config db")
				panic(fmt.Sprintf("cannot define config db at index[%d]", i))
			}

			config.Databases[i].ID = id
			config.Databases[i].URL = url
		}
	})

	return config
}
