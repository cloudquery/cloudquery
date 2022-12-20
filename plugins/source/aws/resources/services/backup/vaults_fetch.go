package backup

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchBackupVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupVaults.html
	for {
		result, err := svc.ListBackupVaults(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.BackupVaultList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func resolveVaultTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: vault.BackupVaultArn}
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

func resolveVaultAccessPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	result, err := svc.GetBackupVaultAccessPolicy(
		ctx,
		&backup.GetBackupVaultAccessPolicyInput{BackupVaultName: vault.BackupVaultName},
		func(o *backup.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if result.Policy == nil {
		return nil
	}

	var p map[string]any
	err = json.Unmarshal([]byte(*result.Policy), &p)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, p)
}

func resolveVaultNotifications(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	result, err := svc.GetBackupVaultNotifications(
		ctx,
		&backup.GetBackupVaultNotificationsInput{BackupVaultName: vault.BackupVaultName},
		func(o *backup.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(col.Name, result)
}

func fetchBackupVaultRecoveryPoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	vault := parent.Item.(types.BackupVaultListMember)
	params := backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)}
	for {
		result, err := svc.ListRecoveryPointsByBackupVault(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.RecoveryPoints
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func resolveRecoveryPointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rp := resource.Item.(types.RecoveryPointByBackupVault)
	if rp.ResourceArn == nil || rp.RecoveryPointArn == nil {
		return nil
	}
	resourceARN, err := arn.Parse(*rp.ResourceArn)
	if err != nil {
		return err
	}

	// decide if the backed up resource supports tags
	switch client.AWSService(resourceARN.Service) {
	case client.S3Service, client.EFSService, client.DynamoDBService:
		// these services are ok
	case client.EC2Service:
		if !strings.HasPrefix(resourceARN.Resource, "instance/") {
			return nil
		}
	default:
		// full backup management not supported, so no tags
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "ERROR_2603") {
				// ignoring "ERROR_2603: Cannot find recovery point."
				return nil
			}
			if resourceARN.Service == string(client.DynamoDBService) && client.IsAWSError(err, "ERROR_3930") {
				// advanced backup features are not enabled for dynamodb
				return nil
			}
			return err
		}

		if result == nil {
			break
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
