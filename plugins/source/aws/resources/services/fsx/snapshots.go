package fsx

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_snapshots",
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_Snapshot.html`,
		Resolver:    fetchFsxSnapshots,
		Transform:   transformers.TransformWithStruct(&types.Snapshot{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "administrative_actions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdministrativeActions"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
