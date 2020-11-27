package iam

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/providers/gcp/internal/resource"
	"go.uber.org/zap"
	"google.golang.org/api/iam/v1"
	"gorm.io/gorm"
)

type Client struct {
	db               *gorm.DB
	log              *zap.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *iam.Service
}

func NewClient(db *gorm.DB, log *zap.Logger,
	projectID string, region string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := iam.NewService(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		db:               db,
		log:              log,
		projectID:        projectID,
		region:           region,
		resourceMigrated: map[string]bool{},
		svc:              computeService,
	}, nil
}

func (c *Client) CollectResource(resource string, config interface{}) error {
	switch resource {
	case "service_accounts":
		return c.ServiceAccounts(config)
	case "project_roles":
		return c.ProjectRoles(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
