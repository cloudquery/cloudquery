package resources

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var organizationIDCol = codegen.ColumnDefinition{
	Name:     "organization_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveOrganizationID",
}
