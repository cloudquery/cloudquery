package iam

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func userAccessKeys() *schema.Table {
	tableName := "aws_iam_user_access_keys"
	return &schema.Table{
		Name:                 tableName,
		Description:          `https://docs.aws.amazon.com/IAM/latest/APIReference/API_AccessKeyMetadata.html`,
		Resolver:             fetchIamUserAccessKeys,
		PostResourceResolver: postIamUserAccessKeyResolver,
		Transform:            transformers.TransformWithStruct(&models.AccessKeyWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:            client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "user_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:       "access_key_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("AccessKeyId"),
				PrimaryKey: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
			{
				Name: "last_used",
				Type: arrow.FixedWidthTypes.Timestamp_us,
			},
			{
				Name: "last_used_service_name",
				Type: arrow.BinaryTypes.String,
			},
		},
	}
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(*types.User)
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	config.UserName = p.UserName
	paginator := iam.NewListAccessKeysPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		keys := make([]models.AccessKeyWrapper, len(page.AccessKeyMetadata))
		for i, key := range page.AccessKeyMetadata {
			switch i {
			case 0:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
			case 1:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
			default:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key}
			}
		}
		res <- keys
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(models.AccessKeyWrapper)
	if r.AccessKeyId == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return err
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return err
		}
	}
	return nil
}
