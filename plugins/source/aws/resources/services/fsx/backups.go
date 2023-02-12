package fsx

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Backups() *schema.Table {
	return &schema.Table{
		Name:        "aws_fsx_backups",
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html`,
		Resolver:    fetchFsxBackups,
		Transform:   transformers.TransformWithStruct(&types.Backup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("fsx"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupId"),
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
