package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func FsxBackups() *schema.Table {
	return &schema.Table{
		Name:         "aws_fsx_backups",
		Resolver:     fetchFsxBackups,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "backup_id",
				Type: schema.TypeString,
			},
			{
				Name: "creation_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "lifecycle",
				Type: schema.TypeString,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name:     "directory_information_active_directory_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryInformation.ActiveDirectoryId"),
			},
			{
				Name:     "directory_information_domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectoryInformation.DomainName"),
			},
			{
				Name:     "failure_details_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureDetails.Message"),
			},
			{
				Name: "kms_key_id",
				Type: schema.TypeString,
			},
			{
				Name: "progress_percent",
				Type: schema.TypeInt,
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceARN"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveFsxBackupTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchFsxBackups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config fsx.DescribeBackupsInput
	c := meta.(*client.Client)
	svc := c.Services().FSX
	for {
		response, err := svc.DescribeBackups(ctx, &config, func(options *fsx.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Backups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveFsxBackupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Backup)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
