// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func Apigatewayv2ApisAuthorizers() *schema.Table {

	return &schema.Table{
		Name:      "aws_apigatewayv2_authorizers",
		Resolver:  fetchApigatewayv2ApisAuthorizers,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "authorizer_credentials_arn",
				Type: schema.TypeString,
			},
			{
				Name: "authorizer_id",
				Type: schema.TypeString,
			},
			{
				Name: "authorizer_payload_format_version",
				Type: schema.TypeString,
			},
			{
				Name: "authorizer_result_ttl_in_seconds",
				Type: schema.TypeInt,
			},
			{
				Name: "authorizer_type",
				Type: schema.TypeString,
			},
			{
				Name: "authorizer_uri",
				Type: schema.TypeString,
			},
			{
				Name: "enable_simple_responses",
				Type: schema.TypeBool,
			},
			{
				Name: "identity_source",
				Type: schema.TypeStringArray,
			},
			{
				Name: "identity_validation_expression",
				Type: schema.TypeString,
			},
			{
				Name: "jwt_configuration",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchApigatewayv2ApisAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r := parent.Item.(types.Api)
	input := apigatewayv2.GetAuthorizersInput{
		ApiId: r.ApiId,
	}

	for {
		response, err := svc.GetAuthorizers(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
