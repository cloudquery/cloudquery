package main

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/cloudquery/cloudquery/sdk"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
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
	Domain    string
	Resources []struct {
		Name  string
		Other map[string]interface{}
	}
}

var tablesArr = [][]interface{}{
	applicationTables,
	userTables,
}

func (p *Provider)Init(driver string, dsn string, verbose bool) error {
	var err error
	p.db, err = database.Open(driver, dsn)
	if err != nil {
		return err
	}
	zapLogger, err := sdk.NewLogger(verbose)
	p.log = zapLogger

	p.log.Info("Creating tables if needed")
	for _, tables := range tablesArr {
		err := p.db.AutoMigrate(tables...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Provider) GenConfig() (string, error) {
	return configYaml, nil
}

func (p *Provider) Fetch(data []byte) error {
	err := yaml.Unmarshal(data, &p.config)
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		p.log.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/okta/tables-reference")
		return nil
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


func main() {
	sdk.ServePlugin(&Provider{})
}