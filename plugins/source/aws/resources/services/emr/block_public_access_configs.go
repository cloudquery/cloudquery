package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EmrBlockPublicAccessConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_emr_block_public_access_configs",
		Resolver:  fetchEmrBlockPublicAccessConfigs,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticmapreduce"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "block_public_security_group_rules",
				Description: "Indicates whether Amazon EMR block public access is enabled or disabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("BlockPublicAccessConfiguration.BlockPublicSecurityGroupRules"),
			},
			{
				Name:          "classification",
				Description:   "The classification within a configuration.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("BlockPublicAccessConfiguration.Classification"),
				IgnoreInTests: true,
			},
			{
				Name:          "configurations",
				Description:   "A list of additional configurations to apply within a configuration object.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("BlockPublicAccessConfiguration.Configurations"),
				IgnoreInTests: true,
			},
			{
				Name:          "properties",
				Description:   "A set of properties specified within a configuration classification.",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("BlockPublicAccessConfiguration.Properties"),
				IgnoreInTests: true,
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
			{
				Name:     "block_public_access_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockPublicAccessConfiguration"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEmrBlockPublicAccessConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EMR
	out, err := svc.GetBlockPublicAccessConfiguration(ctx, &emr.GetBlockPublicAccessConfigurationInput{})
	if err != nil {
		if client.IgnoreNotAvailableRegion(err) {
			meta.Logger().Debug().Err(err).Msg("received InvalidRequestException on GetBlockPublicAccessConfiguration, api is not available in the current Region.")
			return nil
		}
		return err
	}
	res <- out
	return nil
}
