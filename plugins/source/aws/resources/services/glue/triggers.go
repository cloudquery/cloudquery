package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Triggers() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_triggers",
		Description: "Information about a specific trigger",
		Resolver:    fetchGlueTriggers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the trigger.",
				Type:            schema.TypeString,
				Resolver:        resolveGlueTriggerArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
				Resolver:    resolveGlueTriggerTags,
			},
			{
				Name:        "description",
				Description: "A description of this trigger",
				Type:        schema.TypeString,
			},
			{
				Name:        "event_batching_condition_size",
				Description: "Number of events that must be received from Amazon EventBridge before EventBridge event trigger fires",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EventBatchingCondition.BatchSize"),
			},
			{
				Name:        "event_batching_condition_window",
				Description: "Window of time in seconds after which EventBridge event trigger fires",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EventBatchingCondition.BatchWindow"),
			},
			{
				Name:        "id",
				Description: "Reserved for future use",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the trigger",
				Type:        schema.TypeString,
			},
			{
				Name:        "predicate_logical",
				Description: "An optional field if only one condition is listed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Predicate.Logical"),
			},
			{
				Name:        "schedule",
				Description: "A cron expression used to specify the schedule (see Time-Based Schedules for Jobs and Crawlers (https://docsawsamazoncom/glue/latest/dg/monitor-data-warehouse-schedulehtml) For example, to run something every day at 12:15 UTC, you would specify: cron(15 12 * * ? *)",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the trigger",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of trigger that this is",
				Type:        schema.TypeString,
			},
			{
				Name:        "workflow_name",
				Description: "The name of the workflow associated with the trigger",
				Type:        schema.TypeString,
			},
			{
				Name:        "actions",
				Description: "Defines an action to be initiated by a trigger",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Actions"),
			},
			{
				Name:     "predicate",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Predicate"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueTriggers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	return client.ListAndDetailResolver(ctx, meta, res, listTriggers, triggerDetail)
}
func resolveGlueTriggerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name)))
	return resource.Set(c.Name, arn)
}
func resolveGlueTriggerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func triggerARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "trigger", name)
}
func listTriggers(ctx context.Context, meta schema.ClientMeta, detailChan chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glue
	input := glue.ListTriggersInput{MaxResults: aws.Int32(200)}
	for {
		response, err := svc.ListTriggers(ctx, &input)
		if err != nil {
			return err
		}
		for _, item := range response.TriggerNames {
			detailChan <- item
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func triggerDetail(ctx context.Context, meta schema.ClientMeta, resultsChan chan<- interface{}, errorChan chan<- error, listInfo interface{}) {
	c := meta.(*client.Client)
	name := listInfo.(string)
	svc := c.Services().Glue
	dc, err := svc.GetTrigger(ctx, &glue.GetTriggerInput{
		Name: &name,
	})
	if err != nil {
		if c.IsNotFoundError(err) {
			return
		}
		errorChan <- err
		return
	}
	resultsChan <- *dc.Trigger
}
