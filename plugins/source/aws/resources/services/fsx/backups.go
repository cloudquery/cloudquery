package fsx

import (
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Backups() *schema.Table {
	tableName := "aws_fsx_backups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/fsx/latest/APIReference/API_Backup.html`,
		Resolver:    fetchFsxBackups,
		Transform:   client.TransformWithStruct(&types.Backup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "fsx"),
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
