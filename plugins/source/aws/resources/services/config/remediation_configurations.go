package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func remediationConfigurations() *schema.Table {
	tableName := "aws_config_remediation_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_RemediationConfiguration.html`,
		Resolver:    fetchRemediationConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform: transformers.TransformWithStruct(&types.RemediationConfiguration{},
			transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
		Relations: []*schema.Table{},
	}
}

func fetchRemediationConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Configservice

	configRule := parent.Item.(types.ConfigRule).ConfigRuleName
	input := &configservice.DescribeRemediationConfigurationsInput{
		ConfigRuleNames: []string{*configRule},
	}

	// no pagination for this one
	output, err := svc.DescribeRemediationConfigurations(ctx, input, func(options *configservice.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.RemediationConfigurations

	return nil
}
