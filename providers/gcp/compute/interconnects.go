package compute

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Interconnect struct {
	ID                      uint `gorm:"primarykey"`
	ProjectID               string
	AdminEnabled            bool
	CircuitInfos            []*InterconnectCircuitInfo `gorm:"constraint:OnDelete:CASCADE;"`
	CreationTimestamp       string
	CustomerName            string
	Description             string
	ExpectedOutages         []*InterconnectOutageNotification `gorm:"constraint:OnDelete:CASCADE;"`
	GoogleIpAddress         string
	GoogleReferenceId       string
	ResourceID              uint64
	InterconnectAttachments []*InterconnectAttachment `gorm:"constraint:OnDelete:CASCADE;"`
	InterconnectType        string
	Kind                    string
	LinkType                string
	Location                string
	Name                    string
	NocContactEmail         string
	OperationalStatus       string
	PeerIpAddress           string
	ProvisionedLinkCount    int64
	RequestedLinkCount      int64
	SelfLink                string
	State                   string
}

type InterconnectCircuitInfo struct {
	ID               uint `gorm:"primarykey"`
	InterconnectID   uint
	CustomerDemarcId string
	GoogleCircuitId  string
	GoogleDemarcId   string
}

type InterconnectOutageNotification struct {
	ID               uint `gorm:"primarykey"`
	InterconnectID   uint
	AffectedCircuits []*InterconnectOutageNotificationAffectedCircuit `gorm:"constraint:OnDelete:CASCADE;"`
	Description      string
	EndTime          int64
	IssueType        string
	Name             string
	Source           string
	StartTime        int64
	State            string
}

type InterconnectOutageNotificationAffectedCircuit struct {
	ID                               uint `gorm:"primarykey"`
	InterconnectOutageNotificationID uint
	Value                            string
}
type InterconnectAttachment struct {
	ID             uint `gorm:"primarykey"`
	InterconnectID uint
	Value          string
}

func (c *Client) transformInterconnectCircuitInfo(value *compute.InterconnectCircuitInfo) *InterconnectCircuitInfo {
	return &InterconnectCircuitInfo{
		CustomerDemarcId: value.CustomerDemarcId,
		GoogleCircuitId:  value.GoogleCircuitId,
		GoogleDemarcId:   value.GoogleDemarcId,
	}
}

func (c *Client) transformInterconnectCircuitInfos(values []*compute.InterconnectCircuitInfo) []*InterconnectCircuitInfo {
	var tValues []*InterconnectCircuitInfo
	for _, v := range values {
		tValues = append(tValues, c.transformInterconnectCircuitInfo(v))
	}
	return tValues
}

func (c *Client) transformInterconnectOutageNotificationAffectedCircuits(values []string) []*InterconnectOutageNotificationAffectedCircuit {
	var tValues []*InterconnectOutageNotificationAffectedCircuit
	for _, v := range values {
		tValues = append(tValues, &InterconnectOutageNotificationAffectedCircuit{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInterconnectOutageNotification(value *compute.InterconnectOutageNotification) *InterconnectOutageNotification {
	return &InterconnectOutageNotification{
		AffectedCircuits: c.transformInterconnectOutageNotificationAffectedCircuits(value.AffectedCircuits),
		Description:      value.Description,
		EndTime:          value.EndTime,
		IssueType:        value.IssueType,
		Name:             value.Name,
		Source:           value.Source,
		StartTime:        value.StartTime,
		State:            value.State,
	}
}

func (c *Client) transformInterconnectOutageNotifications(values []*compute.InterconnectOutageNotification) []*InterconnectOutageNotification {
	var tValues []*InterconnectOutageNotification
	for _, v := range values {
		tValues = append(tValues, c.transformInterconnectOutageNotification(v))
	}
	return tValues
}

func (c *Client) transformInterconnectInterconnectAttachments(values []string) []*InterconnectAttachment {
	var tValues []*InterconnectAttachment
	for _, v := range values {
		tValues = append(tValues, &InterconnectAttachment{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInterconnect(value *compute.Interconnect) *Interconnect {
	return &Interconnect{
		ProjectID:               c.projectID,
		AdminEnabled:            value.AdminEnabled,
		CircuitInfos:            c.transformInterconnectCircuitInfos(value.CircuitInfos),
		CreationTimestamp:       value.CreationTimestamp,
		CustomerName:            value.CustomerName,
		Description:             value.Description,
		ExpectedOutages:         c.transformInterconnectOutageNotifications(value.ExpectedOutages),
		GoogleIpAddress:         value.GoogleIpAddress,
		GoogleReferenceId:       value.GoogleReferenceId,
		ResourceID:              value.Id,
		InterconnectAttachments: c.transformInterconnectInterconnectAttachments(value.InterconnectAttachments),
		InterconnectType:        value.InterconnectType,
		Kind:                    value.Kind,
		LinkType:                value.LinkType,
		Location:                value.Location,
		Name:                    value.Name,
		NocContactEmail:         value.NocContactEmail,
		OperationalStatus:       value.OperationalStatus,
		PeerIpAddress:           value.PeerIpAddress,
		ProvisionedLinkCount:    value.ProvisionedLinkCount,
		RequestedLinkCount:      value.RequestedLinkCount,
		SelfLink:                value.SelfLink,
		State:                   value.State,
	}
}

func (c *Client) transformInterconnects(values []*compute.Interconnect) []*Interconnect {
	var tValues []*Interconnect
	for _, v := range values {
		tValues = append(tValues, c.transformInterconnect(v))
	}
	return tValues
}

type InterconnectConfig struct {
	Filter string
}

func (c *Client) interconnects(gConfig interface{}) error {
	var config InterconnectConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeInterconnect"] {
		err := c.db.AutoMigrate(
			&Interconnect{},
			&InterconnectCircuitInfo{},
			&InterconnectOutageNotification{},
			&InterconnectOutageNotificationAffectedCircuit{},
			&InterconnectAttachment{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeInterconnect"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.Interconnects.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&Interconnect{})
		common.ChunkedCreate(c.db, output.Items)
		c.log.Info("Fetched resources", zap.String("resource", "compute.interconnects"), zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
