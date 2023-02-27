package docdb

import (
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EngineVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_engine_versions",
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html`,
		Resolver:    fetchDocdbEngineVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
		Transform: transformers.TransformWithStruct(&types.DBEngineVersion{},
			transformers.WithPrimaryKeys("Engine", "EngineVersion")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},

		Relations: []*schema.Table{
			ClusterParameters(),
			OrderableDbInstanceOptions(),
		},
	}
}
