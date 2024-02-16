package backup

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func ReportPlans() *schema.Table {
	tableName := "aws_backup_report_plans"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ReportPlan.html`,
		Resolver:    fetchReportPlans,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   transformers.TransformWithStruct(&types.ReportPlan{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ReportPlanArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveReportPlanTags,
			},
		},
	}
}

func fetchReportPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
	paginator := backup.NewListReportPlansPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ReportPlans
	}
	return nil
}

func resolveReportPlanTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	plan := resource.Item.(types.ReportPlan)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
	params := backup.ListTagsInput{ResourceArn: plan.ReportPlanArn}
	tags := make(map[string]string)
	paginator := backup.NewListTagsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for k, v := range page.Tags {
			tags[k] = v
		}
	}
	return resource.Set(c.Name, tags)
}
