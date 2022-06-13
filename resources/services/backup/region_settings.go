package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource region_settings --config gen.hcl --output .
func RegionSettings() *schema.Table {
	return &schema.Table{
		Name:         "aws_backup_region_settings",
		Resolver:     fetchBackupRegionSettings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("backup"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region"}},
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
				Name:        "resource_type_management_preference",
				Description: "Returns whether Backup fully manages the backups for a resource type",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "resource_type_opt_in_preference",
				Description: "Returns a list of all services along with the opt-in preferences in the Region.",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBackupRegionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Backup
	input := backup.DescribeRegionSettingsInput{}

	output, err := svc.DescribeRegionSettings(ctx, &input, func(o *backup.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- output
	return nil
}
