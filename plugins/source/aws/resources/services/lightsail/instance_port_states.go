package lightsail

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func instancePortStates() *schema.Table {
	tableName := "aws_lightsail_instance_port_states"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_InstancePortState.html`,
		Resolver:    fetchLightsailInstancePortStates,
		Transform: transformers.TransformWithStruct(&types.InstancePortState{},
			transformers.WithPrimaryKeyComponents("FromPort", "ToPort", "Protocol"),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "allow_list",
				Description:         "This column contains a concatenated list of all allowed addresses",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveInstancePortAllowList,
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLightsailInstancePortStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services(client.AWSServiceLightsail).Lightsail
	input := lightsail.GetInstancePortStatesInput{InstanceName: r.Name}
	output, err := svc.GetInstancePortStates(ctx, &input, func(options *lightsail.Options) {
		options.Region = cli.Region
	})
	if err != nil {
		return err
	}

	res <- output.PortStates
	return nil
}

func resolveInstancePortAllowList(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	// Inspired by https://docs.aws.amazon.com/cli/latest/reference/lightsail/put-instance-public-ports.html
	state := resource.Item.(types.InstancePortState)
	cidrs := "cidrs=" + strings.Join(state.Cidrs, ",")
	ipv6Cidrs := "ipv6Cidrs=" + strings.Join(state.Ipv6Cidrs, ",")
	cidrListAliases := "cidrListAliases=" + strings.Join(state.CidrListAliases, ",")
	return resource.Set(c.Name, cidrs+","+ipv6Cidrs+","+cidrListAliases)
}
