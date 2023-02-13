package resiliencehub

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var (
	appARN = schema.Column{
		Name:            "app_arn",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("app_arn"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	appARNTop = schema.Column{
		Name:            "app_arn",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("arn"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	assessmentARN = schema.Column{
		Name:            "assessment_arn",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("arn"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
	appVersion = schema.Column{
		Name:            "app_version",
		Type:            schema.TypeString,
		Resolver:        schema.ParentColumnResolver("app_version"),
		CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
	}
)

func transformARN(path string) transformers.NameTransformer {
	return func(field reflect.StructField) (string, error) {
		if field.Name == path {
			return "arn", nil
		}
		return transformers.DefaultNameTransformer(field)
	}
}
