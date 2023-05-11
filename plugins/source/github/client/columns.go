package client

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	OrgColumn = schema.Column{
		Name:            "org",
		Type:            schema.TypeString,
		Resolver:        ResolveOrg,
		Description:     `The Github Organization of the resource.`,
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	RepositoryIDColumn = schema.Column{
		Name:     "repository_id",
		Type:     schema.TypeInt,
		Resolver: ResolveRepositoryID,
		CreationOptions: schema.ColumnCreationOptions{
			PrimaryKey: true,
		},
	}
)
