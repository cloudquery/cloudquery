package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Plans() *schema.Table {
	tableName := "aws_backup_plans"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html`,
		Resolver:            fetchBackupPlans,
		PreResourceResolver: getPlan,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:           transformers.TransformWithStruct(&backup.GetBackupPlanOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupPlanArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolvePlanTags,
			},
		},

		Relations: []*schema.Table{
			planSelections(),
		},
	}
}

func fetchBackupPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupPlans.html
	paginator := backup.NewListBackupPlansPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.BackupPlansList
	}
	return nil
}

func getPlan(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	m := resource.Item.(types.BackupPlansListMember)

	plan, err := svc.GetBackupPlan(
		ctx,
		&backup.GetBackupPlanInput{BackupPlanId: m.BackupPlanId, VersionId: m.VersionId},
	)
	if err != nil {
		return err
	}
	resource.Item = plan
	return nil
}

func resolvePlanTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	plan := resource.Item.(*backup.GetBackupPlanOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: plan.BackupPlanArn}
	tags := make(map[string]string)
	paginator := backup.NewListTagsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for k, v := range page.Tags {
			tags[k] = v
		}
	}
	return resource.Set(c.Name, tags)
}
