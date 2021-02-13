package iam

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resource"
	"go.uber.org/zap"
	"google.golang.org/api/iam/v1"
)

type Client struct {
	db               *database.Database
	log              *zap.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *iam.Service
}

func NewClient(db *database.Database, log *zap.Logger,
	projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := iam.NewService(ctx)
	if err != nil {
		return nil, err
	}


	return &Client{
		db:               db,
		log:              log,
		projectID:        projectID,
		resourceMigrated: map[string]bool{},
		svc:              computeService,
	}, nil
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "service_accounts":
		return c.serviceAccounts(config)
	case "project_roles":
		return c.projectRoles(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
