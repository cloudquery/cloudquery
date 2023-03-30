package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServerCertificates() *schema.Table {
	tableName := "aws_iam_server_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html`,
		Resolver:    fetchIamServerCertificates,
		Transform:   transformers.TransformWithStruct(&types.ServerCertificateMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCertificateId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
