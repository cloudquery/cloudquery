package inspector2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Findings() *schema.Table {
	return &schema.Table{
		Name:        "aws_inspector2_findings",
		Description: "Details about an Amazon Inspector finding",
		Resolver:    fetchInspector2Findings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("inspector2"),
		Columns: []schema.Column{
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the finding",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("FindingArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:     "remediation",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Remediation"),
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
			{
				Name:        "resources",
				Description: "Details about the resource involved in a finding",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Resources"),
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
			return err
		}
		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
