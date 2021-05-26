package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ConfigConfigurationRecorders() *schema.Table {
	return &schema.Table{
		Name:         "aws_config_configuration_recorders",
		Resolver:     fetchConfigConfigurationRecorders,
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
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name:     "recording_group_all_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RecordingGroup.AllSupported"),
			},
			{
				Name:     "recording_group_include_global_resource_types",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("RecordingGroup.IncludeGlobalResourceTypes"),
			},
			{
				Name:     "recording_group_resource_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("RecordingGroup.ResourceTypes"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleARN"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchConfigConfigurationRecorders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	resp, err := c.Services().ConfigService.DescribeConfigurationRecorders(ctx, &configservice.DescribeConfigurationRecordersInput{}, func(options *configservice.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- resp.ConfigurationRecorders
	return nil
}
