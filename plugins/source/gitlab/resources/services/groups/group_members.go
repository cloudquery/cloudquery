// Code generated by codegen; DO NOT EDIT.

package groups

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:     "gitlab_group_members",
		Resolver: fetchGroupMembers,
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "username",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Username"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvatarURL"),
			},
			{
				Name:     "web_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebURL"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "expires_at",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ExpiresAt"),
			},
			{
				Name:     "access_level",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AccessLevel"),
			},
			{
				Name:     "group_saml_identity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GroupSAMLIdentity"),
			},
		},
	}
}
