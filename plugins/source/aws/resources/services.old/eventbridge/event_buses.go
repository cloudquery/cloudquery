package eventbridge

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource event_buses --config gen.hcl --output .
func EventBuses() *schema.Table {
	return &schema.Table{
		Name:         "aws_eventbridge_event_buses",
		Description:  "An event bus receives events from a source and routes them to rules associated with that event bus",
		Resolver:     fetchEventbridgeEventBuses,
		Multiplex:    client.ServiceAccountRegionMultiplexer("events"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEventbridgeEventBusTags,
			},
			{
				Name:        "arn",
				Description: "The ARN of the event bus",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the event bus",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "The permissions policy of the event bus, describing which other Amazon Web Services accounts can write events to this event bus",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_eventbridge_event_bus_rules",
				Description: "Contains information about a rule in Amazon EventBridge",
				Resolver:    fetchEventbridgeEventBusRules,
				Columns: []schema.Column{
					{
						Name:        "event_bus_cq_id",
						Description: "Unique CloudQuery ID of aws_eventbridge_event_buses table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: resolveEventbridgeEventBusRuleTags,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "event_bus_name",
						Description: "The name or ARN of the event bus associated with the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "event_pattern",
						Description: "The event pattern of the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "managed_by",
						Description: "If the rule was created on behalf of your account by an Amazon Web Services service, this field displays the principal name of the service that created the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "role_arn",
						Description: "The Amazon Resource Name (ARN) of the role that is used for target invocation If you're setting an event bus in another account as the target and that account granted permission to your account through an organization instead of directly by the account ID, you must specify a RoleArn with proper permissions in the Target structure, instead of here in this parameter",
						Type:        schema.TypeString,
					},
					{
						Name:        "schedule_expression",
						Description: "The scheduling expression",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "The state of the rule",
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

func fetchEventbridgeEventBuses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input eventbridge.ListEventBusesInput
	c := meta.(*client.Client)
	svc := c.Services().EventBridge
	for {
		response, err := svc.ListEventBuses(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.EventBuses
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func resolveEventbridgeEventBusTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.EventBus).Arn
	return resolveEventBridgeTags(ctx, meta, resource, c, *eventBusArn)
}
func fetchEventbridgeEventBusRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	p := parent.Item.(types.EventBus)
	input := eventbridge.ListRulesInput{
		EventBusName: p.Arn,
	}
	c := meta.(*client.Client)
	svc := c.Services().EventBridge
	for {
		response, err := svc.ListRules(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Rules
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func resolveEventbridgeEventBusRuleTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	eventBusArn := resource.Item.(types.Rule).Arn
	return resolveEventBridgeTags(ctx, meta, resource, c, *eventBusArn)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func resolveEventBridgeTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column, resourceArn string) error {
	cl := meta.(*client.Client)
	svc := cl.Services().EventBridge
	input := eventbridge.ListTagsForResourceInput{
		ResourceARN: &resourceArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(output.Tags)))
}
