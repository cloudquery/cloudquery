package gcp

import (
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/gcp/compute"
	"github.com/cloudquery/cloudquery/providers/gcp/iam"
	"github.com/cloudquery/cloudquery/providers/gcp/resource"
	"github.com/cloudquery/cloudquery/providers/gcp/storage"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"strings"
)

type Provider struct {
	db              *database.Database
	config          Config
	resourceClients map[string]resource.ClientInterface
	log             *zap.Logger
}

type Config struct {
	Region    string `mapstructure:"region"`
	ProjectID string `mapstructure:"project_id"`
	Resources []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

type NewResourceFunc func(db *database.Database, log *zap.Logger,
	projectID string, region string) (resource.ClientInterface, error)

var resourceFactory = map[string]NewResourceFunc{
	"compute": compute.NewClient,
	"iam":     iam.NewClient,
	"storage": storage.NewClient,
}

var tablesArr = [][]interface{}{
	compute.AddressTables,
	compute.AutoscalerTables,
	compute.DiskTypeTables,
	compute.ImageTables,
	compute.InstanceTables,
	compute.InterconnectTables,
	compute.SSLCertificateTables,
	compute.VPNGatewayTables,
	iam.RoleTables,
	iam.ServiceAccountTables,
	storage.BucketTables,
}

func NewProvider(db *database.Database, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:              db,
		resourceClients: map[string]resource.ClientInterface{},
		log:             log,
	}

	for _, table := range tablesArr {
		err := db.AutoMigrate(table...)
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
		return fmt.Errorf("please specify at least 1 resource in config.yml. see: https://docs.cloudquery.io/gcp/tables-reference")
	}

	for _, resource := range p.config.Resources {
		err := p.collectResource(resource.Name, resource.Other)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Provider) resetClients() {
	p.resourceClients = map[string]resource.ClientInterface{}
}

func (p *Provider) collectResource(fullResourceName string, config interface{}) error {
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		return fmt.Errorf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if resourceFactory[service] == nil {
		return fmt.Errorf("unsupported service %s", service)
	}

	if p.resourceClients[service] == nil {
		var err error
		p.resourceClients[service], err = resourceFactory[service](p.db, p.log, p.config.ProjectID, p.config.Region)
		if err != nil {
			return err
		}
	}
	return p.resourceClients[service].CollectResource(resourceName, config)
}
