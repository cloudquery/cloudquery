package compute

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/providers/gcp/internal/resource"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
	"gorm.io/gorm"
)

type Client struct {
	db               *gorm.DB
	log              *zap.Logger
	projectID        string
	region           string
	resourceMigrated map[string]bool
	svc              *compute.Service
}

func NewClient(db *gorm.DB, log *zap.Logger,
	projectID string, region string) (resource.ClientInterface, error) {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
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
	case "instances":
		return c.Instances(config)
	case "images":
		return c.Images(config)
	case "addresses":
		return c.Addresses(config)
	case "disk_types":
		return c.DiskTypes(config)
	case "autoscalers":
		return c.Autoscalers(config)
	case "interconnects":
		return c.Interconnects(config)
	case "ssl_certificates":
		return c.SSLCertificates(config)
	case "vpn_gateways":
		return c.VpnGateways(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
