package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource findings --config gen.hcl --output .
func Findings() *schema.Table {
	return &schema.Table{
		Name:         "aws_inspector_findings",
		Description:  "Contains information about an Amazon Inspector finding",
		Resolver:     fetchInspectorFindings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("inspector"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The ARN that specifies the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "attributes",
				Description: "The system-defined attributes for the finding",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTagField("Attributes"),
			},
			{
				Name:        "created_at",
				Description: "The time when the finding was generated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "updated_at",
				Description: "The time when AddAttributesToFindings is called",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "user_attributes",
				Description: "The user-defined attributes that are assigned to the finding",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTagField("UserAttributes"),
			},
			{
				Name:        "asset_attributes",
				Description: "A collection of attributes of the host from which the finding is generated",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "asset_type",
				Description: "The type of the host from which the finding is generated",
				Type:        schema.TypeString,
			},
			{
				Name:        "confidence",
				Description: "This data element is currently not used",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "description",
				Description: "The description of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "indicator_of_compromise",
				Description: "This data element is currently not used",
				Type:        schema.TypeBool,
			},
			{
				Name:        "numeric_severity",
				Description: "The numeric value of the finding severity",
				Type:        schema.TypeFloat,
			},
			{
				Name:        "recommendation",
				Description: "The recommendation for the finding",
				Type:        schema.TypeString,
			},
			{
				Name:        "schema_version",
				Description: "The schema version of this data type",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "service",
				Description: "The data element is set to \"Inspector\"",
				Type:        schema.TypeString,
			},
			{
				Name:        "service_attributes",
				Description: "This data type is used in the Finding data type",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "severity",
				Description: "The finding severity",
				Type:        schema.TypeString,
			},
			{
				Name:        "title",
				Description: "The name of the finding",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchInspectorFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Inspector
	input := inspector.ListFindingsInput{MaxResults: aws.Int32(100)}
	for {
		response, err := svc.ListFindings(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		if len(response.FindingArns) > 0 {
			out, err := svc.DescribeFindings(ctx, &inspector.DescribeFindingsInput{FindingArns: response.FindingArns})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- out.Findings
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
