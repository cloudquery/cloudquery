package compute

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/providers/gcp/resource"
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
		return c.instances(config)
	case "images":
		return c.images(config)
	case "addresses":
		return c.addresses(config)
	case "disk_types":
		return c.diskTypes(config)
	case "autoscalers":
		return c.autoscalers(config)
	case "interconnects":
		return c.interconnects(config)
	case "ssl_certificates":
		return c.sslCertificates(config)
	case "vpn_gateways":
		return c.vpnGateways(config)
	default:
		return fmt.Errorf("unsupported resource compute.%s", resource)
	}

}
