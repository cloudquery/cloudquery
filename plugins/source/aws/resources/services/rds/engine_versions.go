package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func EngineVersions() *schema.Table {
	tableName := "aws_rds_engine_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_DBEngineVersion.html`,
		Resolver:    fetchRdsEngineVersions,
		Transform:   transformers.TransformWithStruct(&types.DBEngineVersion{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "_engine_version_hash",
				Type:     schema.TypeString,
				Resolver: client.ResolveObjectHash,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tag_list",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("TagList"),
			},
		},

		Relations: []*schema.Table{
			clusterParameters(),
		},
	}
}

func fetchRdsEngineVersions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Rds
	input := &rds.DescribeDBEngineVersionsInput{}
	p := rds.NewDescribeDBEngineVersionsPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.DBEngineVersions
	}
	return nil
}
