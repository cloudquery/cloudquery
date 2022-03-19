package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotThings() *schema.Table {
	return &schema.Table{
		Name:         "aws_iot_things",
		Description:  "The properties of the thing, including thing name, thing type name, and a list of thing attributes.",
		Resolver:     fetchIotThings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:          "principals",
				Description:   "Principals associated with the thing",
				Type:          schema.TypeStringArray,
				Resolver:      ResolveIotThingPrincipals,
				IgnoreInTests: true,
			},
			{
				Name:        "attributes",
				Description: "A list of thing attributes which are name-value pairs.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "arn",
				Description: "The thing ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingArn"),
			},
			{
				Name:        "name",
				Description: "The name of the thing.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingName"),
			},
			{
				Name:        "type_name",
				Description: "The name of the thing type, if the thing has been associated with a type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingTypeName"),
			},
			{
				Name:        "version",
				Description: "The version of the thing record in the registry.",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotThings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListThingsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThings(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Things
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotThingPrincipals(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(types.ThingAttribute)
	client := meta.(*client.Client)
	svc := client.Services().IOT
	input := iot.ListThingPrincipalsInput{
		ThingName:  i.ThingName,
		MaxResults: aws.Int32(250),
	}
	var principals []string

	for {
		response, err := svc.ListThingPrincipals(ctx, &input, func(options *iot.Options) {
			options.Region = client.Region
		})

		if err != nil {
			return diag.WrapError(err)
		}
		principals = append(principals, response.Principals...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, principals)
}
