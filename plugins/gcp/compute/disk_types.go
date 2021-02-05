package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type DiskType struct {
	ID                    uint `gorm:"primarykey"`
	ProjectID             string
	CreationTimestamp     string
	DefaultDiskSizeGb     int64
	DeprecatedDeleted     string
	DeprecatedDeprecated  string
	DeprecatedObsolete    string
	DeprecatedReplacement string
	DeprecatedState       string
	Description           string
	ResourceID            uint64
	Kind                  string
	Name                  string
	Region                string
	SelfLink              string
	ValidDiskSize         string
	Zone                  string
}

func (DiskType) TableName() string {
	return "gcp_compute_disk_types"
}

func (c *Client) transformDiskType(value *compute.DiskType) *DiskType {
	return &DiskType{
		ProjectID:         c.projectID,
		CreationTimestamp: value.CreationTimestamp,
		DefaultDiskSizeGb: value.DefaultDiskSizeGb,
		//DeprecatedDeleted: value.Deprecated.Deleted,
		//DeprecatedDeprecated: value.Deprecated.Deprecated,
		//DeprecatedObsolete: value.Deprecated.Obsolete,
		//DeprecatedReplacement: value.Deprecated.Replacement,
		//DeprecatedState: value.Deprecated.State,
		Description:   value.Description,
		ResourceID:    value.Id,
		Kind:          value.Kind,
		Name:          value.Name,
		Region:        value.Region,
		SelfLink:      value.SelfLink,
		ValidDiskSize: value.ValidDiskSize,
		Zone:          value.Zone,
	}
}

func (c *Client) transformDiskTypes(values []*compute.DiskType) []*DiskType {
	var tValues []*DiskType
	for _, v := range values {
		tValues = append(tValues, c.transformDiskType(v))
	}
	return tValues
}

type DiskTypeConfig struct {
	Filter string
}

var DiskTypeTables = []interface{}{
	&DiskType{},
}

func (c *Client) diskTypes(gConfig interface{}) error {
	var config DiskTypeConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("project_id", c.projectID).Delete(DiskTypeTables...)
	nextPageToken := ""
	for {
		call := c.svc.DiskTypes.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		call.Filter(config.Filter)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var tValues []*DiskType
		for _, items := range output.Items {
			tValues = append(tValues, c.transformDiskTypes(items.DiskTypes)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.disk_types"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
