package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func InstanceProfiles() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_instance_profiles",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html`,
		Resolver:    fetchIamInstanceProfiles,
		Transform:   transformers.TransformWithStruct(&types.InstanceProfile{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceProfileId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveIamInstanceProfileTags,
			},
		},
	}
}
