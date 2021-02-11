package sql

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cq-provider-gcp/resource"
	"go.uber.org/zap"
	"google.golang.org/api/sql/v1beta4"
)

type Client struct {
	db               *database.Database
	log              *zap.Logger
	projectID        string
	resourceMigrated map[string]bool
	svc              *sql.Service
}

func NewClient(db *database.Database, log *zap.Logger,
	projectID string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := sql.NewService(ctx)
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
	case "instances":
		return c.instances(config)
	default:
		return fmt.Errorf("unsupported resource sql.%s", resource)
	}

}
