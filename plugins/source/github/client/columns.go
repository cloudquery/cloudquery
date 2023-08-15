package client

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	OrgColumn = schema.Column{
		Name:        "org",
		Type:        arrow.BinaryTypes.String,
		Resolver:    ResolveOrg,
		Description: `The Github Organization of the resource.`,
		PrimaryKey:  true,
	}
	RepositoryIDColumn = schema.Column{
		Name:       "repository_id",
		Type:       arrow.PrimitiveTypes.Int64,
		Resolver:   ResolveRepositoryID,
		PrimaryKey: true,
	}
)
