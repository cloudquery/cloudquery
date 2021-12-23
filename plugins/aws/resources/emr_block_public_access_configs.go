package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func EmrBlockPublicAccessConfigs() *schema.Table {
	return &schema.Table{
		Name:         "aws_emr_block_public_access_configs",
		Resolver:     fetchEmrBlockPublicAccessConfigs,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "block_public_security_group_rules",
				Description: "Indicates whether Amazon EMR block public access is enabled or disabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BlockPublicAccessConfiguration.BlockPublicSecurityGroupRules"),
			},
			{
				Name:        "classification",
				Description: "The classification within a configuration.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BlockPublicAccessConfiguration.Classification"),
			},
			{
				Name:        "configurations",
				Description: "A list of additional configurations to apply within a configuration object.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEmrBlockPublicAccessConfigConfigurations,
			},
			{
				Name:        "properties",
				Description: "A set of properties specified within a configuration classification.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("BlockPublicAccessConfiguration.Properties"),
			},
			{
				Name:        "created_by_arn",
				Description: "The Amazon Resource Name that created or last modified the configuration.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BlockPublicAccessConfigurationMetadata.CreatedByArn"),
			},
			{
				Name:        "creation_date_time",
				Description: "The date and time that the configuration was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("BlockPublicAccessConfigurationMetadata.CreationDateTime"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_emr_block_public_access_config_port_ranges",
				Description: "A list of port ranges that are permitted to allow inbound traffic from all public IP addresses",
				Resolver:    fetchEmrBlockPublicAccessConfigPermittedPublicSecurityGroupRuleRanges,
				Columns: []schema.Column{
					{
						Name:        "block_public_access_config_cq_id",
						Description: "Unique CloudQuery ID of aws_emr_block_public_access_configs table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "min_range",
						Description: "The smallest port number in a specified range of port numbers.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "max_range",
						Description: "The smallest port number in a specified range of port numbers.",
						Type:        schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEmrBlockPublicAccessConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EMR
	out, err := svc.GetBlockPublicAccessConfiguration(ctx, &emr.GetBlockPublicAccessConfigurationInput{}, func(options *emr.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- out
	return nil
}

func resolveEmrBlockPublicAccessConfigConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	out, ok := resource.Item.(*emr.GetBlockPublicAccessConfigurationOutput)
	if !ok {
		return fmt.Errorf("not an *emr.GetBlockPublicAccessConfigurationOutput: %T", resource.Item)
	}
	if out.BlockPublicAccessConfiguration == nil {
		return nil
	}
	b, err := json.Marshal(out.BlockPublicAccessConfiguration.Configurations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, b)
}

func fetchEmrBlockPublicAccessConfigPermittedPublicSecurityGroupRuleRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	out, ok := parent.Item.(*emr.GetBlockPublicAccessConfigurationOutput)
	if !ok {
		return fmt.Errorf("not an *emr.GetBlockPublicAccessConfigurationOutput: %T", parent.Item)
	}
	if out.BlockPublicAccessConfiguration == nil {
		return nil
	}
	res <- out.BlockPublicAccessConfiguration.PermittedPublicSecurityGroupRuleRanges
	return nil
}
