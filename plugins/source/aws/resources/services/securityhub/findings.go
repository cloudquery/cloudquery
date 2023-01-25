package securityhub

import (
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "aws_securityhub_findings",
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetFindings.html`,
		Resolver:    fetchFindings,
		Transform:   transformers.TransformWithStruct(&types.AwsSecurityFinding{}, transformers.WithPrimaryKeys("AwsAccountId", "Region", "CreatedAt", "Description", "GeneratorId,Id", "ProductArn", "SchemaVersion", "Title")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("securityhub"),
		Columns: []schema.Column{
			{
				Name:     "request_account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "request_region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
		},
	}
}
