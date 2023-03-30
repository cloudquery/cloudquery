package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DirectoryConfigs() *schema.Table {
	tableName := "aws_appstream_directory_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_DirectoryConfig.html`,
		Resolver:    fetchAppstreamDirectoryConfigs,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   client.TransformWithStruct(&types.DirectoryConfig{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "directory_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
