package translate

import (
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Glossaries() *schema.Table {
	return &schema.Table{
		Name:        "gcp_translate_glossaries",
		Description: `https://cloud.google.com/translate/docs/reference/rest/v3/projects.locations.glossaries#resource:-glossary`,
		Resolver:    fetchGlossaries,
		Multiplex:   client.ProjectMultiplexEnabledServices("translate.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Glossary{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
