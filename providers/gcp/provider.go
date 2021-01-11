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
	"google.golang.org/api/googleapi"
	"log"
	"strings"
	"sync"
)

type Provider struct {
	db              *database.Database
	config          Config
	resourceClients map[string]resource.ClientInterface
	log             *zap.Logger
}

type Config struct {
	ProjectIDs [] string `mapstructure:"project_ids"`
	Resources []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

type NewResourceFunc func(db *database.Database, log *zap.Logger,
	projectID string) (resource.ClientInterface, error)

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

	if len(p.config.ProjectIDs) == 0 {
		return fmt.Errorf("please specify at least 1 project_id in config.yml. see: https://docs.cloudquery.io/gcp/tables-reference")
	}

	for _, projectID := range p.config.ProjectIDs {
		if projectID == "<CHANGE_THIS_TO_YOUR_PROJECT_ID>" {
			return fmt.Errorf("please specify a valid project_id in config.yml instead of <CHANGE_THIS_TO_YOUR_PROJECT_ID>")
		}
		err := p.initClients(projectID)
		if err != nil {
			return err
		}
		var wg sync.WaitGroup
		for _, resource := range p.config.Resources {
			wg.Add(1)
			go p.collectResource(&wg, projectID, resource.Name, resource.Other)

		}
		wg.Wait()
	}

	return nil
}

func (p *Provider) initClients(projectID string) error{
	var err error
	for serviceName, newFunc := range resourceFactory {
		zapLog := p.log.With(zap.String("project_id", projectID))
		p.resourceClients[serviceName], err = newFunc(p.db, zapLog, projectID)
		if err != nil {
			return err
		}
	}
	return nil
}


func (p *Provider) collectResource(wg *sync.WaitGroup, projectID string, fullResourceName string, config interface{}) {
	defer wg.Done()
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		log.Fatalf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if resourceFactory[service] == nil {
		log.Fatalf("unsupported service %s", service)
	}

	err := p.resourceClients[service].CollectResource(resourceName, config)
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == 403 && len(e.Errors) > 0 && e.Errors[0].Reason == "accessNotConfigured" {
				p.log.Info("access not configured. skipping.",
					zap.String("project_id", projectID), zap.String("resource", fullResourceName))
				return
			} else if e.Code == 403 && len(e.Errors) > 0 && e.Errors[0].Reason == "forbidden"{
				p.log.Info("access denied. skipping.",
					zap.String("project_id", projectID), zap.String("resource", fullResourceName))
				return
			}
		}
		p.log.Error(err.Error())
		return
	}
}
