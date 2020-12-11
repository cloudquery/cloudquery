package iam

import (
	"github.com/cloudquery/cloudquery/providers/common"
	//"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/iam/v1"
)

type ServiceAccount struct {
	ID             uint `gorm:"primarykey"`
	ProjectID      string
	Region         string
	Description    string
	Disabled       bool
	DisplayName    string
	Email          string
	Etag           string
	Name           string
	Oauth2ClientId string
	UniqueId       string
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

func (c *Client) serviceAccounts(_ interface{}) error {
	if !c.resourceMigrated["iamServiceAccount"] {
		err := c.db.AutoMigrate(
			&ServiceAccount{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["iamServiceAccount"] = true
	}
	nextPageToken := ""
	for {
		call := c.svc.Projects.ServiceAccounts.List("projects/" + c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id = ?", c.projectID).Delete(&ServiceAccount{})
		common.ChunkedCreate(c.db, c.transformServiceAccounts(output.Accounts))
		c.log.Info("Fetched resources", zap.Int("count", len(output.Accounts)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
