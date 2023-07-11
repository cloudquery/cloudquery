package iam

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func InstanceProfiles() *schema.Table {
	tableName := "aws_iam_instance_profiles"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html`,
		Resolver:    fetchIamInstanceProfiles,
		Transform:   transformers.TransformWithStruct(&types.InstanceProfile{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("InstanceProfileId"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamInstanceProfileTags,
			},
		},
	}
}

func fetchIamInstanceProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.ListInstanceProfilesInput{}
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	p := iam.NewListInstanceProfilesPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.InstanceProfiles
	}
	return nil
}

func resolveIamInstanceProfileTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.InstanceProfile)
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	response, err := svc.ListInstanceProfileTags(ctx, &iam.ListInstanceProfileTagsInput{InstanceProfileName: r.InstanceProfileName}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}
