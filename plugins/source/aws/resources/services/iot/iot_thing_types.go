package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotThingTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_thing_types",
		Description: "The definition of the thing type, including thing type name and description.",
		Resolver:    fetchIotThingTypes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:        "tags",
				Description: "Tags of the resource",
				Type:        schema.TypeJSON,
				Resolver:    ResolveIotThingTypeTags,
			},
			{
				Name:            "arn",
				Description:     "The thing type ARN.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ThingTypeArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "thing_type_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ThingTypeMetadata"),
			},
			{
				Name:        "name",
				Description: "The name of the thing type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ThingTypeName"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotThingTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListThingTypesInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListThingTypes(ctx, &input)
		if err != nil {
			return err
		}

		res <- response.ThingTypes

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
func ResolveIotThingTypeTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(types.ThingTypeDefinition)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListTagsForResourceInput{
		ResourceArn: i.ThingTypeArn,
	}
	tags := make(map[string]string)

	for {
		response, err := svc.ListTagsForResource(ctx, &input)

		if err != nil {
			return err
		}

		client.TagsIntoMap(response.Tags, tags)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, tags)
}
