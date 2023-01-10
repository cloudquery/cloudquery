package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIotThings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingsInput{
		MaxResults: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	for {
		response, err := svc.ListThings(ctx, &input)
		if err != nil {
			return err
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
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListThingPrincipalsInput{
		ThingName:  i.ThingName,
		MaxResults: aws.Int32(250),
	}
	var principals []string

	for {
		response, err := svc.ListThingPrincipals(ctx, &input)

		if err != nil {
			return err
		}
		principals = append(principals, response.Principals...)

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return resource.Set(c.Name, principals)
}
