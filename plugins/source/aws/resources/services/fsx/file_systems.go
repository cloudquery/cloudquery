package fsx

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FileSystems() *schema.Table {
	tableName := "aws_fsx_file_systems"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileSystem.html`,
		Resolver:    fetchFsxFileSystems,
		Transform:   client.TransformWithStruct(&types.FileSystem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
