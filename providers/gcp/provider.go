package gcp

import (
	"fmt"
	"github.com/cloudquery/cloudquery/providers/gcp/compute"
	"github.com/cloudquery/cloudquery/providers/gcp/iam"
	"github.com/cloudquery/cloudquery/providers/gcp/resource"
	"github.com/cloudquery/cloudquery/providers/gcp/storage"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

type Provider struct {
	db              *gorm.DB
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

type NewResourceFunc func(db *gorm.DB, log *zap.Logger,
	projectID string, region string) (resource.ClientInterface, error)

var resourceFactory = map[string]NewResourceFunc{
	"compute": compute.NewClient,
	"iam":     iam.NewClient,
	"storage": storage.NewClient,
}

func NewProvider(db *gorm.DB, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:              db,
		resourceClients: map[string]resource.ClientInterface{},
		log:             log,
	}
	p.db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "gcp_",
	}
	return &p, nil
}

func (p *Provider) Run(config interface{}) error {
	err := mapstructure.Decode(config, &p.config)
	if err != nil {
		return err
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
	p.db.NamingStrategy = schema.NamingStrategy{
		TablePrefix: fmt.Sprintf("gcp_%s_", service),
	}
	return p.resourceClients[service].CollectResource(resourceName, config)
}
