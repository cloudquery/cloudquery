package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/pkg/errors"
)

type WrappedBucket struct {
	types.Bucket
	Location string
	Public   bool
	acls     []types.Grant
}

const publicAccessURI = "http://acs.amazonaws.com/groups/global/AllUsers"

func Spaces() *schema.Table {
	return &schema.Table{
		Name:                 "digitalocean_spaces",
		Description:          "",
		Multiplex:            client.SpacesRegionMultiplex,
		Resolver:             fetchSpaces,
		DeleteFilter:         client.DeleteFilter,
		PostResourceResolver: resolveSpaceAttributes,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"name"}},
		Columns: []schema.Column{
			{
				Name:        "name",
				Description: "name of the space.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date",
				Description: "the date of the bucket’s creation.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "location",
				Description: "A “slug” representing the region where the bucket is located (e.g. nyc3).",
				Type:        schema.TypeString,
			},
			{
				Name:        "public",
				Description: "Whether anyone can list the contents of this Space.",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "digitalocean_space_cors",
				Description: " list of elements describing allowed methods for a specific origin.",
				Resolver:    fetchSpaceCorsRules,
				Columns: []schema.Column{
					{
						Name:        "space_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_spaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:          "space_name",
						Description:   "name of the space.",
						Type:          schema.TypeString,
						Resolver:      schema.ParentPathResolver("name"),
						IgnoreInTests: true,
					},
					{
						Name:        "allowed_methods",
						Description: "HTTP methods (e.g. GET) which are allowed from the specified origin.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_origins",
						Description: "orgins from which requests using the specified methods are allowed.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "allowed_headers",
						Description: "headers that will be included in the CORS preflight request’s Access-Control-Request-Headers",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "max_age_seconds",
						Description: "access control max age in seconds",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:          "digitalocean_space_acls",
				Description:   " list of elements describing allowed methods for a specific origin.",
				Resolver:      fetchSpacesAcls,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "space_cq_id",
						Description: "Unique CloudQuery ID of digitalocean_spaces table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "space_name",
						Description: "name of the space.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("name"),
					},
					{
						Name:        "permission",
						Description: "The level of access granted. At this time, the only supported values are FULL_CONTROL and READ.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type of grantee",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.Type"),
					},
					{
						Name:        "display_name",
						Description: "Screen name of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.DisplayName"),
					},
					{
						Name:        "email_address",
						Description: "Email address of the grantee",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.EmailAddress"),
					},
					{
						Name:        "grantee_id",
						Description: "The canonical user ID of the grantee.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.ID"),
					},
					{
						Name:        "uri",
						Description: "A URI specifying a group of users. At this time, only http://acs.amazonaws.com/groups/global/AllUsers is supported.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Grantee.URI"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSpaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	log := meta.Logger()

	buckets, err := svc.S3.ListBuckets(ctx, &s3.ListBucketsInput{}, func(options *s3.Options) {
		options.Region = svc.SpacesRegion
	})
	if err != nil {
		if !svc.CredentialStatus.Spaces {
			log.Warn("Spaces credentials not set. skipping")
			return nil
		}

		return err
	}

	wb := make([]*WrappedBucket, len(buckets.Buckets))
	for i, b := range buckets.Buckets {
		wb[i] = &WrappedBucket{
			Bucket:   b,
			Location: svc.SpacesRegion,
		}
	}
	res <- wb
	return nil
}

func resolveSpaceAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	log := meta.Logger()
	r := resource.Item.(*WrappedBucket)
	log.Debug("fetching space attributes", "space", r.Name)

	acls, err := resolveSpacesAcls(ctx, meta, r)
	if err != nil {
		log.Error("failed to fetch space acls", "space", r.Name)
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
	return nil
}

func resolveSpacesAcls(ctx context.Context, meta schema.ClientMeta, space *WrappedBucket) ([]types.Grant, error) {
	var ae smithy.APIError
	svc := meta.(*client.Client).S3
	aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: space.Name}, func(options *s3.Options) {
		options.Region = space.Location
	})
	if err != nil && !(errors.As(err, &ae) && ae.ErrorCode() == "ServerSideEncryptionConfigurationNotFoundError") {
		return nil, err
	}
	return aclOutput.Grants, nil
}

func fetchSpacesAcls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*WrappedBucket)
	if r == nil {
		return nil
	}
	res <- r.acls
	return nil
}

func fetchSpaceCorsRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var ae smithy.APIError
	r := parent.Item.(*WrappedBucket)
	svc := meta.(*client.Client).S3
	corsOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
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
