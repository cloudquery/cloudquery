package resources

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	smithy "github.com/aws/smithy-go"
)

func ConfigConformancePack() *schema.Table {
	return &schema.Table{
		Name:         "aws_config_conformance_packs",
		Description:  "Returns details of a conformance pack.",
		Resolver:     fetchConfigConformancePacks,
		Multiplex:    client.ServiceAccountRegionMultiplexer("config"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options: schema.TableCreationOptions{
			PrimaryKeys: []string{"arn"},
		},
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
				Description: "Amazon Resource Name (ARN) of the conformance pack.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ConformancePackArn"),
			},
			{
				Name:        "conformance_pack_id",
				Description: "ID of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_name",
				Description: "Name of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_input_parameters",
				Description: "A list of ConformancePackInputParameter objects.",
				Type:        schema.TypeJSON,
				Resolver:    resolveConfigConformancePackConformancePackInputParameters,
			},
			{
				Name:        "created_by",
				Description: "AWS service that created the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_bucket",
				Description: "Amazon S3 bucket where AWS Config stores conformance pack templates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_key_prefix",
				Description: "The prefix for the Amazon S3 bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_update_requested_time",
				Description: "Last time when conformation pack update was requested.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	var ae smithy.APIError
	for {
		resp, err := c.Services().ConfigService.DescribeConformancePacks(ctx, &config, func(options *configservice.Options) {
			options.Region = c.Region
		})

		// This is a workaround until this bug is fixed = https://github.com/aws/aws-sdk-go-v2/issues/1539
		if c.Region == "af-south-1" && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
			return nil
		}

		if err != nil {
			return err
		}
		res <- resp.ConformancePackDetails
		if resp.NextToken == nil {
			break
		}
		config.NextToken = resp.NextToken
	}
	return nil
}
func resolveConfigConformancePackConformancePackInputParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	conformancePack := resource.Item.(types.ConformancePackDetail)
	params := make(map[string]*string, len(conformancePack.ConformancePackInputParameters))
	for _, p := range conformancePack.ConformancePackInputParameters {
		params[*p.ParameterName] = p.ParameterValue
	}
	return resource.Set(c.Name, params)
}
