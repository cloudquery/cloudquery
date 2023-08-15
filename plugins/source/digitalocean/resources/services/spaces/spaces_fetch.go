package spaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/pkg/errors"
)

const publicAccessURI = "http://acs.amazonaws.com/groups/global/AllUsers"

type WrappedBucket struct {
	types.Bucket
	Location string
	Public   bool
	ACLs     []types.Grant
}

func fetchSpacesSpaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	log := meta.(*client.Client).Logger()

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

func resolveSpaceAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	log := meta.(*client.Client).Logger()
	r := resource.Item.(*WrappedBucket)
	log.Debug().Str("space", *r.Name).Msg("fetching space attributes")

	acls, err := resolveSpacesAcls(ctx, meta, r)
	if err != nil {
		log.Error().Str("space", *r.Name).Msg("failed to fetch space acls")
		return nil
	}
	for _, a := range acls {
		if a.Grantee == nil || a.Grantee.URI == nil {
			continue
		}
		if *a.Grantee.URI == publicAccessURI {
			if err := resource.Set("public", true); err != nil {
				return err
			}
			break
		}
	}
	return resource.Set("acls", acls)
}

func resolveSpacesAcls(ctx context.Context, meta schema.ClientMeta, space *WrappedBucket) ([]types.Grant, error) {
	var ae smithy.APIError
	svc := meta.(*client.Client).Services
	aclOutput, err := svc.Spaces.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: space.Name}, func(options *s3.Options) {
		options.Region = space.Location
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "ServerSideEncryptionConfigurationNotFoundError") {
		return nil, err
	}
	return aclOutput.Grants, nil
}

func fetchSpacesCors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
