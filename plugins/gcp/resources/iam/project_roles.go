package iam

import (
	"google.golang.org/api/iam/v1"
)

type Role struct {
	ID                  uint `gorm:"primarykey"`
	ProjectID           string
	Region              string
	Deleted             bool
	Description         string
	Etag                string
	IncludedPermissions []*RolePermission `gorm:"constraint:OnDelete:CASCADE;"`
	Name                string
	Stage               string
	Title               string
}

func (Role) TableName() string {
	return "gcp_iam_roles"
}

type RolePermission struct {
	ID        uint   `gorm:"primarykey"`
	RoleID    uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     string
}

func (RolePermission) TableName() string {
	return "gcp_iam_role_permissions"
}

func (c *Client) transformRolePermission(value string) *RolePermission {
	return &RolePermission{
		ProjectID: c.projectID,
		Value:     value,
	}
}

func (c *Client) transformRolePermissions(values []string) []*RolePermission {
	var tValues []*RolePermission
	for _, v := range values {
		tValues = append(tValues, c.transformRolePermission(v))
	}
	return tValues
}

func (c *Client) transformRole(value *iam.Role) *Role {
	return &Role{
		Region:              c.region,
		ProjectID:           c.projectID,
		Deleted:             value.Deleted,
		Description:         value.Description,
		Etag:                value.Etag,
		IncludedPermissions: c.transformRolePermissions(value.IncludedPermissions),
		Name:                value.Name,
		Stage:               value.Stage,
		Title:               value.Title,
	}
}

func (c *Client) transformRoles(values []*iam.Role) []*Role {
	var tValues []*Role
	for _, v := range values {
		tValues = append(tValues, c.transformRole(v))
	}
	return tValues
}

var RoleTables = []interface{}{
	&Role{},
	&RolePermission{},
}

func (c *Client) projectRoles(_ interface{}) error {
	c.db.Where("region", c.region).Where("project_id", c.projectID).Delete(RoleTables...)
	nextPageToken := ""
	for {
		call := c.svc.Projects.Roles.List("projects/" + c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformRoles(output.Roles))
		c.log.Info("Fetched resources", "resource", "iam.project_roles", "count", len(output.Roles))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
