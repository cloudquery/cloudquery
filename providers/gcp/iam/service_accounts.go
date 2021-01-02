package iam

import (
	//"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/iam/v1"
)

type ServiceAccount struct {
	ID             uint   `gorm:"primarykey"`
	ProjectID      string `neo:"unique"`
	Region         string
	Description    string
	Disabled       bool
	DisplayName    string
	Email          string
	Etag           string
	Name           string
	Oauth2ClientId string
	UniqueId       string `neo:"unique"`
}

func (ServiceAccount) TableName() string {
	return "gcp_iam_service_accounts"
}

func (c *Client) transformServiceAccount(value *iam.ServiceAccount) *ServiceAccount {
	return &ServiceAccount{
		Region:         c.region,
		ProjectID:      c.projectID,
		Description:    value.Description,
		Disabled:       value.Disabled,
		DisplayName:    value.DisplayName,
		Email:          value.Email,
		Etag:           value.Etag,
		Name:           value.Name,
		Oauth2ClientId: value.Oauth2ClientId,
		UniqueId:       value.UniqueId,
	}
}

func (c *Client) transformServiceAccounts(values []*iam.ServiceAccount) []*ServiceAccount {
	var tValues []*ServiceAccount
	for _, v := range values {
		tValues = append(tValues, c.transformServiceAccount(v))
	}
	return tValues
}

var ServiceAccountTables = []interface{}{
	&ServiceAccount{},
}

func (c *Client) serviceAccounts(_ interface{}) error {

	c.db.Where("project_id", c.projectID).Delete(ServiceAccountTables...)
	nextPageToken := ""
	for {
		call := c.svc.Projects.ServiceAccounts.List("projects/" + c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformServiceAccounts(output.Accounts))
		c.log.Info("Fetched resources", zap.String("resource", "iam.service_accounts"), zap.Int("count", len(output.Accounts)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
