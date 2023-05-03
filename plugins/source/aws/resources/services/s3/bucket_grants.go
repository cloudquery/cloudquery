package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/s3/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func bucketGrants() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_grants",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html`,
		Resolver:    fetchS3BucketGrants,
		Transform:   transformers.TransformWithStruct(&types.Grant{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:            "bucket_arn",
				Type:            schema.TypeString,
				Resolver:        schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "grantee_type",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Grantee.Type"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "grantee_id",
				Type:            schema.TypeString,
				Resolver:        resolveBucketGranteeID,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "permission",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("Permission"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
func fetchS3BucketGrants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*models.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services().S3
	region := parent.Get("region").(*schema.Text)
	if region == nil {
		return nil
	}
	aclOutput, err := svc.GetBucketAcl(ctx, &s3.GetBucketAclInput{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = region.Str
	})
	if err != nil {
		if client.IsAWSError(err, "NoSuchBucket") {
			return nil
		}
		return err
	}
	res <- aclOutput.Grants
	return nil
}

func resolveBucketGranteeID(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	grantee := resource.Item.(types.Grant).Grantee
	switch grantee.Type {
	case types.TypeCanonicalUser:
		return resource.Set(c.Name, *grantee.ID)
	case types.TypeAmazonCustomerByEmail:
		return resource.Set(c.Name, *grantee.EmailAddress)
	case types.TypeGroup:
		return resource.Set(c.Name, *grantee.URI)
	default:
		return fmt.Errorf("unsupported grantee type %q", grantee.Type)
	}
}
