package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource findings --config gen.hcl --output .
func Findings() *schema.Table {
	return &schema.Table{
		Name:         "aws_inspector2_findings",
		Description:  "Details about an Amazon Inspector finding",
		Resolver:     fetchInspector2Findings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("inspector2"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the finding",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("FindingArn"),
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "account_id",
				Description: "The Amazon Web Services account ID associated with the finding",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("AwsAccountId"),
			},
			{
				Name:        "description",
				Description: "The description of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "finding_arn",
				Description: "The Amazon Resource Number (ARN) of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "first_observed_at",
				Description: "The date and time that the finding was first observed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_observed_at",
				Description: "The date and time that the finding was last observed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "remediation_recommendation_text",
				Description: "The recommended course of action to remediate the finding",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Remediation.Recommendation.Text"),
			},
			{
				Name:        "remediation_recommendation_url",
				Description: "The URL address to the CVE remediation recommendations",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Remediation.Recommendation.Url"),
			},
			{
				Name:        "severity",
				Description: "The severity of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "inspector_score",
				Description: "The Amazon Inspector score given to the finding",
				Type:        schema.TypeFloat,
			},
			{
				Name:          "inspector_score_details",
				Description:   "An object that contains details of the Amazon Inspector score",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "network_reachability_details",
				Description:   "An object that contains the details of a network reachability finding",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:          "package_vulnerability_details",
				Description:   "An object that contains the details of a package vulnerability finding",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
			{
				Name:        "title",
				Description: "The title of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "updated_at",
				Description: "The date and time the finding was last updated at",
				Type:        schema.TypeTimestamp,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_inspector2_finding_resources",
				Description: "Details about the resource involved in a finding",
				Resolver:    schema.PathTableResolver("Resources"),
				Columns: []schema.Column{
					{
						Name:        "finding_cq_id",
						Description: "Unique CloudQuery ID of aws_inspector2_findings table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "The ID of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "The type of resource",
						Type:        schema.TypeString,
					},
					{
						Name:          "aws_ec2_instance",
						Description:   "An object that contains details about the Amazon EC2 instance involved in the finding",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Details.AwsEc2Instance"),
						IgnoreInTests: true,
					},
					{
						Name:          "aws_ecr_container_image",
						Description:   "An object that contains details about the Amazon ECR container image involved in the finding",
						Type:          schema.TypeJSON,
						Resolver:      schema.PathResolver("Details.AwsEcrContainerImage"),
						IgnoreInTests: true,
					},
					{
						Name:        "partition",
						Description: "The partition of the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "region",
						Description: "The Amazon Web Services Region the impacted resource is located in",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The tags attached to the resource",
						Type:        schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchInspector2Findings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().InspectorV2
	input := inspector2.ListFindingsInput{MaxResults: aws.Int32(100)}
	for {
		response, err := svc.ListFindings(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
