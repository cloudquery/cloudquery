package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource usage_plans --config usage_plans.hcl --output .
func UsagePlans() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_usage_plans",
		Description:  "Represents a usage plan used to specify who can assess associated API stages",
		Resolver:     fetchApigatewayUsagePlans,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayUsagePlanArn,
			},
			{
				Name:        "description",
				Description: "The description of a usage plan",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier of a UsagePlan resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of a usage plan",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_code",
				Description: "The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace",
				Type:        schema.TypeString,
			},
			{
				Name:        "quota_limit",
				Description: "The target maximum number of requests that can be made in a given time period",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Quota.Limit"),
			},
			{
				Name:        "quota_offset",
				Description: "The number of requests subtracted from the given limit in the initial time period",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Quota.Offset"),
			},
			{
				Name:        "quota_period",
				Description: "The time period in which the limit applies",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Quota.Period"),
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "throttle_burst_limit",
				Description: "The API target request burst rate limit",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Throttle.BurstLimit"),
			},
			{
				Name:        "throttle_rate_limit",
				Description: "The API target request rate limit",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Throttle.RateLimit"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigateway_usage_plan_api_stages",
				Description: "API stage name of the associated API stage in a usage plan",
				Resolver:    schema.PathTableResolver("ApiStages"),
				Columns: []schema.Column{
					{
						Name:        "usage_plan_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_usage_plans table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "usage_plan_id",
						Description: "The identifier of a UsagePlan resource",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "api_id",
						Description: "API Id of the associated API stage in a usage plan",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "API stage name of the associated API stage in a usage plan",
						Type:        schema.TypeString,
					},
					{
						Name:        "throttle",
						Description: "Map containing method level throttling information for API stage in a usage plan",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "aws_apigateway_usage_plan_keys",
				Description: "Represents a usage plan key to identify a plan customer",
				Resolver:    fetchApigatewayUsagePlanKeys,
				Columns: []schema.Column{
					{
						Name:        "usage_plan_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_usage_plans table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "usage_plan_id",
						Description: "The identifier of a UsagePlan resource",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayUsagePlanKeyArn,
					},
					{
						Name:        "id",
						Description: "The Id of a usage plan key",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of a usage plan key",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of a usage plan key",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The value of a usage plan key",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayUsagePlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetUsagePlansInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetUsagePlansPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayUsagePlanArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	up := resource.Item.(types.UsagePlan)
	arn := cl.RegionGlobalARN(client.ApigatewayService, usagePlanIDPart, *up.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayUsagePlanKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.UsagePlan)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetUsagePlanKeysInput{UsagePlanId: r.Id}
	for p := apigateway.NewGetUsagePlanKeysPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayUsagePlanKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	up := resource.Parent.Item.(types.UsagePlan)
	key := resource.Item.(types.UsagePlanKey)
	arn := cl.RegionGlobalARN(client.ApigatewayService, usagePlanIDPart, *up.Id, "keys", *key.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
}
