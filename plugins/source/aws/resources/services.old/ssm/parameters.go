package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource parameters --config resources/services/ssm/parameters.hcl --output .
func Parameters() *schema.Table {
	return &schema.Table{
		Name:         "aws_ssm_parameters",
		Description:  "Metadata includes information like the ARN of the last user and the date/time the parameter was last used",
		Resolver:     fetchSsmParameters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ssm"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
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
				Name:        "allowed_pattern",
				Description: "A parameter name can include only the following letters and symbols a-zA-Z0-9_-",
				Type:        schema.TypeString,
			},
			{
				Name:        "data_type",
				Description: "The data type of the parameter, such as text or aws:ec2:image",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "Description of the parameter actions",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_id",
				Description: "The ID of the query key used for this parameter",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_date",
				Description: "Date the parameter was last changed or updated",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "last_modified_user",
				Description: "Amazon Resource Name (ARN) of the Amazon Web Services user who last changed the parameter",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The parameter name",
				Type:        schema.TypeString,
			},
			{
				Name:        "tier",
				Description: "The parameter tier",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of parameter",
				Type:        schema.TypeString,
			},
			{
				Name:        "version",
				Description: "The parameter version",
				Type:        schema.TypeBigInt,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ssm_parameter_policies",
				Description: "One or more policies assigned to a parameter",
				Resolver:    schema.PathTableResolver("Policies"),
				Columns: []schema.Column{
					{
						Name:        "parameter_cq_id",
						Description: "Unique CloudQuery ID of aws_ssm_parameters table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "policy_status",
						Description: "The status of the policy",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_text",
						Description: "The JSON text of the policy",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_type",
						Description: "The type of policy",
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

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().SSM
	params := ssm.DescribeParametersInput{}
	for {
		output, err := svc.DescribeParameters(ctx, &params)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Parameters
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
