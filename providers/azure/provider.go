package main

import (
	"context"
	"fmt"
	"github.com/Azure/go-autorest/autorest"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/azure/compute"
	"github.com/cloudquery/cloudquery/providers/azure/keyvault"
	"github.com/cloudquery/cloudquery/providers/azure/mysql"
	"github.com/cloudquery/cloudquery/providers/azure/postgresql"
	"github.com/cloudquery/cloudquery/providers/azure/resources"
	"github.com/cloudquery/cloudquery/providers/azure/sql"
	"github.com/cloudquery/cloudquery/sdk"
	"gopkg.in/yaml.v3"

	//"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"go.uber.org/zap"
)

type ResourceFunc func(subscriptionID string, auth autorest.Authorizer, db *database.Database, log *zap.Logger, gConfig interface{}) error

type Provider struct {
	region         string
	db             *database.Database
	config         Config
	subscriptionID string
	resourceFuncs  map[string]ResourceFunc
	log            *zap.Logger
}

type Config struct {
	Subscriptions []string
	Resources     []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

var TablesArr = [][]interface{}{
	resources.GroupTables,
	sql.ServerTables,
	sql.DatabaseTables,
	postgresql.ServerTables,
	mysql.ServerTables,
	compute.DiskTables,
	keyvault.VaultTables,
}

func (p *Provider)Init(driver string, dsn string, verbose bool) error {
	var err error
	p.db, err = database.Open(driver, dsn)
	if err != nil {
		return err
	}

	zapLogger, err := sdk.NewLogger(verbose)
	p.log = zapLogger
	p.resourceFuncs = map[string]ResourceFunc{
		"resources.groups":   resources.Groups,
		"sql.servers":        sql.Servers,
		"sql.databases":      sql.Databases,
		"postgresql.servers": postgresql.Servers,
		"mysql.servers":      mysql.Servers,
		"compute.disks":      compute.Disks,
		"keyvault.vaults":    keyvault.Vaults,
	}
	p.log.Info("Creating tables if needed")
	for _, tables := range TablesArr {
		err := p.db.AutoMigrate(tables...)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (p *Provider) GenConfig() (string, error) {
	return configYaml, nil
}

func (p *Provider) Fetch(data []byte) error {
	ctx := context.Background()

	err := yaml.Unmarshal(data, &p.config)
	if err != nil {
		return err
	}

	if len(p.config.Resources) == 0 {
		p.log.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/azure/tables-reference")
		return nil
	}

	azureAuth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return err
	}
	var subscriptions []string
	if len(p.config.Subscriptions) == 0 {
		client := subscription.NewSubscriptionsClient()
		client.Authorizer = azureAuth
		res, err := client.List(ctx)
		if err != nil {
			return err
		}
		for res.NotDone() {
			for _, sub := range res.Values() {
				subscriptions = append(subscriptions, *sub.SubscriptionID)
			}
			err := res.NextWithContext(ctx)
			if err != nil {
				return err
			}
		}
		p.log.Info(fmt.Sprintf("No subscriptions specified going to use: %v", subscriptions))
	} else {
		subscriptions = p.config.Subscriptions
	}

	for _, subscriptionID := range subscriptions {
		logger := p.log.With(zap.String("subscription_id", subscriptionID))
		//var wg sync.WaitGroup
		for _, resource := range p.config.Resources {
			f := p.resourceFuncs[resource.Name]
			if f == nil {
				return fmt.Errorf("resource %s is not supported", resource.Name)
			}
			logger := logger.With(zap.String("resource", resource.Name))
			//wg.Add(1)
			err := f(subscriptionID, azureAuth, p.db, logger, resource.Other)
			if err != nil {
				return err
			}
		}
		//wg.Wait()
	}

	return nil
}

func main() {
	sdk.ServePlugin(&Provider{})
}