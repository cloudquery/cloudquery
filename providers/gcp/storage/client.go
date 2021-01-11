package storage

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/gcp/resource"
	"go.uber.org/zap"
	"google.golang.org/api/storage/v1"
)

type Client struct {
	db               *database.Database
	log              *zap.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *storage.Service
}

func NewClient(db *database.Database, log *zap.Logger,
	projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := storage.NewService(ctx)
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
	case "buckets":
		return c.buckets(config)
	default:
		return fmt.Errorf("unsupported resource storage.%s", resource)
	}

}
