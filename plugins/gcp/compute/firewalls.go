package compute
import (
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)


type Firewall struct {
	ID                uint `gorm:"primarykey"`
	ProjectID         string
	Allowed           []*FirewallAllowed `gorm:"constraint:OnDelete:CASCADE;"`
	CreationTimestamp string
	Denied            []*FirewallDenied `gorm:"constraint:OnDelete:CASCADE;"`
	Description       string
	DestinationRanges []*FirewallDestinationRanges `gorm:"constraint:OnDelete:CASCADE;"`
	Direction         string
	Disabled          bool
	ResourceID        uint64
	Kind              string

	LogConfigEnable bool
	LogConfigMetadata string

	Name string
	Network string
	Priority int64
	SelfLink string
	SourceRanges []*FirewallSourceRange                   `gorm:"constraint:OnDelete:CASCADE;"`
	SourceServiceAccounts []*FirewallSourceServiceAccount `gorm:"constraint:OnDelete:CASCADE;"`
	SourceTags []*FirewallSourceTag                       `gorm:"constraint:OnDelete:CASCADE;"`
	TargetServiceAccounts []*FirewallTargetServiceAccount `gorm:"constraint:OnDelete:CASCADE;"`
	TargetTags []*FirewallTargetTag                       `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Firewall) TableName() string {
	return "gcp_compute_firewalls"
}

type FirewallAllowedPorts struct {
	ID uint`gorm:"primarykey"`
	FirewallAllowedID uint
	Value string
}
func (FirewallAllowedPorts) TableName() string {
	return "gcp_compute_firewall_allowed_ports"
}

type FirewallAllowed struct {
	ID uint `gorm:"primarykey"`
	ProjectID string `gorm:"-"`
	FirewallID uint `neo:"ignore"`
	IPProtocol string
	Ports []*FirewallAllowedPorts `gorm:"constraint:OnDelete:CASCADE;"`
}

func (FirewallAllowed) TableName() string {
	return "gcp_compute_firewall_allowed"
}

type FirewallDeniedPort struct {
	ID uint`gorm:"primarykey"`
	FirewallDeniedID uint
	Value string
}
func (FirewallDeniedPort) TableName() string {
	return "gcp_compute_firewall_denied_ports"
}

type FirewallDenied struct {
	ID uint `gorm:"primarykey"`
	ProjectID string `gorm:"-"`
	FirewallID uint `neo:"ignore"`
	IPProtocol string
	Ports []*FirewallDeniedPort `gorm:"constraint:OnDelete:CASCADE;"`
}

func (FirewallDenied) TableName() string {
	return "gcp_compute_firewall_denied"
}

type FirewallDestinationRanges struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallDestinationRanges) TableName() string {
	return "gcp_firewall_destination_ranges"
}

type FirewallSourceRange struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallSourceRange) TableName() string {
	return "gcp_compute_firewall_source_ranges"
}

type FirewallSourceServiceAccount struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallSourceServiceAccount) TableName() string {
	return "gcp_compute_firewall_source_service_accounts"
}

type FirewallSourceTag struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallSourceTag) TableName() string {
	return "gcp_compute_firewall_source_tags"
}

type FirewallTargetServiceAccount struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallTargetServiceAccount) TableName() string {
	return "gcp_compute_firewall_target_service_accounts"
}

type FirewallTargetTag struct {
	ID uint`gorm:"primarykey"`
	FirewallID uint
	Value string
}
func (FirewallTargetTag) TableName() string {
	return "gcp_compute_firewall_target_tags"
}

func (c *Client) transformFirewalls(values []*compute.Firewall) []*Firewall {
	var tValues []*Firewall
	for _, value := range values {
		tValue := Firewall {
			ProjectID:             c.projectID,
			Allowed:               c.transformFirewallAlloweds(value.Allowed),
			CreationTimestamp:     value.CreationTimestamp,
			Denied:                c.transformFirewallDenieds(value.Denied),
			Description:           value.Description,
			DestinationRanges:     c.transformFirewallDestinationRanges(value.DestinationRanges),
			Direction:             value.Direction,
			Disabled:              value.Disabled,
			ResourceID:            value.Id,
			Kind:                  value.Kind,
			Name:                  value.Name,
			Network:               value.Network,
			Priority:              value.Priority,
			SelfLink:              value.SelfLink,
			SourceRanges:          c.transformFirewallSourceRanges(value.SourceRanges),
			SourceServiceAccounts: c.transformFirewallSourceServiceAccounts(value.SourceServiceAccounts),
			SourceTags:            c.transformFirewallSourceTags(value.SourceTags),
			TargetServiceAccounts: c.transformFirewallTargetServiceAccounts(value.TargetServiceAccounts),
			TargetTags:            c.transformFirewallTargetTags(value.TargetTags),
		}
		if value.LogConfig != nil {

			tValue.LogConfigEnable = value.LogConfig.Enable
			tValue.LogConfigMetadata = value.LogConfig.Metadata

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformFirewallAllowedPorts(values []string) []*FirewallAllowedPorts {
	var tValues []*FirewallAllowedPorts
	for _, v := range values {
		tValues = append(tValues, &FirewallAllowedPorts{
			Value: v,
		})
	}
	return tValues
}


func (c *Client) transformFirewallAlloweds(values []*compute.FirewallAllowed) []*FirewallAllowed {
	var tValues []*FirewallAllowed
	for _, value := range values {
		tValue := FirewallAllowed {
			ProjectID: c.projectID,
			IPProtocol: value.IPProtocol,
			Ports: c.transformFirewallAllowedPorts(value.Ports),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformFirewallDeniedPorts(values []string) []*FirewallDeniedPort {
	var tValues []*FirewallDeniedPort
	for _, v := range values {
		tValues = append(tValues, &FirewallDeniedPort{
			Value: v,
		})
	}
	return tValues
}


func (c *Client) transformFirewallDenieds(values []*compute.FirewallDenied) []*FirewallDenied {
	var tValues []*FirewallDenied
	for _, value := range values {
		tValue := FirewallDenied {
			ProjectID: c.projectID,
			IPProtocol: value.IPProtocol,
			Ports: c.transformFirewallDeniedPorts(value.Ports),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}
func (c *Client) transformFirewallDestinationRanges(values []string) []*FirewallDestinationRanges {
	var tValues []*FirewallDestinationRanges
	for _, v := range values {
		tValues = append(tValues, &FirewallDestinationRanges{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformFirewallSourceRanges(values []string) []*FirewallSourceRange {
	var tValues []*FirewallSourceRange
	for _, v := range values {
		tValues = append(tValues, &FirewallSourceRange{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformFirewallSourceServiceAccounts(values []string) []*FirewallSourceServiceAccount {
	var tValues []*FirewallSourceServiceAccount
	for _, v := range values {
		tValues = append(tValues, &FirewallSourceServiceAccount{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformFirewallSourceTags(values []string) []*FirewallSourceTag {
	var tValues []*FirewallSourceTag
	for _, v := range values {
		tValues = append(tValues, &FirewallSourceTag{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformFirewallTargetServiceAccounts(values []string) []*FirewallTargetServiceAccount {
	var tValues []*FirewallTargetServiceAccount
	for _, v := range values {
		tValues = append(tValues, &FirewallTargetServiceAccount{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformFirewallTargetTags(values []string) []*FirewallTargetTag {
	var tValues []*FirewallTargetTag
	for _, v := range values {
		tValues = append(tValues, &FirewallTargetTag{
			Value: v,
		})
	}
	return tValues
}



var FirewallTables = []interface{} {
	&Firewall{},
	&FirewallAllowed{},
	&FirewallAllowedPorts{},
	&FirewallDenied{},
	&FirewallDeniedPort{},
	&FirewallDestinationRanges{},
	&FirewallSourceRange{},
	&FirewallSourceServiceAccount{},
	&FirewallSourceTag{},
	&FirewallTargetServiceAccount{},
	&FirewallTargetTag{},
}

func (c *Client)firewalls(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(FirewallTables...)
	for {
		c.svc.FirewallPolicies.List()
		call := c.svc.Firewalls.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformFirewalls(output.Items))
		c.log.Info("populating Firewalls", zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

