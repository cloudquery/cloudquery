package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)


func Parameters() *schema.Table {
	return &schema.Table{
		Name:         "aws_ssm_parameters",
		Description:  "Metadata includes information like the ARN of the last user and the date/time the parameter was last used",
		Resolver:     fetchSsmParameters,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ssm"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Type:        schema.TypeInt,
			},
			{
				Name:        "policies",
				Description: "A list of policies assigned to a parameter",
				Type:        schema.TypeJSON,
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
			return err
		}
		res <- output.Parameters
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
