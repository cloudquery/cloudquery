package azure

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

	//"github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2018-09-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
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
	Subscriptions []string `mapstructure:"subscriptions"`
	Resources     []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
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

func NewProvider(db *database.Database, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:  db,
		log: log,
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
	log.Info("Creating tables if needed")
	for _, tables := range TablesArr {
		err := db.AutoMigrate(tables...)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
}

func (p *Provider) Run(config interface{}) error {
	ctx := context.Background()
	err := mapstructure.Decode(config, &p.config)
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
