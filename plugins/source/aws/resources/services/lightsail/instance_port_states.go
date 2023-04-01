package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func instancePortStates() *schema.Table {
	tableName := "aws_lightsail_instance_port_states"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html`,
		Resolver:    fetchLightsailInstancePortStates,
		Transform:   transformers.TransformWithStruct(&types.InstancePortState{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "instance_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchLightsailInstancePortStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstancePortStatesInput{InstanceName: r.Name}
	output, err := svc.GetInstancePortStates(ctx, &input)
	if err != nil {
		return err
	}

	res <- output.PortStates
	return nil
}
