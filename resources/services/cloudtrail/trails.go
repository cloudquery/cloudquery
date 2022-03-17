package cloudtrail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudtrailTrails() *schema.Table {
	return &schema.Table{
		Name:                 "aws_cloudtrail_trails",
		Description:          "The settings for a trail.",
		Resolver:             fetchCloudtrailTrails,
		Multiplex:            client.AccountMultiplex,
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountFilter,
		PostResourceResolver: postCloudtrailTrailResolver,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "arn"}},
		IgnoreInTests:        true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Description: "Any tags assigned to the resource",
			},
			{
				Name:     "cloudwatch_logs_log_group_name",
				Type:     schema.TypeString,
				Resolver: resolveCloudtrailTrailCloudwatchLogsLogGroupName,
			},
			{
				Name:        "is_logging",
				Description: "Whether the CloudTrail is currently logging AWS API calls.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "latest_cloud_watch_logs_delivery_error",
				Description: "Displays any CloudWatch Logs error that CloudTrail encountered when attempting to deliver logs to CloudWatch Logs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_cloud_watch_logs_delivery_time",
				Description: "Displays the most recent date and time when CloudTrail delivered logs to CloudWatch Logs.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "latest_delivery_error",
				Description: "Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver log files to the designated bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_delivery_time",
				Description: "Specifies the date and time that CloudTrail last delivered log files to an account's Amazon S3 bucket.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "latest_digest_delivery_error",
				Description: "Displays any Amazon S3 error that CloudTrail encountered when attempting to deliver a digest file to the designated bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_digest_delivery_time",
				Description: "Specifies the date and time that CloudTrail last delivered a digest file to an account's Amazon S3 bucket.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "latest_notification_error",
				Description: " Displays any Amazon SNS error that CloudTrail encountered when attempting to send a notification.",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_notification_time",
				Description: "Specifies the date and time of the most recent Amazon SNS notification that CloudTrail has written a new log file to an account's Amazon S3 bucket.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "start_logging_time",
				Description: "Specifies the most recent date and time when CloudTrail started recording API calls for an AWS account.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "stop_logging_time",
				Description: "Specifies the most recent date and time when CloudTrail stopped recording API calls for an AWS account.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "cloud_watch_logs_log_group_arn",
				Description: "Specifies an Amazon Resource Name (ARN), a unique identifier that represents the log group to which CloudTrail logs will be delivered.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cloud_watch_logs_role_arn",
				Description: "Specifies the role for the CloudWatch Logs endpoint to assume to write to a user's log group.",
				Type:        schema.TypeString,
			},
			{
				Name:        "has_custom_event_selectors",
				Description: "Specifies if the trail has custom event selectors.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "has_insight_selectors",
				Description: "Specifies whether a trail has insight types specified in an InsightSelector list.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "region",
				Description: "The region in which the trail was created.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HomeRegion"),
			},
			{
				Name:        "include_global_service_events",
				Description: "Set to True to include AWS API calls from AWS global services such as IAM. Otherwise, False.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "is_multi_region_trail",
				Description: "Specifies whether the trail exists only in one region or exists in all regions.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "is_organization_trail",
				Description: "Specifies whether the trail is an organization trail.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "kms_key_id",
				Description: "Specifies the KMS key ID that encrypts the logs delivered by CloudTrail",
				Type:        schema.TypeString,
			},
			{
				Name:        "log_file_validation_enabled",
				Description: "Specifies whether log file validation is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "name",
				Description: "Name of the trail set by calling CreateTrail",
				Type:        schema.TypeString,
			},
			{
				Name:        "s3_bucket_name",
				Description: "Name of the Amazon S3 bucket into which CloudTrail delivers your trail files. See Amazon S3 Bucket Naming Requirements (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/create_trail_naming_policy.html).",
				Type:        schema.TypeString,
			},
			{
				Name:        "s3_key_prefix",
				Description: "Specifies the Amazon S3 key prefix that comes after the name of the bucket you have designated for log file delivery",
				Type:        schema.TypeString,
			},
			{
				Name:        "sns_topic_arn",
				Description: "Specifies the ARN of the Amazon SNS topic that CloudTrail uses to send notifications when log files are delivered",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SnsTopicARN"),
			},
			{
				Name:        "sns_topic_name",
				Description: "This field is no longer in use",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "Specifies the ARN of the trail",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TrailARN"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cloudtrail_trail_event_selectors",
				Description:   "Use event selectors to further specify the management and data event settings for your trail",
				Resolver:      fetchCloudtrailTrailEventSelectors,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "trail_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudtrail_trails table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "trail_arn",
						Description: "Specifies the ARN of the trail",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("TrailARN"),
					},
					{
						Name:        "exclude_management_event_sources",
						Description: "An optional list of service event sources from which you do not want management events to be logged on your trail",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "include_management_events",
						Description: "Specify if you want your event selector to include management events for your trail",
						Type:        schema.TypeBool,
					},
					{
						Name:        "read_write_type",
						Description: "Specify if you want your trail to log read-only events, write-only events, or all",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudtrailTrails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	log := meta.(*client.Client).Logger()
	response, err := svc.DescribeTrails(ctx, nil, func(options *cloudtrail.Options) {
		options.Region = c.Region
	})

	if err != nil {
		return diag.WrapError(err)
	}

	getBundledTrailsWithTags := func(trails []types.Trail, region string) ([]CloudTrailWrapper, error) {
		processed := make([]CloudTrailWrapper, len(trails))

		input := cloudtrail.ListTagsInput{
			ResourceIdList: make([]string, 0, len(trails)),
		}

		for i, h := range trails {
			processed[i] = CloudTrailWrapper{
				Trail: h,
				Tags:  make(map[string]interface{}),
			}

			// Before fetching trail tags we have to check if the trail is organization trail
			// If the trail is organization trail and the account id is not matched with current account id
			// We skip, and not fetch the trail tags
			arnParts, err := arn.Parse(*h.TrailARN)
			if err != nil {
				log.Warn("cloud not parse cloudtrail ARN", "arn", *h.TrailARN)
				continue
			}
			if *h.IsOrganizationTrail && c.AccountID != arnParts.AccountID {
				log.Warn("the trail is an organization level trail, cloud not fetch tags", "arn", *h.TrailARN)
				continue
			}

			input.ResourceIdList = append(input.ResourceIdList, *h.TrailARN)
		}

		if len(input.ResourceIdList) == 0 {
			return processed, nil
		}

		for {
			response, err := svc.ListTags(ctx, &input, func(options *cloudtrail.Options) {
				options.Region = region
			})
			if err != nil {
				return nil, err
			}
			for _, tr := range processed {
				for _, t := range getCloudTrailTagsByResourceID(*tr.TrailARN, response.ResourceTagList) {
					tr.Tags[*t.Key] = t.Value
				}
			}
			if aws.ToString(response.NextToken) == "" {
				break
			}
			input.NextToken = response.NextToken
		}

		return processed, nil
	}

	// since api returns all the cloudtrails despite region we aggregate trails by region to get tags.
	aggregatedTrails, err := aggregateCloudTrails(response.TrailList)
	if err != nil {
		return diag.WrapError(err)
	}
	for region, trails := range aggregatedTrails {
		for i := 0; i < len(trails); i += 20 {
			end := i + 20

			if end > len(trails) {
				end = len(trails)
			}
			t := trails[i:end]
			processed, err := getBundledTrailsWithTags(t, region)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- processed
		}
	}

	return nil
}

func postCloudtrailTrailResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	r, ok := resource.Item.(CloudTrailWrapper)
	if !ok {
		return fmt.Errorf("expected CloudTrailWrapper but got %T", resource.Item)
	}
	response, err := svc.GetTrailStatus(ctx,
		&cloudtrail.GetTrailStatusInput{Name: r.TrailARN}, func(o *cloudtrail.Options) {
			o.Region = *r.HomeRegion
		})
	if err != nil {
		return diag.WrapError(err)
	}
	if err := resource.Set("is_logging", response.IsLogging); err != nil {
		return err
	}
	if err := resource.Set("latest_cloud_watch_logs_delivery_error", response.LatestCloudWatchLogsDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_cloud_watch_logs_delivery_time", response.LatestCloudWatchLogsDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_delivery_error", response.LatestDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_delivery_time", response.LatestDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_digest_delivery_error", response.LatestDigestDeliveryError); err != nil {
		return err
	}
	if err := resource.Set("latest_digest_delivery_time", response.LatestDigestDeliveryTime); err != nil {
		return err
	}
	if err := resource.Set("latest_notification_error", response.LatestNotificationError); err != nil {
		return err
	}
	if err := resource.Set("latest_notification_time", response.LatestNotificationTime); err != nil {
		return err
	}
	if err := resource.Set("start_logging_time", response.StartLoggingTime); err != nil {
		return err
	}
	if err := resource.Set("stop_logging_time", response.StopLoggingTime); err != nil {
		return err
	}
	return nil
}

func resolveCloudtrailTrailCloudwatchLogsLogGroupName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	groupName := ""
	log := meta.(*client.Client).Logger()
	r, ok := resource.Item.(CloudTrailWrapper)
	if !ok {
		return fmt.Errorf("expected CloudTrailWrapper but got %T", resource.Item)
	}
	if r.CloudWatchLogsLogGroupArn != nil {
		matches := client.GroupNameRegex.FindStringSubmatch(*r.CloudWatchLogsLogGroupArn)
		if len(matches) < 2 {
			log.Warn("CloudWatchLogsLogGroupARN doesn't fit standard regex", "arn", *r.CloudWatchLogsLogGroupArn)
		} else {
			groupName = matches[1]
		}
	} else {
		log.Info("CloudWatchLogsLogGroupARN is empty")
	}

	return resource.Set("cloudwatch_logs_log_group_name", groupName)
}

func fetchCloudtrailTrailEventSelectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(CloudTrailWrapper)
	if !ok {
		return fmt.Errorf("expected CloudTrailWrapper but got %T", parent.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudtrail
	response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
		options.Region = *r.HomeRegion
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- response.EventSelectors
	return nil
}

func getCloudTrailTagsByResourceID(id string, set []types.ResourceTag) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.TagsList
		}
	}
	return nil
}

func aggregateCloudTrails(trails []types.Trail) (map[string][]types.Trail, error) {
	resp := make(map[string][]types.Trail)
	for _, t := range trails {
		if t.HomeRegion == nil {
			return nil, fmt.Errorf("got cloudtrail with HomeRegion == nil")
		}
		resp[*t.HomeRegion] = append(resp[*t.HomeRegion], t)
	}
	return resp, nil
}

type CloudTrailWrapper struct {
	types.Trail
	Tags map[string]interface{}
}
