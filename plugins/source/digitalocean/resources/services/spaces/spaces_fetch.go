package spaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"
)

type WrappedBucket struct {
	types.Bucket
	Location string
	Public   bool
	ACLs     []types.Grant
}

func fetchSpaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	log := meta.Logger()

	buckets, err := c.Services.Spaces.ListBuckets(ctx, &s3.ListBucketsInput{}, func(options *s3.Options) {
		options.Region = c.SpacesRegion
	})
	if err != nil {
		if !c.CredentialStatus.Spaces {
			log.Warn().Msg("Spaces credentials not set. skipping")
			return nil
		}

		return err
	}

	wb := make([]*WrappedBucket, len(buckets.Buckets))
	for i, b := range buckets.Buckets {
		wb[i] = &WrappedBucket{
			Bucket:   b,
			Location: c.SpacesRegion,
		}
	}
	res <- wb
	return nil
}

func fetchCors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var ae smithy.APIError
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).Services
	corsOutput, err := svc.Spaces.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
		options.Region = r.Location
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "NoSuchCORSConfiguration") {
		return err
	}
	if corsOutput != nil {
		res <- corsOutput.CORSRules
	}
	return nil
}
