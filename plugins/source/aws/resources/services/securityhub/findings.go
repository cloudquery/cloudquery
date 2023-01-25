package securityhub

import (
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "aws_sqs_queues",
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html`,
		Resolver:    fetchFindings,
		Transform:   transformers.TransformWithStruct(&types.AwsSecurityFinding{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("securityhub"),
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
		},
	}
}
