package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var OrganizationID = schema.Column{
	Name:       "organization_id",
	Type:       arrow.BinaryTypes.String,
	PrimaryKey: true,
	NotNull:    true,
	Resolver:   ResolveOrganizationID,
}
