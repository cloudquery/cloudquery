package provider

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-azure/resources/compute"
	"github.com/cloudquery/cq-provider-azure/resources/keyvault"
	"github.com/cloudquery/cq-provider-azure/resources/mysql"
	"github.com/cloudquery/cq-provider-azure/resources/postgresql"
	"github.com/cloudquery/cq-provider-azure/resources/resources"
	"github.com/cloudquery/cq-provider-azure/resources/sql"
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v3"
)

type ResourceFunc func(subscriptionID string, auth autorest.Authorizer, db *database.Database, log hclog.Logger, gConfig interface{}) error

type Provider struct {
	region         string
	db             *database.Database
	config         Config
	subscriptionID string
	resourceFuncs  map[string]ResourceFunc
	Logger         hclog.Logger
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

	p.resourceFuncs = map[string]ResourceFunc{
		"resources.groups":   resources.Groups,
		"sql.servers":        sql.Servers,
		"sql.databases":      sql.Databases,
		"postgresql.servers": postgresql.Servers,
		"mysql.servers":      mysql.Servers,
		"compute.disks":      compute.Disks,
		"keyvault.vaults":    keyvault.Vaults,
	}
	p.Logger.Info("Creating tables if needed")
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
		p.Logger.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/azure/tables-reference")
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
		p.Logger.Info(fmt.Sprintf("No subscriptions specified going to use: %v", subscriptions))
	} else {
		subscriptions = p.config.Subscriptions
	}

	for _, subscriptionID := range subscriptions {
		logger := p.Logger.With("subscription_id", subscriptionID)
		for _, resource := range p.config.Resources {
			f := p.resourceFuncs[resource.Name]
			if f == nil {
				return fmt.Errorf("resource %s is not supported", resource.Name)
			}
			logger := logger.With("resource", resource.Name)
			err := f(subscriptionID, azureAuth, p.db, logger, resource.Other)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
