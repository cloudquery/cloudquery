package gcp

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/gcp/compute"
	"github.com/cloudquery/cloudquery/providers/gcp/iam"
	"github.com/cloudquery/cloudquery/providers/gcp/resource"
	"github.com/cloudquery/cloudquery/providers/gcp/sql"
	"github.com/cloudquery/cloudquery/providers/gcp/storage"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/cloudresourcemanager/v1"
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
	ProjectFilter string `mapstructure:"project_filter"`
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
	"sql": sql.NewClient,
}

var tablesArr = [][]interface{}{
	compute.AddressTables,
	compute.AutoscalerTables,
	compute.DiskTypeTables,
	compute.ForwardingRuleTables,
	compute.ImageTables,
	compute.InstanceTables,
	compute.InterconnectTables,
	compute.SSLCertificateTables,
	compute.VPNGatewayTables,
	iam.RoleTables,
	iam.ServiceAccountTables,
	storage.BucketTables,
	sql.DatabaseInstanceTables,
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
	ctx := context.Background()
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		p.log.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/gcp/tables-reference")
		return nil
	}

	var projectIDs []string
	if len(p.config.ProjectIDs) == 0 {
		service, err := cloudresourcemanager.NewService(ctx)
		if err != nil {
			return err
		}

		call := service.Projects.List()
		if p.config.ProjectFilter != "" {
			call.Filter(p.config.ProjectFilter)
		}
		for {
			output, err := call.Do()
			if err != nil {
				return err
			}
			for _, project := range output.Projects {
				projectIDs = append(projectIDs, project.ProjectId)
			}
			if output.NextPageToken == "" {
				break
			}
			call.PageToken(output.NextPageToken)
		}
		p.log.Info("No project_ids specified in config.yml assuming all projects", zap.Int("count", len(projectIDs)))
	} else {
		projectIDs = p.config.ProjectIDs
	}


	for _, projectID := range projectIDs {
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
