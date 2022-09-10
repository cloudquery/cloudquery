package config

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type configurationRecorderWrapper struct {
	types.ConfigurationRecorder
	StatusLastErrorCode        *string
	StatusLastErrorMessage     *string
	StatusLastStartTime        *time.Time
	StatusLastStatus           types.RecorderStatus
	StatusLastStatusChangeTime *time.Time
	StatusLastStopTime         *time.Time
	StatusRecording            bool
}

func ConfigConfigurationRecorders() *schema.Table {
	return &schema.Table{
		Name:          "aws_config_configuration_recorders",
		Description:   "An object that represents the recording of configuration changes of an AWS resource.",
		Resolver:      fetchConfigConfigurationRecorders,
		Multiplex:     client.ServiceAccountRegionMultiplexer("config"),
		IgnoreInTests: true,
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
				Name:            "arn",
				Description:     "Amazon Resource Name (ARN) of the config recorder.",
				Type:            schema.TypeString,
				Resolver:        generateConfigRecorderArn,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of the recorder.",
				Type:        schema.TypeString,
			},
			{
				Name:     "recording_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RecordingGroup"),
			},
			{
				Name:        "role_arn",
				Description: "Amazon Resource Name (ARN) of the IAM role used to describe the AWS resources associated with the account.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RoleARN"),
			},
			{
				Name:        "status_last_error_code",
				Description: "The error code indicating that the recording failed.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_last_error_message",
				Description: "The message indicating that the recording failed due to an error.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_last_start_time",
				Description: "The time the recorder was last started.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status_last_status",
				Description: "The last (previous) status of the recorder.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status_last_status_change_time",
				Description: "The time when the status was last changed.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status_last_stop_time",
				Description: "The time the recorder was last stopped.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status_recording",
				Description: "Specifies whether or not the recorder is currently recording.",
				Type:        schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchConfigConfigurationRecorders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	resp, err := c.Services().ConfigService.DescribeConfigurationRecorders(ctx, &configservice.DescribeConfigurationRecordersInput{})
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
	status, err := c.Services().ConfigService.DescribeConfigurationRecorderStatus(ctx, &configservice.DescribeConfigurationRecorderStatusInput{
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
				res <- configurationRecorderWrapper{
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
	cfg := resource.Item.(configurationRecorderWrapper)
	return resource.Set(c.Name, cl.ARN("config", "config-recorder", *cfg.Name))
}
