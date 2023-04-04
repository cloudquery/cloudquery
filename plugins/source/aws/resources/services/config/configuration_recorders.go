package config

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/config/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ConfigurationRecorders() *schema.Table {
	tableName := "aws_config_configuration_recorders"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigurationRecorder.html`,
		Resolver:    fetchConfigConfigurationRecorders,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "config"),
		Transform:   transformers.TransformWithStruct(&models.ConfigurationRecorderWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: generateConfigRecorderArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchConfigConfigurationRecorders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	resp, err := c.Services().Configservice.DescribeConfigurationRecorders(ctx, &configservice.DescribeConfigurationRecordersInput{})
	if err != nil {
		return err
	}
	if len(resp.ConfigurationRecorders) == 0 {
		return nil
	}
	names := make([]string, len(resp.ConfigurationRecorders))
	for i, configurationRecorder := range resp.ConfigurationRecorders {
		names[i] = *configurationRecorder.Name
	}
	status, err := c.Services().Configservice.DescribeConfigurationRecorderStatus(ctx, &configservice.DescribeConfigurationRecorderStatusInput{
		ConfigurationRecorderNames: names,
	})
	if err != nil {
		return err
	}
	for _, configurationRecorder := range resp.ConfigurationRecorders {
		if configurationRecorder.Name == nil {
			continue
		}
		var configurationRecorderStatus types.ConfigurationRecorderStatus
		for _, s := range status.ConfigurationRecordersStatus {
			if s.Name == nil {
				continue
			}
			if *s.Name == *configurationRecorder.Name {
				configurationRecorderStatus = s
				res <- models.ConfigurationRecorderWrapper{
					ConfigurationRecorder:      configurationRecorder,
					StatusLastErrorCode:        configurationRecorderStatus.LastErrorCode,
					StatusLastErrorMessage:     configurationRecorderStatus.LastErrorMessage,
					StatusLastStartTime:        configurationRecorderStatus.LastStartTime,
					StatusLastStatus:           configurationRecorderStatus.LastStatus,
					StatusLastStatusChangeTime: configurationRecorderStatus.LastStatusChangeTime,
					StatusLastStopTime:         configurationRecorderStatus.LastStopTime,
					StatusRecording:            configurationRecorderStatus.Recording,
				}

				break
			}
		}
	}
	return nil
}

func generateConfigRecorderArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cfg := resource.Item.(models.ConfigurationRecorderWrapper)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "config",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("config-recorder/%s", aws.ToString(cfg.Name)),
	}.String())
}
