package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var OrgColumn = schema.Column{
	Name:            "org",
	Type:            schema.TypeString,
	Resolver:        ResolveOrg,
	Description:     `The Github Organization of the resource.`,
	CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
}
