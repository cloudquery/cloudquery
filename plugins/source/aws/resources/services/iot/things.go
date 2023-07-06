package iot

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Things() *schema.Table {
	tableName := "aws_iot_things"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_ThingAttribute.html`,
		Resolver:    fetchIotThings,
		Transform:   transformers.TransformWithStruct(&types.ThingAttribute{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "principals",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: ResolveIotThingPrincipals,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ThingArn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchIotThings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListThingsInput{
		MaxResults: aws.Int32(250),
	}
	cl := meta.(*client.Client)

	svc := cl.Services().Iot
	paginator := iot.NewListThingsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Things
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
	paginator := iot.NewListThingPrincipalsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})

		if err != nil {
			return err
		}
		principals = append(principals, page.Principals...)
	}
	return resource.Set(c.Name, principals)
}
