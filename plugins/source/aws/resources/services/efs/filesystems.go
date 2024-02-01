package efs

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
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
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("FileSystemArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "backup_policy_status",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveEfsFilesystemBackupPolicyStatus,
			},
			{
				Name:     "file_system_policy",
				Type:     arrow.BinaryTypes.String,
				Resolver: resolveEfsFilesystemPolicy,
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
	svc := cl.Services(client.AWSServiceEfs).Efs
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

func resolveEfsFilesystemBackupPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.FileSystemDescription)
	config := efs.DescribeBackupPolicyInput{
		FileSystemId: p.FileSystemId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEfs).Efs
	response, err := svc.DescribeBackupPolicy(ctx, &config, func(options *efs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if response.BackupPolicy == nil {
		return nil
	}

	return resource.Set(c.Name, response.BackupPolicy.Status)
}

func resolveEfsFilesystemPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.FileSystemDescription)
	config := efs.DescribeFileSystemPolicyInput{
		FileSystemId: p.FileSystemId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEfs).Efs
	response, err := svc.DescribeFileSystemPolicy(ctx, &config, func(options *efs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if response.Policy == nil {
		return nil
	}

	return resource.Set(c.Name, response.Policy)
}
