package kms

import (
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Keys() *schema.Table {
	return &schema.Table{
		Name:                "aws_kms_keys",
		Description:         `https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html`,
		Resolver:            fetchKmsKeys,
		PreResourceResolver: getKey,
		Transform:           transformers.TransformWithStruct(&types.KeyMetadata{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("kms"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "rotation_enabled",
				Type:     schema.TypeBool,
				Resolver: resolveKeysRotationEnabled,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKeysTags,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "replica_keys",
				Type:     schema.TypeJSON,
				Resolver: resolveKeysReplicaKeys,
			},
		},

		Relations: []*schema.Table{
			KeyGrants(),
			KeyPolicies(),
		},
	}
}
