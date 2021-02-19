package compute
import (
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)


type Subnetwork struct {
	ID uint `gorm:"primarykey"`
	ProjectID string
	CreationTimestamp string
	Description string
	EnableFlowLogs bool
	Fingerprint string
	GatewayAddress string
	Id uint64
	IpCidrRange string
	Ipv6CidrRange string
	Kind string

	LogConfigAggregationInterval string
	LogConfigEnable bool
	LogConfigFilterExpr string
	LogConfigFlowSampling float64
	LogConfigMetadata string
	LogConfigMetadataFields []*SubnetworkLogConfigMetadataField `gorm:"constraint:OnDelete:CASCADE;"`

	Name string
	Network string
	PrivateIpGoogleAccess bool
	PrivateIpv6GoogleAccess string
	Purpose string
	Region string
	Role string
	SecondaryIpRanges []*SubnetworkSecondaryRange `gorm:"constraint:OnDelete:CASCADE;"`
	SelfLink string
	State string
}

func (Subnetwork) TableName() string {
	return "gcp_compute_subnetworks"
}

type SubnetworkLogConfigMetadataField struct {
	ID uint`gorm:"primarykey"`
	SubnetworkID uint
	Value string
}
func (SubnetworkLogConfigMetadataField) TableName() string {
	return "gcp_compute_subnetwork_log_config_metadata_fields"
}

type SubnetworkSecondaryRange struct {
	ID uint `gorm:"primarykey"`
	ProjectID string `gorm:"-"`
	SubnetworkID uint `neo:"ignore"`
	IpCidrRange string
	RangeName string
}

func (SubnetworkSecondaryRange) TableName() string {
	return "gcp_compute_subnetwork_secondary_ranges"
}

func (c *Client) transformSubnetworks(values []*compute.Subnetwork) []*Subnetwork {
	var tValues []*Subnetwork
	for _, value := range values {
		tValue := Subnetwork {
			ProjectID: c.projectID,
			CreationTimestamp: value.CreationTimestamp,
			Description: value.Description,
			EnableFlowLogs: value.EnableFlowLogs,
			Fingerprint: value.Fingerprint,
			GatewayAddress: value.GatewayAddress,
			Id: value.Id,
			IpCidrRange: value.IpCidrRange,
			Ipv6CidrRange: value.Ipv6CidrRange,
			Kind: value.Kind,
			Name: value.Name,
			Network: value.Network,
			PrivateIpGoogleAccess: value.PrivateIpGoogleAccess,
			PrivateIpv6GoogleAccess: value.PrivateIpv6GoogleAccess,
			Purpose: value.Purpose,
			Region: value.Region,
			Role: value.Role,
			SecondaryIpRanges: c.transformSubnetworkSecondaryRanges(value.SecondaryIpRanges),
			SelfLink: value.SelfLink,
			State: value.State,
		}
		if value.LogConfig != nil {

			tValue.LogConfigAggregationInterval = value.LogConfig.AggregationInterval
			tValue.LogConfigEnable = value.LogConfig.Enable
			tValue.LogConfigFilterExpr = value.LogConfig.FilterExpr
			tValue.LogConfigFlowSampling = value.LogConfig.FlowSampling
			tValue.LogConfigMetadata = value.LogConfig.Metadata
			tValue.LogConfigMetadataFields = c.transformSubnetworkLogConfigMetadataFields(value.LogConfig.MetadataFields)

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformSubnetworkLogConfigMetadataFields(values []string) []*SubnetworkLogConfigMetadataField {
	var tValues []*SubnetworkLogConfigMetadataField
	for _, v := range values {
		tValues = append(tValues, &SubnetworkLogConfigMetadataField{
			Value: v,
		})
	}
	return tValues
}


func (c *Client) transformSubnetworkSecondaryRanges(values []*compute.SubnetworkSecondaryRange) []*SubnetworkSecondaryRange {
	var tValues []*SubnetworkSecondaryRange
	for _, value := range values {
		tValue := SubnetworkSecondaryRange {
			ProjectID: c.projectID,
			IpCidrRange: value.IpCidrRange,
			RangeName: value.RangeName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}


var SubnetworkTables = []interface{} {
	&Subnetwork{},
	&SubnetworkLogConfigMetadataField{},
	&SubnetworkSecondaryRange{},
}

func (c *Client)subnetworks(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(SubnetworkTables...)
	for {
		call := c.svc.Subnetworks.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}
		count := 0
		for _, scopedNetworkList := range output.Items {
			c.db.ChunkedCreate(c.transformSubnetworks(scopedNetworkList.Subnetworks))
			count += len(scopedNetworkList.Subnetworks)
		}

		c.log.Info("populating Subnetworks", zap.Int("count", count))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

