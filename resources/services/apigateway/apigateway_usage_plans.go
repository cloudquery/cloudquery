package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayUsagePlans() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_usage_plans",
		Description:  "Represents a usage plan than can specify who can assess associated API stages with specified request limits and quotas.",
		Resolver:     fetchApigatewayUsagePlans,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "description",
				Description: "The description of a usage plan.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier of a UsagePlan resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "name",
				Description: "The name of a usage plan.",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_code",
				Description: "The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.",
				Type:        schema.TypeString,
			},
			{
				Name:        "quota_limit",
				Description: "The maximum number of requests that can be made in a given time period.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Quota.Limit"),
			},
			{
				Name:        "quota_offset",
				Description: "The day that a time period starts. For example, with a time period of WEEK, an offset of 0 starts on Sunday, and an offset of 1 starts on Monday.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Quota.Offset"),
			},
			{
				Name:        "quota_period",
				Description: "The time period in which the limit applies. Valid values are \"DAY\", \"WEEK\" or \"MONTH\".",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Quota.Period"),
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "throttle_burst_limit",
				Description: "The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Throttle.BurstLimit"),
			},
			{
				Name:        "throttle_rate_limit",
				Description: "The API request steady-state rate limit.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Throttle.RateLimit"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigateway_usage_plan_api_stages",
				Description: "API stage name of the associated API stage in a usage plan.",
				Resolver:    fetchApigatewayUsagePlanApiStages,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"usage_plan_cq_id", "api_id", "stage"}},
				Columns: []schema.Column{
					{
						Name:        "usage_plan_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_usage_plans table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "usage_plan_id",
						Description: "The identifier of a UsagePlan resource.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "api_id",
						Description: "API Id of the associated API stage in a usage plan.",
						Type:        schema.TypeString,
					},
					{
						Name:        "stage",
						Description: "API stage name of the associated API stage in a usage plan.",
						Type:        schema.TypeString,
					},
					{
						Name:        "throttle",
						Description: "Map containing method level throttling information for API stage in a usage plan.",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "aws_apigateway_usage_plan_keys",
				Description: "Represents a usage plan key to identify a plan customer.",
				Resolver:    fetchApigatewayUsagePlanKeys,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"usage_plan_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "usage_plan_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_usage_plans table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "usage_plan_id",
						Description: "The identifier of a UsagePlan resource.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "id",
						Description: "The Id of a usage plan key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "name",
						Description: "The name of a usage plan key.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of a usage plan key. Currently, the valid key type is API_KEY.",
						Type:        schema.TypeString,
					},
					{
						Name:        "value",
						Description: "The value of a usage plan key.",
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
func fetchApigatewayUsagePlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigateway.GetUsagePlansInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetUsagePlans(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayUsagePlanApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.UsagePlan)
	if !ok {
		return fmt.Errorf("expected UsagePlan but got %T", r)
	}
	res <- r.ApiStages
	return nil
}
func fetchApigatewayUsagePlanKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.UsagePlan)
	if !ok {
		return fmt.Errorf("expected UsagePlan but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetUsagePlanKeysInput{UsagePlanId: r.Id}
	for {
		response, err := svc.GetUsagePlanKeys(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
