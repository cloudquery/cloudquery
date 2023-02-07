package eks

import (
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FargateProfiles() *schema.Table {
	return &schema.Table{
		Name:                "aws_eks_fargate_profiles",
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html`,
		Resolver:            fetchFargateProfiles,
		PreResourceResolver: getFargateProfile,
		Multiplex:           client.ServiceAccountRegionMultiplexer("eks"),
		Transform:           transformers.TransformWithStruct(&types.FargateProfile{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FargateProfileArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{},
	}
}
