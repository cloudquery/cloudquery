package compute

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type SSLCertificate struct {
	_                       interface{} `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                      uint        `gorm:"primarykey"`
	ProjectID               string      `neo:"unique"`
	Certificate             string
	CreationTimestamp       string
	Description             string
	ExpireTime              string
	ResourceID              uint64 `neo:"unique"`
	Kind                    string
	ManagedDomains          []*SSLCertificateManagedDomain `gorm:"constraint:OnDelete:CASCADE;"`
	ManagedStatus           string
	Name                    string
	Region                  string
	SelfLink                string
	SelfManagedCertificate  string
	SubjectAlternativeNames []*SSLCertificateSubjectAlternativeName `gorm:"constraint:OnDelete:CASCADE;"`
	Type                    string
}

func (SSLCertificate) TableName() string {
	return "gcp_compute_ssl_certificates"
}

type SSLCertificateManagedDomain struct {
	ID               uint   `gorm:"primarykey"`
	SSLCertificateID uint   `neo:"ignore"`
	ProjectID        string `gorm:"-"`
	Value            string
}

func (SSLCertificateManagedDomain) TableName() string {
	return "gcp_compute_ssl_certificate_managed_domains"
}

type SSLCertificateSubjectAlternativeName struct {
	ID               uint   `gorm:"primarykey"`
	SSLCertificateID uint   `neo:"ignore"`
	ProjectID        string `gorm:"-"`
	Value            string
}

func (SSLCertificateSubjectAlternativeName) TableName() string {
	return "gcp_compute_ssl_certificate_subject_alternative_names"
}

func (c *Client) transformSSLCertificateManagedSslCertificateDomains(values []string) []*SSLCertificateManagedDomain {
	var tValues []*SSLCertificateManagedDomain
	for _, v := range values {
		tValues = append(tValues, &SSLCertificateManagedDomain{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformSSLCertificateSubjectAlternativeNames(values []string) []*SSLCertificateSubjectAlternativeName {
	var tValues []*SSLCertificateSubjectAlternativeName
	for _, v := range values {
		tValues = append(tValues, &SSLCertificateSubjectAlternativeName{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformSSLCertificate(value *compute.SslCertificate) *SSLCertificate {
	res := SSLCertificate{
		ProjectID:         c.projectID,
		Certificate:       value.Certificate,
		CreationTimestamp: value.CreationTimestamp,
		Description:       value.Description,
		ExpireTime:        value.ExpireTime,
		ResourceID:        value.Id,
		Kind:              value.Kind,
		Name:              value.Name,
		Region:            value.Region,
		SelfLink:          value.SelfLink,

		SubjectAlternativeNames: c.transformSSLCertificateSubjectAlternativeNames(value.SubjectAlternativeNames),
		Type:                    value.Type,
	}
	if value.SelfManaged != nil {
		res.SelfManagedCertificate = value.SelfManaged.Certificate
	}

	if value.Managed != nil {
		res.ManagedStatus = value.Managed.Status
		res.ManagedDomains = c.transformSSLCertificateManagedSslCertificateDomains(value.Managed.Domains)
	}
	return &res
}

func (c *Client) transformSSLCertificates(values []*compute.SslCertificate) []*SSLCertificate {
	var tValues []*SSLCertificate
	for _, v := range values {
		tValues = append(tValues, c.transformSSLCertificate(v))
	}
	return tValues
}

type SslCertificateConfig struct {
	Filter string
}

var SSLCertificateTables = []interface{}{
	&SSLCertificate{},
	&SSLCertificateManagedDomain{},
	&SSLCertificateSubjectAlternativeName{},
}

func (c *Client) sslCertificates(gConfig interface{}) error {
	var config SslCertificateConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	c.db.Where("project_id", c.projectID).Delete(SSLCertificateTables...)
	nextPageToken := ""
	for {
		call := c.svc.SslCertificates.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var tValues []*SSLCertificate
		for _, items := range output.Items {
			tValues = append(tValues, c.transformSSLCertificates(items.SslCertificates)...)
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "compute.ssl_certificates"), zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
