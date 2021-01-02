package okta

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

type Provider struct {
	db               *database.Database
	config           Config
	resourceClients  map[string]common.ClientInterface
	log              *zap.Logger
	client           *okta.Client
	resourceMigrated map[string]bool
}

type Config struct {
	Domain    string `mapstructure:"domain"`
	Resources []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

type NewResourceFunc func(client *okta.Client, db *gorm.DB, log *zap.Logger) (common.ClientInterface, error)

var tablesArr = [][]interface{}{
	applicationTables,
	userTables,
}

func NewProvider(db *database.Database, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:               db,
		resourceClients:  map[string]common.ClientInterface{},
		resourceMigrated: map[string]bool{},
		log:              log,
	}
	log.Info("Creating tables if needed")
	for _, tables := range tablesArr {
		err := p.db.AutoMigrate(tables...)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
}

func (p *Provider) Run(config interface{}) error {
	err := mapstructure.Decode(config, &p.config)
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		return fmt.Errorf("please specify at least 1 resource in config.yml. see: https://docs.cloudquery.io/okta/tables-reference")
	}

	oktaToken, ok := os.LookupEnv("OKTA_TOKEN")
	if !ok {
		return fmt.Errorf("please set OKTA_TOKEN")
	}

	if p.config.Domain == "" {
		return fmt.Errorf("please set your okta \"domain\" in config.yml")
	}

	_, p.client, err = okta.NewClient(context.Background(), okta.WithOrgUrl(p.config.Domain), okta.WithToken(oktaToken))
	if err != nil {
		return err
	}

	for _, resource := range p.config.Resources {
		switch resource.Name {
		case "users":
			return p.users(resource.Other)
		case "applications":
			return p.applications(resource.Other)
		default:
			return fmt.Errorf("unsupported resource %s", resource)
		}
	}

	return nil
}
