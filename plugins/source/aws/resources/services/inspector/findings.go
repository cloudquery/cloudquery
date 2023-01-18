package inspector

import (
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "aws_inspector_findings",
		Description: `https://docs.aws.amazon.com/inspector/v1/APIReference/API_Finding.html`,
		Resolver:    fetchInspectorFindings,
		Transform:   transformers.TransformWithStruct(&types.Finding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("inspector"),
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
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("Attributes"),
			},
			{
				Name:     "user_attributes",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTagField("UserAttributes"),
			},
		},
	}
}
