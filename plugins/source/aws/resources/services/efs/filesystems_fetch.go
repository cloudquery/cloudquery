package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEfsFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config efs.DescribeFileSystemsInput
	c := meta.(*client.Client)
	svc := c.Services().Efs
	for {
		response, err := svc.DescribeFileSystems(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.FileSystems
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.NextMarker
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
	response, err := svc.DescribeBackupPolicy(ctx, &config)
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
