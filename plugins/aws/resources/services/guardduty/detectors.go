package guardduty

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func GuarddutyDetectors() *schema.Table {
	return &schema.Table{
		Name:         "aws_guardduty_detectors",
		Resolver:     fetchGuarddutyDetectors,
		Multiplex:    client.ServiceAccountRegionMultiplexer("guardduty"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "id"}},
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.GuardDutyService, func(resource *schema.Resource) ([]string, error) {
					return []string{"detector", resource.Item.(Detector).Id}, nil
				}),
			},
			{
				Name:        "id",
				Description: "The Unique Identifier of the Detector.",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_role",
				Description: "The GuardDuty service role.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The detector status.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp of when the detector was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreatedAt"),
			},
			{
				Name:        "data_sources_cloud_trail_status",
				Description: "Describes whether CloudTrail is enabled as a data source for the detector.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataSources.CloudTrail.Status"),
			},
			{
				Name:        "data_sources_dns_logs_status",
				Description: "Denotes whether DNS logs is enabled as a data source.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataSources.DNSLogs.Status"),
			},
			{
				Name:        "data_sources_flow_logs_status",
				Description: "Denotes whether VPC flow logs is enabled as a data source.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataSources.FlowLogs.Status"),
			},
			{
				Name:        "data_sources_s3_logs_status",
				Description: "A value that describes whether S3 data event logs are automatically enabled for new members of the organization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DataSources.S3Logs.Status"),
			},
			{
				Name:        "finding_publishing_frequency",
				Description: "The publishing frequency of the finding.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags of the detector resource.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "updated_at",
				Description: "The last-updated timestamp for the detector.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("UpdatedAt"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_guardduty_detector_members",
				Description: "Contains information about the member account.",
				Resolver:    fetchGuarddutyDetectorMembers,
				Columns: []schema.Column{
					{
						Name:        "detector_cq_id",
						Description: "Unique CloudQuery ID of aws_guardduty_detectors table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "account_id",
						Description: "The ID of the member account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "email",
						Description: "The email address of the member account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "master_id",
						Description: "The administrator account ID.",
						Type:        schema.TypeString,
					},
					{
						Name:        "relationship_status",
						Description: "The status of the relationship between the member and the administrator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "updated_at",
						Description: "The last-updated timestamp of the member.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("UpdatedAt"),
					},
					{
						Name:        "detector_id",
						Description: "The detector ID of the member account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "invited_at",
						Description: "The timestamp when the invitation was sent.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.DateResolver("InvitedAt"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGuarddutyDetectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().GuardDuty
	config := &guardduty.ListDetectorsInput{}
	for {
		output, err := svc.ListDetectors(ctx, config, func(o *guardduty.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, dId := range output.DetectorIds {
			d, err := svc.GetDetector(ctx, &guardduty.GetDetectorInput{DetectorId: aws.String(dId)}, func(o *guardduty.Options) {
				o.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- Detector{d, dId}
		}
		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

func fetchGuarddutyDetectorMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	detector, ok := parent.Item.(Detector)
	if !ok {
		return fmt.Errorf("expected type *guardduty.GetDetectorOutput got %T", parent.Item)
	}
	c := meta.(*client.Client)
	svc := c.Services().GuardDuty
	config := &guardduty.ListMembersInput{DetectorId: aws.String(detector.Id)}
	for {
		output, err := svc.ListMembers(ctx, config, func(options *guardduty.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- output.Members
		if output.NextToken == nil {
			return nil
		}
		config.NextToken = output.NextToken
	}
}

type Detector struct {
	*guardduty.GetDetectorOutput
	Id string
}
