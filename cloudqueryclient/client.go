package cloudqueryclient

import (
	"fmt"
	"io/ioutil"

	"github.com/cloudquery/cloudquery/providers/aws"
	"github.com/cloudquery/cloudquery/providers/provider"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var ProviderMap = map[string]func(*gorm.DB, *zap.Logger) (provider.Interface, error){
	"aws": aws.NewProvider,
	//"azure": aws.NewProvider,
}

type Config struct {
	Providers []struct {
		Name string
		Rest map[string]interface{} `yaml:",inline"`
	}
}

type Client struct {
	db     *gorm.DB
	config Config
	log    *zap.Logger
}

func New(driver string, dsn string) (*Client, error) {
	client := Client{}
	gormLogger := logger.Default.LogMode(logger.Error)
	var err error = nil
	switch driver {
	case "sqlite":
		client.db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
		client.db.Exec("PRAGMA foreign_keys = ON")
	case "postgresql":
		client.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	case "mysql":
		client.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	case "sqlserver":
		client.db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
	default:
		return nil, fmt.Errorf("database driver only supports one of sqlite,postgresql,mysql,sqlserver")
	}
	if err != nil {
		return nil, err
	}
	client.log, err = zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *Client) Run(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	for _, provider := range config.Providers {
		if provider.Name == "" {
			return fmt.Errorf("provider must contain key: name")
		}
		if ProviderMap[provider.Name] == nil {
			return fmt.Errorf("provider %s is not supported\n", provider.Name)
		}
		p, err := ProviderMap[provider.Name](c.db, c.log)
		if err != nil {
			return err
		}
		err = p.Run(provider.Rest)
		if err != nil {
			return err
		}

	}

	return nil
}
