package postgres

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/zulfikar4568/go-todo/internal/config"
	"github.com/zulfikar4568/go-todo/pkg/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Driver struct {
		immutableConfig config.IImmutableConfig
	}

	IDriverPostgres interface {
		Start() (map[string]*gorm.DB, error)
		GetInstance(name string) (*gorm.DB, error)
	}
)

var (
	once     = &sync.Once{}
	logger   = log.NewAppLogger("[driver][postgres]")
	instance = make(map[string]*gorm.DB)
)

func (d *Driver) GetDatabaseURLs() map[string]string {
	URLs := make(map[string]string)

	for _, db := range config.NewImmutableConfig().GetDatabaseConfig() {
		if db.Driver == "postgres" {
			URLs[db.ID] = db.URL
		}
	}
	return URLs
}

func (d *Driver) GetGormSources() map[string]string {
	result := make(map[string]string)

	for key, val := range d.GetDatabaseURLs() {
		obj, _ := url.Parse(val)

		hostWithPort := strings.Split(obj.Host, ":")
		password, _ := obj.User.Password()
		baseURLPath := strings.Split(obj.Path, "/")

		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			hostWithPort[0], obj.User.Username(), password, baseURLPath[1], obj.Port(), obj.Query().Get("sslmode"),
		)

		result[key] = dsn
	}

	return result
}

func (d *Driver) Start() (map[string]*gorm.DB, error) {
	var e error = nil

	once.Do(func() {
		for key, dsn := range d.GetGormSources() {
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

			if err != nil {
				logger.Error(fmt.Sprintf("failed to open connection to postgres with id [%s]", key), err)
				e = err

				break
			}

			sql, _ := db.DB()
			sql.SetMaxIdleConns(16)
			sql.SetMaxOpenConns(100)
			sql.SetConnMaxLifetime(30 * time.Second)

			instance[key] = db
		}
	})

	return instance, e
}

func (d *Driver) GetInstance(name string) (*gorm.DB, error) {
	if instance[name] == nil {
		return nil, fmt.Errorf("db connection with id [%s] is undefined", name)
	}
	return instance[name], nil
}

func NewDriver(cfg config.IImmutableConfig) IDriverPostgres {
	return &Driver{
		cfg,
	}
}
