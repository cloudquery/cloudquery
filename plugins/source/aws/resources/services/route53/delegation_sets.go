package route53

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DelegationSets() *schema.Table {
	tableName := "aws_route53_delegation_sets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/Route53/latest/APIReference/API_DelegationSet.html`,
		Resolver:    fetchRoute53DelegationSets,
		Transform:   transformers.TransformWithStruct(&types.DelegationSet{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:        "arn",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveDelegationSetArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				PrimaryKey:  true,
			},
		},
	}
}

func fetchRoute53DelegationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config route53.ListReusableDelegationSetsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Route53
	// no paginator available
	for {
		response, err := svc.ListReusableDelegationSets(ctx, &config, func(options *route53.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.DelegationSets
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveDelegationSetArn() schema.ColumnResolver {
	return client.ResolveARNGlobal(client.Route53Service, func(resource *schema.Resource) ([]string, error) {
		return []string{"delegationset", *resource.Item.(types.DelegationSet).Id}, nil
	})
}
