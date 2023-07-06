package mq

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func brokerConfigurations() *schema.Table {
	tableName := "aws_mq_broker_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id.html`,
		Resolver:    fetchMqBrokerConfigurations,
		Transform:   transformers.TransformWithStruct(&mq.DescribeConfigurationOutput{}, transformers.WithSkipFields("ResultMetadata"), transformers.WithPrimaryKeys("Arn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "broker_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},

		Relations: []*schema.Table{
			brokerConfigurationRevisions(),
		},
	}
}

func fetchMqBrokerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	// Ensure Configurations is not nil
	// This *might* occur during initial creation of broker
	if broker.Configurations == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Mq
	list := broker.Configurations.History
	if broker.Configurations.Current != nil {
		list = append(list, *broker.Configurations.Current)
	}

	// History might contain same Id multiple times (maybe with different revisions) but we're only interested in the latest revision of each
	dupes := make(map[string]struct{}, len(list))
	configurations := make([]mq.DescribeConfigurationOutput, 0, len(list))
	for _, cfg := range list {
		if cfg.Id == nil {
			continue
		}

		if _, ok := dupes[*cfg.Id]; ok {
			continue
		}
		dupes[*cfg.Id] = struct{}{}

		input := mq.DescribeConfigurationInput{ConfigurationId: cfg.Id}
		output, err := svc.DescribeConfiguration(ctx, &input, func(options *mq.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		configurations = append(configurations, *output)
	}
	res <- configurations
	return nil
}
