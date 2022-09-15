package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func Installations() []*Resource {
	return []*Resource{
		{
			Service:    "installations",
			SubService: "installations",
			Struct:     new(github.Installation),
			TableName:  "installations",
			SkipFields: append(skipID, "HTMLURL"),
			ExtraColumns: append(orgColumns, idColumn,
				codegen.ColumnDefinition{
					Name:     "html_url",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("HTMLURL")`,
				},
			),
		},
	}
}
