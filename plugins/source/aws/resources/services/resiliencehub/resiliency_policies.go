package resiliencehub

import (
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResiliencyPolicies() *schema.Table {
	tableName := "aws_resiliencehub_resiliency_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResiliencyPolicy.html`,
		Resolver:    fetchResiliencyPolicies,
		Transform:   client.TransformWithStruct(&types.ResiliencyPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("PolicyArn")},
	}
}
