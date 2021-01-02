package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type Interconnect struct {
	_                       interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                      uint        `gorm:"primarykey"`
	ProjectID               string      `neo:"unique"`
	AdminEnabled            bool
	CircuitInfos            []*InterconnectCircuitInfo `gorm:"constraint:OnDelete:CASCADE;"`
	CreationTimestamp       string
	CustomerName            string
	Description             string
	ExpectedOutages         []*InterconnectOutage `gorm:"constraint:OnDelete:CASCADE;"`
	GoogleIpAddress         string
	GoogleReferenceId       string
	ResourceID              uint64                    `neo:"unique"`
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

func (Interconnect) TableName() string {
	return "gcp_compute_interconnects"
}

type InterconnectCircuitInfo struct {
	ID               uint   `gorm:"primarykey"`
	InterconnectID   uint   `neo:"ignore"`
	ProjectID        string `gorm:"-"`
	CustomerDemarcId string
	GoogleCircuitId  string
	GoogleDemarcId   string
}

func (InterconnectCircuitInfo) TableName() string {
	return "gcp_compute_interconnect_circuit_info"
}

type InterconnectOutage struct {
	ID               uint                                 `gorm:"primarykey"`
	InterconnectID   uint                                 `neo:"ignore"`
	ProjectID        string                               `gorm:"-"`
	AffectedCircuits []*InterconnectOutageAffectedCircuit `gorm:"constraint:OnDelete:CASCADE;"`
	Description      string
	EndTime          int64
	IssueType        string
	Name             string
	Source           string
	StartTime        int64
	State            string
}

func (InterconnectOutage) TableName() string {
	return "gcp_compute_interconnect_outages"
}

type InterconnectOutageAffectedCircuit struct {
	ID                   uint   `gorm:"primarykey"`
	InterconnectOutageID uint   `neo:"ignore"`
	ProjectID            string `gorm:"-"`
	Value                string
}

func (InterconnectOutageAffectedCircuit) TableName() string {
	return "gcp_compute_interconnect_outage_notification_affected_circuits"
}

type InterconnectAttachment struct {
	ID             uint   `gorm:"primarykey"`
	InterconnectID uint   `neo:"ignore"`
	ProjectID      string `gorm:"-"`
	Value          string
}

func (InterconnectAttachment) TableName() string {
	return "gcp_compute_interconnect_attachments"
}

func (c *Client) transformInterconnectCircuitInfo(value *compute.InterconnectCircuitInfo) *InterconnectCircuitInfo {
	return &InterconnectCircuitInfo{
		ProjectID:        c.projectID,
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

func (c *Client) transformInterconnectOutageNotificationAffectedCircuits(values []string) []*InterconnectOutageAffectedCircuit {
	var tValues []*InterconnectOutageAffectedCircuit
	for _, v := range values {
		tValues = append(tValues, &InterconnectOutageAffectedCircuit{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformInterconnectOutageNotification(value *compute.InterconnectOutageNotification) *InterconnectOutage {
	return &InterconnectOutage{
		ProjectID:        c.projectID,
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

func (c *Client) transformInterconnectOutageNotifications(values []*compute.InterconnectOutageNotification) []*InterconnectOutage {
	var tValues []*InterconnectOutage
	for _, v := range values {
		tValues = append(tValues, c.transformInterconnectOutageNotification(v))
	}
	return tValues
}

func (c *Client) transformInterconnectInterconnectAttachments(values []string) []*InterconnectAttachment {
	var tValues []*InterconnectAttachment
	for _, v := range values {
		tValues = append(tValues, &InterconnectAttachment{
			ProjectID: c.projectID,
			Value:     v,
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

var InterconnectTables = []interface{}{
	&Interconnect{},
	&InterconnectCircuitInfo{},
	&InterconnectOutage{},
	&InterconnectOutageAffectedCircuit{},
	&InterconnectAttachment{},
}

func (c *Client) interconnects(gConfig interface{}) error {
	var config InterconnectConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	c.db.Where("project_id", c.projectID).Delete(InterconnectTables...)
	nextPageToken := ""
	for {
		call := c.svc.Interconnects.List(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(output.Items)
		c.log.Info("Fetched resources", zap.String("resource", "compute.interconnects"), zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
