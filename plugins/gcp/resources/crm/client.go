package crm

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resources/resource"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/api/cloudresourcemanager/v1"
)

type Client struct {
	db               *database.Database
	log              hclog.Logger
	projectID        string
	region           string
	svc              *cloudresourcemanager.Service
}

func NewClient(db *database.Database, log hclog.Logger,
	projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	service, err := cloudresourcemanager.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:               db,
		log:              log,
		projectID:        projectID,
		svc:              service,
	}, nil
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "projects":
		return c.projects(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
