package backup

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
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
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("BackupVaultArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "access_policy",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveVaultAccessPolicy,
			},
			{
				Name:     "notifications",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveVaultNotifications,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
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
	svc := cl.Services(client.AWSServiceBackup).Backup
	params := backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupVaults.html
	paginator := backup.NewListBackupVaultsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *backup.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.BackupVaultList
	}
	return nil
}

func resolveVaultTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
	params := backup.ListTagsInput{ResourceArn: vault.BackupVaultArn}
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

func resolveVaultAccessPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceBackup).Backup
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
	svc := cl.Services(client.AWSServiceBackup).Backup
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
