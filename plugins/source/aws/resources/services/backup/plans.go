package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:        "aws_backup_plans",
		Description: "Contains metadata about a backup plan.",
		Resolver:    fetchBackupPlans,
		Multiplex:   client.ServiceAccountRegionMultiplexer("backup"),
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
				Description:     "An Amazon Resource Name (ARN) that uniquely identifies a backup plan.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("BackupPlanArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "id",
				Description: "Uniquely identifies a backup plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupPlanId"),
			},
			{
				Name:     "backup_plan",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BackupPlan"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time a resource backup plan is created, in Unix format and Coordinated Universal Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "creator_request_id",
				Description:   "A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "last_execution_date",
				Description:   "The last time a job to back up resources was run with this rule.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "version_id",
				Description: "Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most 1,024 bytes long.",
				Type:        schema.TypeString,
			},
			{
				Name:          "advanced_backup_settings",
				Description:   "Contains a list of backup options for a resource type.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("AdvancedBackupSettings"),
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolvePlanTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_backup_plan_selections",
				Description: "Contains metadata about a BackupSelection object.",
				Resolver:    fetchBackupSelections,

				Columns: []schema.Column{
					{
						Name:        "plan_cq_id",
						Description: "Unique CloudQuery ID of aws_backup_plan table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "creation_date",
						Description: "The date and time a backup plan is created, in Unix format and Coordinated Universal Time (UTC).",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:          "creator_request_id",
						Description:   "A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "iam_role_arn",
						Description: "Specifies the IAM role Amazon Resource Name (ARN) to create the target recovery point; for example, arn:aws:iam::123456789012:role/S3Access.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackupSelection.IamRoleArn"),
					},
					{
						Name:        "selection_id",
						Description: "Uniquely identifies a request to assign a set of resources to a backup plan.",
						Type:        schema.TypeString,
					},
					{
						Name:     "backup_selection",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("BackupSelection"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBackupPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupPlans.html
	for {
		result, err := svc.ListBackupPlans(ctx, &params)
		if err != nil {
			return err
		}
		for _, m := range result.BackupPlansList {
			plan, err := svc.GetBackupPlan(
				ctx,
				&backup.GetBackupPlanInput{BackupPlanId: m.BackupPlanId, VersionId: m.VersionId},
				func(o *backup.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if plan != nil {
				res <- *plan
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func resolvePlanTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	plan := resource.Item.(backup.GetBackupPlanOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: plan.BackupPlanArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if result == nil {
			break
		}
		if err != nil {
			return err
		}
		for k, v := range result.Tags {
			tags[k] = v
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}

func fetchBackupSelections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	plan := parent.Item.(backup.GetBackupPlanOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupSelectionsInput{
		BackupPlanId: plan.BackupPlanId,
		MaxResults:   aws.Int32(1000), // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupSelections.html
	}
	for {
		result, err := svc.ListBackupSelections(ctx, &params)
		if err != nil {
			return err
		}
		for _, m := range result.BackupSelectionsList {
			s, err := svc.GetBackupSelection(
				ctx,
				&backup.GetBackupSelectionInput{BackupPlanId: plan.BackupPlanId, SelectionId: m.SelectionId},
				func(o *backup.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if s != nil {
				res <- *s
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
