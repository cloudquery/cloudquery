package backup

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Vaults() *schema.Table {
	tableName := "aws_backup_vaults"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupVaultListMember.html`,
		Resolver:    fetchBackupVaults,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   transformers.TransformWithStruct(&types.BackupVaultListMember{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "access_policy",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultAccessPolicy,
			},
			{
				Name:     "notifications",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultNotifications,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultTags,
			},
		},

		Relations: []*schema.Table{
			vaultRecoveryPoints(),
		},
	}
}

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
		// This is a service/SDK issue.
		// Workaround is suggested here https://github.com/aws/aws-sdk-go-v2/issues/1885#issuecomment-1282663934
		if strings.Contains(err.Error(), " Failed reading notifications from database for Backup vault ") {
			return nil
		}
		return err
	}
	return resource.Set(col.Name, result)
}
