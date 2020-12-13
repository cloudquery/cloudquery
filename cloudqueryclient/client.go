package cloudqueryclient

import (
	"database/sql"
	"fmt"
	"github.com/cloudquery/cloudquery/providers/aws"
	"github.com/cloudquery/cloudquery/providers/gcp"
	"github.com/cloudquery/cloudquery/providers/okta"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"os"
)

var ProviderMap = map[string]func(*gorm.DB, *zap.Logger) (provider.Interface, error){
	"aws":  aws.NewProvider,
	"gcp":  gcp.NewProvider,
	"okta": okta.NewProvider,
	//"azure": aws.NewProvider,
}

type Config struct {
	Providers []struct {
		Name string
		Rest map[string]interface{} `yaml:",inline"`
	}
}

type PolicyConfig struct {
	Queries []struct {
		Name  string
		Query string
	}
}

type Client struct {
	db     *gorm.DB
	config Config
	log    *zap.Logger
}

func NewLogger(verbose bool, options ...zap.Option) (*zap.Logger, error) {
	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	disableCaller := true
	if verbose {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
		disableCaller = false
	}
	return zap.Config{
		Sampling:         nil,
		Level:            level,
		Development:      true,
		DisableCaller:    disableCaller,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
}

func New(driver string, dsn string, verbose bool) (*Client, error) {
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

	zapLogger, err := NewLogger(verbose)
	client.log = zapLogger
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (c *Client) Run(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s doesn't exist. you can create one via 'gen config' command", path)
		}
		return err
	}

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
		log := c.log.With(zap.String("provider", provider.Name))
		p, err := ProviderMap[provider.Name](c.db, log)
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

func (c *Client) RunQuery(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s doesn't exist. you can create one via 'gen policy' command", path)
		}
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	config := PolicyConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	c.log.Info("Executing queries", zap.Int("count", len(config.Queries)))
	for _, query := range config.Queries {
		c.log.Info("Executing query", zap.String("name", query.Name))
		//var res string
		rows, err := c.db.Raw(query.Query).Rows()
		if err != nil {
			return err
		}
		columns, err := rows.Columns()
		if err != nil {
			return err
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoFormatHeaders(false)
		table.SetHeader(columns)
		nc := len(columns)
		prettyRow := make([]string, nc)
		res := make([]sql.NullString, nc)
		resPtrs := make([]interface{}, nc)
		for i := 0; i < nc; i++ {
			resPtrs[i] = &res[i]
		}
		resultsCount := 0
		for rows.Next() {
			err := rows.Scan(resPtrs...)
			if err != nil {
				return err
			}
			resultsCount += 1
			for i, v := range res {
				prettyRow[i] = v.String
			}
			table.Append(prettyRow)
		}
		if resultsCount > 0 {
			c.log.Info("Check failed. Query returned results.", zap.String("name", query.Name), zap.Int("count", resultsCount))
			table.Render()
		} else {
			c.log.Info("Check passed. Query returned no results.", zap.String("name", query.Name))
		}

	}
	return nil
}
