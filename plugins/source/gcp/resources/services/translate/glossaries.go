package translate

import (
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Glossaries() *schema.Table {
	return &schema.Table{
		Name:        "gcp_translate_glossaries",
		Description: `https://cloud.google.com/translate/docs/reference/rest/v3/projects.locations.glossaries#resource:-glossary`,
		Resolver:    fetchGlossaries,
		Multiplex:   client.ProjectMultiplexEnabledServices("translate.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Glossary{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
