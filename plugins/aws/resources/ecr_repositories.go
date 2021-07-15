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
		Description:  "An object representing a repository.",
		Resolver:     fetchEcrRepositories,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "name"}},
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
				Name:        "created_at",
				Description: "The date and time, in JavaScript date format, when the repository was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "encryption_configuration_encryption_type",
				Description: "The encryption type to use.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.EncryptionType"),
			},
			{
				Name:        "encryption_configuration_kms_key",
				Description: "If you use the KMS encryption type, specify the CMK to use for encryption.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionConfiguration.KmsKey"),
			},
			{
				Name:        "image_scanning_configuration_scan_on_push",
				Description: "The setting that determines whether images are scanned after being pushed to a repository.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ImageScanningConfiguration.ScanOnPush"),
			},
			{
				Name:        "image_tag_mutability",
				Description: "The tag mutability setting for the repository.",
				Type:        schema.TypeString,
			},
			{
				Name:        "registry_id",
				Description: "The AWS account ID associated with the registry that contains the repository.",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) that identifies the repository.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryArn"),
			},
			{
				Name:        "name",
				Description: "The name of the repository.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryName"),
			},
			{
				Name:        "uri",
				Description: "The URI for the repository.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RepositoryUri"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ecr_repository_images",
				Description: "An object that describes an image returned by a DescribeImages operation.",
				Resolver:    fetchEcrRepositoryImages,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"repository_cq_id", "image_digest"}},
				Columns: []schema.Column{
					{
						Name:        "repository_cq_id",
						Description: "Unique CloudQuery ID of aws_ecr_repositories table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The AWS Account ID of the resource.",
						Type:        schema.TypeString,
						Resolver:    client.ResolveAWSAccount,
					},
					{
						Name:     "region",
						Type:     schema.TypeString,
						Resolver: client.ResolveAWSRegion,
					},
					{
						Name:        "artifact_media_type",
						Description: "The artifact media type of the image.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_digest",
						Description: "The sha256 digest of the image manifest.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_manifest_media_type",
						Description: "The media type of the image manifest.",
						Type:        schema.TypeString,
					},
					{
						Name:        "image_pushed_at",
						Description: "The date and time, expressed in standard JavaScript date format, at which the current image was pushed to the repository.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "image_scan_findings_summary_finding_severity_counts",
						Description: "The image vulnerability counts, sorted by severity.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.FindingSeverityCounts"),
					},
					{
						Name:        "image_scan_findings_summary_image_scan_completed_at",
						Description: "The time of the last completed image scan.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.ImageScanCompletedAt"),
					},
					{
						Name:        "image_scan_findings_summary_vulnerability_source_updated_at",
						Description: "The time when the vulnerability data was last scanned.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("ImageScanFindingsSummary.VulnerabilitySourceUpdatedAt"),
					},
					{
						Name:        "image_scan_status_description",
						Description: "The description of the image scan status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageScanStatus.Description"),
					},
					{
						Name:        "image_scan_status",
						Description: "The current state of an image scan.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ImageScanStatus.Status"),
					},
					{
						Name:        "image_size_in_bytes",
						Description: "The size, in bytes, of the image in the repository.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "image_tags",
						Description: "The list of tags associated with this image.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "registry_id",
						Description: "The AWS account ID associated with the registry to which this image belongs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "repository_name",
						Description: "The name of the repository to which this image belongs.",
						Type:        schema.TypeString,
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
