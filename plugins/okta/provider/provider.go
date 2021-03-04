package provider

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-okta/client"
	"github.com/cloudquery/cq-provider-okta/resources"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

const defaultDomain = "https://<CHANGE_THIS_TO_YOUR_OKTA_DOMAIN>.okta.com"

type Provider struct {
	db     *database.Database
	config Config
	client client.Client
	Logger hclog.Logger
}

type Config struct {
	Domain    string
	Resources []struct {
		Name  string
		Other map[string]interface{}
	}
}


func (p *Provider) Init(driver string, dsn string, verbose bool) error {
	var err error
	p.db, err = database.Open(driver, dsn)
	if err != nil {
		return err
	}

	p.Logger.Info("Updating and creating required provider tables")
	for _, table := range resources.ResourceTables {
		p.Logger.Info("Migrating provider table", "table", table.TableName())
		err := p.db.AutoMigrate(table)
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
		p.Logger.Info("no resources specified. See available resources: see: https://docs.cloudquery.io/okta/tables-reference")
		return nil
	}

	oktaToken, ok := os.LookupEnv("OKTA_TOKEN")
	if !ok {
		return fmt.Errorf("please set OKTA_TOKEN")
	}

	if p.config.Domain == "" || p.config.Domain == defaultDomain {
		return fmt.Errorf(`please set your okta "domain" in config.yml`)
	}

	ctx := context.Background()
	c, err := client.New(ctx, p.db, p.Logger, p.config.Domain, oktaToken)
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, resource := range p.config.Resources {
		r := resource
		g.Go(func() error {
			switch strings.ToLower(r.Name) {
			case "users":
				return c.FetchUsers(ctx)
			case "usertypes":
				return c.FetchUserTypes(ctx)
			case "groups":
				return c.FetchGroups(ctx)
			case "applications":
				return c.FetchApplications(ctx)
			default:
				return fmt.Errorf("unsupported resource %s", resource)
			}
		})
	}
	return g.Wait()
}
