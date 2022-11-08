package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

func fetchS3StorageLensConfigurations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().S3control
	config := s3control.ListStorageLensConfigurationsInput{
		AccountId: aws.String(c.AccountID),
	}

	paginator := s3control.NewListStorageLensConfigurationsPaginator(svc, &config)
	var ae smithy.APIError

	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			// This is a workaround, AWS CLI returns "Region is not supported as home region for S3 Storage Lens" on this region
			if c.Region == "ap-northeast-3" && errors.As(err, &ae) && ae.ErrorCode() == "UnknownError" {
				return nil
			}

			return err
		}
		res <- v.StorageLensConfigurationList
	}

	return nil
}

func getStorageLensConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().S3control
	item := resource.Item.(types.ListStorageLensConfigurationEntry)
	response, err := svc.GetStorageLensConfiguration(ctx, &s3control.GetStorageLensConfigurationInput{
		AccountId: aws.String(c.AccountID),
		ConfigId:  item.Id,
	})
	if err != nil {
		return err
	}
	resource.Item = response.StorageLensConfiguration
	return nil
}

func resolveStorageLensTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().S3control
	item := resource.Item.(*types.StorageLensConfiguration)

	output, err := svc.GetStorageLensConfigurationTagging(ctx, &s3control.GetStorageLensConfigurationTaggingInput{
		AccountId: aws.String(cl.AccountID),
		ConfigId:  item.Id,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
