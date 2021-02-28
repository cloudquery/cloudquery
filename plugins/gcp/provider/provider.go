package provider

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resources/cloudfunctions"
	"github.com/cloudquery/cq-provider-gcp/resources/compute"
	"github.com/cloudquery/cq-provider-gcp/resources/crm"
	"github.com/cloudquery/cq-provider-gcp/resources/iam"
	"github.com/cloudquery/cq-provider-gcp/resources/kms"
	"github.com/cloudquery/cq-provider-gcp/resources/resource"
	"github.com/cloudquery/cq-provider-gcp/resources/sql"
	"github.com/cloudquery/cq-provider-gcp/resources/storage"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/googleapi"
	"gopkg.in/yaml.v3"
	"strings"
)

type Provider struct {
	db              *database.Database
	config          Config
	resourceClients map[string]resource.ClientInterface
	Logger             hclog.Logger
}

type Config struct {
	ProjectFilter string   `yaml:"project_filter"`
	ProjectIDs    []string `yaml:"project_ids"`
	Resources     []struct {
		Name  string
		Other map[string]interface{} `yaml:",inline"`
	}
}

type NewResourceFunc func(db *database.Database, log hclog.Logger,
	projectID string) (resource.ClientInterface, error)

var resourceFactory = map[string]NewResourceFunc{
	"compute": compute.NewClient,
	"iam":     iam.NewClient,
	"storage": storage.NewClient,
	"sql":     sql.NewClient,
	"crm":		crm.NewClient,
	"cloudfunctions": cloudfunctions.NewClient,
	"kms": kms.NewClient,
}

var tablesArr = [][]interface{}{
	compute.AddressTables,
	compute.AutoscalerTables,
	compute.BackendServiceTables,
	compute.DiskTypeTables,
	compute.ForwardingRuleTables,
	compute.ImageTables,
	compute.InstanceTables,
	compute.InterconnectTables,
	compute.NetworkTables,
	compute.SSLCertificateTables,
	compute.VPNGatewayTables,
	compute.SubnetworkTables,
	compute.FirewallTables,
	crm.ProjectTables,
	iam.RoleTables,
	iam.ServiceAccountTables,
	storage.BucketTables,
	sql.DatabaseInstanceTables,
	cloudfunctions.CloudFunctionTables,
	kms.CryptoKeyTables,
}

func (p *Provider) Init(driver string, dsn string, verbose bool) error {
	var err error
	p.db, err = database.Open(driver, dsn)
	if err != nil {
		return err
	}
	p.Logger.Info("Creating tables if needed")
	for _, table := range tablesArr {
		err := p.db.AutoMigrate(table...)
		if err != nil {
			return err
		}
	}

	p.resourceClients = map[string]resource.ClientInterface{}

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
	ctx := context.Background()
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		p.Logger.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/gcp/tables-reference")
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
		p.Logger.Info("No project_ids specified in config.yml assuming all projects", "count", len(projectIDs))
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
		g := errgroup.Group{}
		for _, resource := range p.config.Resources {
			resourceName := resource.Name
			resourceConfig := resource.Other
			g.Go(func() error {
				return p.collectResource(projectID, resourceName, resourceConfig)
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
	}

	return nil
}

func (p *Provider) initClients(projectID string) error {
	var err error
	for serviceName, newFunc := range resourceFactory {
		innerLog := p.Logger.With("project_id", projectID)
		p.resourceClients[serviceName], err = newFunc(p.db, innerLog, projectID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Provider) collectResource(projectID string, fullResourceName string, config interface{}) error{
	resourcePath := strings.Split(fullResourceName, ".")
	if len(resourcePath) != 2 {
		return fmt.Errorf("resource %s should be in format {service}.{resource}", fullResourceName)
	}
	service := resourcePath[0]
	resourceName := resourcePath[1]

	if resourceFactory[service] == nil {
		return fmt.Errorf("unsupported service %s", service)
	}

	err := p.resourceClients[service].CollectResource(resourceName, config)
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == 403 && len(e.Errors) > 0 && e.Errors[0].Reason == "accessNotConfigured" {
				p.Logger.Info("access not configured. skipping.",
					"project_id", projectID, "resource", fullResourceName)
				return nil
			} else if e.Code == 403 && len(e.Errors) > 0 && e.Errors[0].Reason == "forbidden" {
				p.Logger.Info("access denied. skipping.",
					"project_id", projectID, "resource", fullResourceName)
				return nil
			}
		}
		return err
	}
	return nil
}


