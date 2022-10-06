package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchBackupPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupPlans.html
	for {
		result, err := svc.ListBackupPlans(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.BackupPlansList

		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
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
	for {
		result, err := svc.ListTags(ctx, &params)
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

func fetchBackupPlanSelections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	plan := parent.Item.(*backup.GetBackupPlanOutput)
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
