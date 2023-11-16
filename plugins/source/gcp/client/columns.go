package client

import (
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func ProjectIDColumn(pk bool) schema.Column {
	return schema.Column{
		Name:       "project_id",
		Type:       arrow.BinaryTypes.String,
		Resolver:   ResolveProject,
		PrimaryKey: pk,
	}
}
