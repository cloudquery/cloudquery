package glue

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Triggers() *schema.Table {
	tableName := "aws_glue_triggers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_Trigger.html`,
		Resolver:            fetchGlueTriggers,
		PreResourceResolver: getTrigger,
		Transform:           transformers.TransformWithStruct(&types.Trigger{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGlueTriggerArn,
				PrimaryKeyComponent: true,
			},
			tagsCol(func(cl *client.Client, resource *schema.Resource) string {
				return triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name))
			}),
		},
	}
}

func fetchGlueTriggers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	input := glue.ListTriggersInput{MaxResults: aws.Int32(200)}
	paginator := glue.NewListTriggersPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TriggerNames
	}
	return nil
}

func getTrigger(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	name := resource.Item.(string)
	svc := cl.Services(client.AWSServiceGlue).Glue
	dc, err := svc.GetTrigger(ctx, &glue.GetTriggerInput{
		Name: &name,
	}, func(options *glue.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.Trigger
	return nil
}

func resolveGlueTriggerArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, triggerARN(cl, aws.ToString(resource.Item.(types.Trigger).Name)))
}

func triggerARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "trigger/" + name,
	}.String()
}
