package ecr

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func lifeCyclePolicy() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecr_repository_lifecycle_policies",
		Description: `https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_GetLifecyclePolicy.html`,
		Resolver:    fetchRepositoryLifecyclePolicy,
		Transform: transformers.TransformWithStruct(&ecr.GetLifecyclePolicyOutput{},
			transformers.WithPrimaryKeyComponents("RegistryId"),
			transformers.WithSkipFields("ResultMetadata"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "repository_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "policy_json",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("LifecyclePolicyText"),
			},
		},
	}
}
func fetchRepositoryLifecyclePolicy(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcr).Ecr
	repository := parent.Item.(types.Repository)
	config := ecr.GetLifecyclePolicyInput{
		RepositoryName: repository.RepositoryName,
		RegistryId:     repository.RegistryId,
	}
	resp, err := svc.GetLifecyclePolicy(ctx, &config, func(options *ecr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if client.IsAWSError(err, "LifecyclePolicyNotFound") {
			return nil
		}
		return err
	}
	res <- resp
	return nil
}
