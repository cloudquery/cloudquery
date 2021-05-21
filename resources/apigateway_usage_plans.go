package resources

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
		Resolver:     fetchApigatewayUsagePlans,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "product_code",
				Type: schema.TypeString,
			},
			{
				Name:     "quota_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Quota.Limit"),
			},
			{
				Name:     "quota_offset",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Quota.Offset"),
			},
			{
				Name:     "quota_period",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Quota.Period"),
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:     "throttle_burst_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Throttle.BurstLimit"),
			},
			{
				Name:     "throttle_rate_limit",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Throttle.RateLimit"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigateway_usage_plan_api_stages",
				Resolver: fetchApigatewayUsagePlanApiStages,
				Columns: []schema.Column{
					{
						Name:     "usage_plan_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_id",
						Type: schema.TypeString,
					},
					{
						Name: "stage",
						Type: schema.TypeString,
					},
					{
						Name: "throttle",
						Type: schema.TypeJSON,
					},
				},
			},
			{
				Name:     "aws_apigateway_usage_plan_keys",
				Resolver: fetchApigatewayUsagePlanKeys,
				Columns: []schema.Column{
					{
						Name:     "usage_plan_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
					{
						Name: "value",
						Type: schema.TypeString,
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
