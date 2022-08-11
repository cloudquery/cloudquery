package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource alarms --config gen.hcl --output .
func Alarms() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_alarms",
		Description:   "Describes an alarm",
		Resolver:      fetchLightsailAlarms,
		Multiplex:     client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the alarm",
				Type:        schema.TypeString,
			},
			{
				Name:        "comparison_operator",
				Description: "The arithmetic operation used when comparing the specified statistic and threshold",
				Type:        schema.TypeString,
			},
			{
				Name:        "contact_protocols",
				Description: "The contact protocols for the alarm, such as Email, SMS (text messaging), or both",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the alarm was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "datapoints_to_alarm",
				Description: "The number of data points that must not within the specified threshold to trigger the alarm",
				Type:        schema.TypeInt,
			},
			{
				Name:        "evaluation_periods",
				Description: "The number of periods over which data is compared to the specified threshold",
				Type:        schema.TypeInt,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "metric_name",
				Description: "The name of the metric associated with the alarm",
				Type:        schema.TypeString,
			},
			{
				Name:        "monitored_resource_info_arn",
				Description: "The Amazon Resource Name (ARN) of the resource being monitored",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MonitoredResourceInfo.Arn"),
			},
			{
				Name:        "monitored_resource_name",
				Description: "The name of the Lightsail resource being monitored",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MonitoredResourceInfo.Name"),
			},
			{
				Name:        "monitored_resource_resource_type",
				Description: "The Lightsail resource type of the resource being monitored",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MonitoredResourceInfo.ResourceType"),
			},
			{
				Name:        "name",
				Description: "The name of the alarm",
				Type:        schema.TypeString,
			},
			{
				Name:        "notification_enabled",
				Description: "Indicates whether the alarm is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "notification_triggers",
				Description: "The alarm states that trigger a notification",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "period",
				Description: "The period, in seconds, over which the statistic is applied",
				Type:        schema.TypeInt,
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type (eg, Alarm)",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the alarm",
				Type:        schema.TypeString,
			},
			{
				Name:        "statistic",
				Description: "The statistic for the metric associated with the alarm",
				Type:        schema.TypeString,
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "threshold",
				Description: "The value against which the specified statistic is compared",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "treat_missing_data",
				Description: "Specifies how the alarm handles missing data points",
				Type:        schema.TypeString,
			},
			{
				Name:        "unit",
				Description: "The unit of the metric associated with the alarm",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailAlarms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetAlarmsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetAlarms(ctx, &input, func(options *lightsail.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Alarms
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
