package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SshPublicKeys() *schema.Table {
	tableName := "aws_iam_ssh_public_keys"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SSHPublicKeyMetadata.html`,
		Resolver:    fetchIamSshPublicKeys,
		Transform:   client.TransformWithStruct(&types.SSHPublicKeyMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "ssh_public_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SSHPublicKeyId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
