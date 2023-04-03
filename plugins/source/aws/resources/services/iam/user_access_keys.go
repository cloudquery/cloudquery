package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "access_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKeyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name: "last_used",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "last_used_service_name",
				Type: schema.TypeString,
			},
		},
	}
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(*types.User)
	svc := meta.(*client.Client).Services().Iam
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return err
		}

		keys := make([]models.AccessKeyWrapper, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
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
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(models.AccessKeyWrapper)
	if r.AccessKeyId == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().Iam
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
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
