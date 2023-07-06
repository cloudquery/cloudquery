package lightsail

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Instances() *schema.Table {
	tableName := "aws_lightsail_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Instance.html`,
		Resolver:    fetchLightsailInstances,
		Transform:   transformers.TransformWithStruct(&types.Instance{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "access_details",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveLightsailInstanceAccessDetails,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			instancePortStates(),
		},
	}
}

func fetchLightsailInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	input := lightsail.GetInstancesInput{}
	// No paginator available
	for {
		output, err := svc.GetInstances(ctx, &input, func(options *lightsail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Instances

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		input.PageToken = output.NextPageToken
	}
	return nil
}
func resolveLightsailInstanceAccessDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstanceAccessDetailsInput{InstanceName: r.Name}
	output, err := svc.GetInstanceAccessDetails(ctx, &input, func(options *lightsail.Options) {
		options.Region = cli.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, output.AccessDetails)
}
