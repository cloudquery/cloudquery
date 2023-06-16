package resiliencehub

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var (
	appARN = schema.Column{
		Name:       "app_arn",
		Type:       arrow.BinaryTypes.String,
		Resolver:   schema.ParentColumnResolver("app_arn"),
		PrimaryKey: true,
	}
	appARNTop = schema.Column{
		Name:       "app_arn",
		Type:       arrow.BinaryTypes.String,
		Resolver:   schema.ParentColumnResolver("arn"),
		PrimaryKey: true,
	}
	assessmentARN = schema.Column{
		Name:       "assessment_arn",
		Type:       arrow.BinaryTypes.String,
		Resolver:   schema.ParentColumnResolver("arn"),
		PrimaryKey: true,
	}
	appVersion = schema.Column{
		Name:       "app_version",
		Type:       arrow.BinaryTypes.String,
		Resolver:   schema.ParentColumnResolver("app_version"),
		PrimaryKey: true,
	}
)

func arnColumn(path string) schema.Column {
	return schema.Column{
		Name:       "arn",
		Type:       arrow.BinaryTypes.String,
		Resolver:   schema.PathResolver(path),
		PrimaryKey: true,
	}
}
