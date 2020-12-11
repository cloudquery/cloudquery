package compute

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type SSLCertificate struct {
	ID                      uint `gorm:"primarykey"`
	ProjectID               string
	Certificate             string
	CreationTimestamp       string
	Description             string
	ExpireTime              string
	ResourceID              uint64
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

type SSLCertificateManagedDomain struct {
	ID               uint `gorm:"primarykey"`
	SSLCertificateID uint
	Value            string
}

type SSLCertificateSubjectAlternativeName struct {
	ID               uint `gorm:"primarykey"`
	SSLCertificateID uint
	Value            string
}

func (c *Client) transformSSLCertificateManagedSslCertificateDomains(values []string) []*SSLCertificateManagedDomain {
	var tValues []*SSLCertificateManagedDomain
	for _, v := range values {
		tValues = append(tValues, &SSLCertificateManagedDomain{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformSSLCertificateSubjectAlternativeNames(values []string) []*SSLCertificateSubjectAlternativeName {
	var tValues []*SSLCertificateSubjectAlternativeName
	for _, v := range values {
		tValues = append(tValues, &SSLCertificateSubjectAlternativeName{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformSSLCertificate(value *compute.SslCertificate) *SSLCertificate {
	res := SSLCertificate{
		ProjectID:               c.projectID,
		Certificate:             value.Certificate,
		CreationTimestamp:       value.CreationTimestamp,
		Description:             value.Description,
		ExpireTime:              value.ExpireTime,
		ResourceID:              value.Id,
		Kind:                    value.Kind,
		Name:                    value.Name,
		Region:                  value.Region,
		SelfLink:                value.SelfLink,
		SelfManagedCertificate:  value.SelfManaged.Certificate,
		SubjectAlternativeNames: c.transformSSLCertificateSubjectAlternativeNames(value.SubjectAlternativeNames),
		Type:                    value.Type,
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

func (c *Client) sslCertificates(gConfig interface{}) error {
	var config SslCertificateConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["computeSslCertificate"] {
		err := c.db.AutoMigrate(
			&SSLCertificate{},
			&SSLCertificateManagedDomain{},
			&SSLCertificateSubjectAlternativeName{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["computeSslCertificate"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.SslCertificates.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&SSLCertificate{})
		var tValues []*SSLCertificate
		for _, items := range output.Items {
			tValues = append(tValues, c.transformSSLCertificates(items.SslCertificates)...)
		}
		common.ChunkedCreate(c.db, tValues)
		c.log.Info("populating SSLCertificates", zap.Int("count", len(tValues)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
