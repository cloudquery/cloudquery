package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EcrRepositories() *schema.Table {
	return &schema.Table{
		Name:         "aws_ecr_repositories",
		Resolver:     fetchEcrRepositories,
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
				Name: "created_at",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "encryption_configuration_encryption_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionConfiguration.EncryptionType"),
			},
			{
				Name:     "encryption_configuration_kms_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionConfiguration.KmsKey"),
			},
			{
				Name:     "image_scanning_configuration_scan_on_push",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ImageScanningConfiguration.ScanOnPush"),
			},
			{
				Name: "image_tag_mutability",
				Type: schema.TypeString,
			},
			{
				Name: "registry_id",
				Type: schema.TypeString,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryArn"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryName"),
			},
			{
				Name:     "uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryUri"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ecr_repository_images",
				Resolver: fetchEcrRepositoryImages,
				Columns: []schema.Column{
					{
						Name:     "repository_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
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
						Name: "artifact_media_type",
						Type: schema.TypeString,
					},
					{
						Name: "image_digest",
						Type: schema.TypeString,
					},
					{
						Name: "image_manifest_media_type",
						Type: schema.TypeString,
					},
					{
						Name: "image_pushed_at",
						Type: schema.TypeTimestamp,
					},
					{
						Name:     "image_scan_findings_summary_finding_severity_counts",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("ImageScanFindingsSummary.FindingSeverityCounts"),
					},
					{
						Name:     "image_scan_findings_summary_image_scan_completed_at",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("ImageScanFindingsSummary.ImageScanCompletedAt"),
					},
					{
						Name:     "image_scan_findings_summary_vulnerability_source_updated_at",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("ImageScanFindingsSummary.VulnerabilitySourceUpdatedAt"),
					},
					{
						Name:     "image_scan_status_description",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ImageScanStatus.Description"),
					},
					{
						Name:     "image_scan_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ImageScanStatus.Status"),
					},
					{
						Name: "image_size_in_bytes",
						Type: schema.TypeBigInt,
					},
					{
						Name: "image_tags",
						Type: schema.TypeStringArray,
					},
					{
						Name: "registry_id",
						Type: schema.TypeString,
					},
					{
						Name: "repository_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEcrRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	maxResults := int32(1000)
	config := ecr.DescribeRepositoriesInput{
		MaxResults: &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().ECR
	for {
		output, err := svc.DescribeRepositories(ctx, &config, func(options *ecr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.Repositories
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchEcrRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	maxResults := int32(1000)
	p := parent.Item.(types.Repository)
	config := ecr.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().ECR
	for {
		output, err := svc.DescribeImages(ctx, &config, func(options *ecr.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.ImageDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
