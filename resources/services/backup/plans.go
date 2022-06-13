package backup

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Plans() *schema.Table {
	return &schema.Table{
		Name:         "aws_backup_plans",
		Description:  "Contains metadata about a backup plan.",
		Resolver:     fetchBackupPlans,
		Multiplex:    client.ServiceAccountRegionMultiplexer("backup"),
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
				Name:        "arn",
				Description: "An Amazon Resource Name (ARN) that uniquely identifies a backup plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupPlanArn"),
			},
			{
				Name:        "id",
				Description: "Uniquely identifies a backup plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupPlanId"),
			},
			{
				Name:        "name",
				Description: "The display name of a saved backup plan.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupPlan.BackupPlanName"),
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
				Resolver:      resolvePlanAdvancedBackupSettings,
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
				Name:        "aws_backup_plan_rules",
				Description: "Specifies a scheduled task used to back up a selection of resources.",
				Resolver:    fetchPlanRules,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
				Columns: []schema.Column{
					{
						Name:        "plan_cq_id",
						Description: "Unique CloudQuery ID of aws_backup_plan table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "A display name for a backup rule.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RuleName"),
					},
					{
						Name:        "target_backup_vault_name",
						Description: "The name of a logical container where backups are stored.",
						Type:        schema.TypeString,
					},
					{
						Name:        "completion_window_minutes",
						Description: "A value in minutes after a backup job is successfully started before it must be completed or it will be canceled by Backup.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:          "copy_actions",
						Description:   "The details of the copy operation.",
						Type:          schema.TypeJSON,
						Resolver:      resolveRuleCopyActions,
						IgnoreInTests: true,
					},
					{
						Name:        "enable_continuous_backup",
						Description: "Specifies whether Backup creates continuous backups.",
						Type:        schema.TypeBool,
					},
					{
						Name:          "delete_after_days",
						Description:   "Specifies the number of days after creation that a recovery point is deleted.",
						Type:          schema.TypeBigInt,
						Resolver:      schema.PathResolver("Lifecycle.DeleteAfterDays"),
						IgnoreInTests: true,
					},
					{
						Name:          "move_to_cold_storage_after_days",
						Description:   "Specifies the number of days after creation that a recovery point is moved to cold storage.",
						Type:          schema.TypeBigInt,
						Resolver:      schema.PathResolver("Lifecycle.MoveToColdStorageAfterDays"),
						IgnoreInTests: true,
					},
					{
						Name:          "recovery_point_tags",
						Description:   "An array of key-value pair strings that are assigned to resources that are associated with this rule when restored from backup.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:        "id",
						Description: "Uniquely identifies a rule that is used to schedule the backup of a selection of resources.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RuleId"),
					},
					{
						Name:        "schedule_expression",
						Description: "A cron expression in UTC specifying when Backup initiates a backup job.",
						Type:        schema.TypeString,
					},
					{
						Name:        "start_window_minutes",
						Description: "A value in minutes after a backup is scheduled before a job will be canceled if it doesn't start successfully.",
						Type:        schema.TypeBigInt,
					},
				},
			},
			{
				Name:        "aws_backup_plan_selections",
				Description: "Contains metadata about a BackupSelection object.",
				Resolver:    fetchBackupSelections,
				IgnoreError: client.IgnoreAccessDeniedServiceDisabled,
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
						Name:        "selection_name",
						Description: "The display name of a resource selection document.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackupSelection.SelectionName"),
					},
					{
						Name:        "conditions",
						Description: "A list of conditions that you define to assign resources to your backup plans using tags.",
						Type:        schema.TypeJSON,
						Resolver:    resolveSelectionConditions,
					},
					{
						Name:        "list_of_tags",
						Description: "A list of conditions that you define to assign resources to your backup plans using tags.",
						Type:        schema.TypeJSON,
						Resolver:    resolveSelectionListOfTags,
					},
					{
						Name:        "not_resources",
						Description: "A list of Amazon Resource Names (ARNs) to exclude from a backup plan.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("BackupSelection.NotResources"),
					},
					{
						Name:        "resources",
						Description: "A list of Amazon Resource Names (ARNs) to assign to a backup plan.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("BackupSelection.Resources"),
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
		result, err := svc.ListBackupPlans(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, m := range result.BackupPlansList {
			plan, err := svc.GetBackupPlan(
				ctx,
				&backup.GetBackupPlanInput{BackupPlanId: m.BackupPlanId, VersionId: m.VersionId},
				func(o *backup.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
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

func resolvePlanAdvancedBackupSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	plan := resource.Item.(backup.GetBackupPlanOutput)
	b, err := json.Marshal(plan.AdvancedBackupSettings)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolvePlanTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	plan := resource.Item.(backup.GetBackupPlanOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: plan.BackupPlanArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) { o.Region = cl.Region })
		if result == nil {
			break
		}
		if err != nil {
			return diag.WrapError(err)
		}
		for k, v := range result.Tags {
			tags[k] = v
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return diag.WrapError(resource.Set(c.Name, tags))
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
		result, err := svc.ListBackupSelections(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, m := range result.BackupSelectionsList {
			s, err := svc.GetBackupSelection(
				ctx,
				&backup.GetBackupSelectionInput{BackupPlanId: plan.BackupPlanId, SelectionId: m.SelectionId},
				func(o *backup.Options) { o.Region = cl.Region },
			)
			if err != nil {
				return diag.WrapError(err)
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

func fetchPlanRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	plan := parent.Item.(backup.GetBackupPlanOutput)
	res <- plan.BackupPlan.Rules
	return nil
}

func resolveRuleCopyActions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rule := resource.Item.(types.BackupRule)
	b, err := json.Marshal(rule.CopyActions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveSelectionConditions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(backup.GetBackupSelectionOutput)
	if s.BackupSelection == nil || s.BackupSelection.Conditions == nil {
		return nil
	}
	b, err := json.Marshal(s.BackupSelection.Conditions)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}

func resolveSelectionListOfTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	s := resource.Item.(backup.GetBackupSelectionOutput)
	if s.BackupSelection == nil || len(s.BackupSelection.ListOfTags) == 0 {
		return nil
	}
	b, err := json.Marshal(s.BackupSelection.ListOfTags)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, b))
}
