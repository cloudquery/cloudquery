package efs

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FileSystemArn"),
				PrimaryKey: true,
			},
			{
				Name:     "backup_policy_status",
				Type:     arrow.BinaryTypes.String,
				Resolver: ResolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEfsFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config efs.DescribeFileSystemsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Efs
	paginator := efs.NewDescribeFileSystemsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *efs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.FileSystems
	}
	return nil
}

func ResolveEfsFilesystemBackupPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.FileSystemDescription)
	config := efs.DescribeBackupPolicyInput{
		FileSystemId: p.FileSystemId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Efs
	response, err := svc.DescribeBackupPolicy(ctx, &config, func(options *efs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return resource.Set(c.Name, types.StatusDisabled)
		}
		return err
	}
	if response.BackupPolicy == nil {
		return nil
	}

	return resource.Set(c.Name, response.BackupPolicy.Status)
}
