package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ConfigConformancePack() *schema.Table {
	return &schema.Table{
		Name:         "aws_config_conformance_packs",
		Resolver:     fetchConfigConformancePacks,
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
				Name: "conformance_pack_arn",
				Type: schema.TypeString,
			},
			{
				Name: "conformance_pack_id",
				Type: schema.TypeString,
			},
			{
				Name: "conformance_pack_name",
				Type: schema.TypeString,
			},
			{
				Name:     "conformance_pack_input_parameters",
				Type:     schema.TypeJSON,
				Resolver: resolveConfigConformancePackConformancePackInputParameters,
			},
			{
				Name: "created_by",
				Type: schema.TypeString,
			},
			{
				Name: "delivery_s3_bucket",
				Type: schema.TypeString,
			},
			{
				Name: "delivery_s3_key_prefix",
				Type: schema.TypeString,
			},
			{
				Name: "last_update_requested_time",
				Type: schema.TypeTimestamp,
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
	for {
		resp, err := c.Services().ConfigService.DescribeConformancePacks(ctx, &config, func(options *configservice.Options) {
			options.Region = c.Region
		})
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
