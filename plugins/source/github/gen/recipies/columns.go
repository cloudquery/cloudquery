package recipies

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	htmlURL = "HTMLURL"
	sshURL  = "SSHURL"
	svnURL  = "SVNURL"
)

var (
	orgColumns = []codegen.ColumnDefinition{
		{
			Name:        "org",
			Description: "The Github Organization of the resource.",
			Type:        schema.TypeString,
			Resolver:    `client.ResolveOrg`,
			Options:     schema.ColumnCreationOptions{PrimaryKey: true},
		},
	}
	htmlURLCol = codegen.ColumnDefinition{
		Name:     "html_url",
		Type:     schema.TypeString,
		Resolver: `schema.PathResolver("` + htmlURL + `")`,
	}
	svnURLCol = codegen.ColumnDefinition{
		Name:     "svn_url",
		Type:     schema.TypeString,
		Resolver: `schema.PathResolver("` + svnURL + `")`,
	}
	sshURLCol = codegen.ColumnDefinition{
		Name:     "ssh_url",
		Type:     schema.TypeString,
		Resolver: `schema.PathResolver("` + sshURL + `")`,
	}
	idColumn = pkColumn("id", "ID")
	skipID   = []string{"ID"}
)

func timestampField(name, path string) codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:     name,
		Type:     schema.TypeTimestamp,
		Resolver: `schema.PathResolver("` + path + `.Time")`,
	}
}

func pkColumn(name, path string) codegen.ColumnDefinition {
	return codegen.ColumnDefinition{
		Name:     name,
		Type:     schema.TypeInt,
		Resolver: `schema.PathResolver("` + path + `")`,
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	}
}
