package efs

import (
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Filesystems() *schema.Table {
	tableName := "aws_efs_filesystems"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/efs/latest/ug/API_FileSystemDescription.html`,
		Resolver:    fetchEfsFilesystems,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticfilesystem"),
		Transform:   transformers.TransformWithStruct(&types.FileSystemDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FileSystemArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "backup_policy_status",
				Type:     schema.TypeString,
				Resolver: ResolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
