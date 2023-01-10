package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchLightsailBuckets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetBucketsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetBuckets(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Buckets
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
func fetchLightsailBucketAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Bucket)
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	input := lightsail.GetBucketAccessKeysInput{
		BucketName: r.Name,
	}
	response, err := svc.GetBucketAccessKeys(ctx, &input)
	if err != nil {
		return err
	}
	res <- response.AccessKeys
	return nil
}
