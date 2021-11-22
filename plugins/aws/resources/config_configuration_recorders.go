package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/configservice/types"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ConfigConfigurationRecorders() *schema.Table {
	return &schema.Table{
		Name:         "aws_config_configuration_recorders",
		Description:  "An object that represents the recording of configuration changes of an AWS resource.",
		Resolver:     fetchConfigConfigurationRecorders,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Description: "Amazon Resource Name (ARN) of the config recorder.",
				Type:        schema.TypeString,
				Resolver:    generateConfigRecorderArn,
			},
			{
				Name:        "name",
				Description: "The name of the recorder.",
				Type:        schema.TypeString,
			},
			{
				Name:        "recording_group_all_supported",
				Description: "Specifies whether AWS Config records configuration changes for every supported type of regional resource.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RecordingGroup.AllSupported"),
			},
			{
				Name:        "recording_group_include_global_resource_types",
				Description: "Specifies whether AWS Config includes all supported types of global resources (for example, IAM resources) with the resources that it records.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("RecordingGroup.IncludeGlobalResourceTypes"),
			},
			{
				Name:        "recording_group_resource_types",
				Description: "A comma-separated list that specifies the types of AWS resources for which AWS Config records configuration changes (for example, AWS::EC2::Instance or AWS::CloudTrail::Trail).",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("RecordingGroup.ResourceTypes"),
			},
			{
				Name:        "role_arn",
				Description: "Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources associated with the account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleARN"),
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

func generateConfigRecorderArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cfg, ok := resource.Item.(types.ConfigurationRecorder)
	if !ok {
		return fmt.Errorf("not config config recorder")
	}
	return resource.Set(c.Name, client.GenerateResourceARN("config", "config-recorder", *cfg.Name, cl.Region, cl.AccountID))

}
