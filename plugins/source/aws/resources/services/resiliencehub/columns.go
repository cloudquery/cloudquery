package resiliencehub

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	appARN = schema.Column{
		Name:            "app_arn",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("app_arn"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	assessmentARN = schema.Column{
		Name:            "assessment_arn",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("assessment_arn"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	appVersion = schema.Column{
		Name:            "app_version",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("app_version"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
)
