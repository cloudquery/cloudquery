package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Apigatewayv2VpcLinks() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_vpc_links",
		Resolver:     fetchApigatewayv2VpcLinks,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "security_group_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "subnet_ids",
				Type: schema.TypeStringArray,
			},
			{
				Name: "vpc_link_id",
				Type: schema.TypeString,
			},
			{
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "vpc_link_status",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_link_status_message",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_link_version",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayv2VpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigatewayv2.GetVpcLinksInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetVpcLinks(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
		})

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
