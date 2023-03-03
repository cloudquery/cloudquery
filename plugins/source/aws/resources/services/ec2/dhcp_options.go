package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DHCPOptions() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_dhcp_options",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DhcpOptions.html`,
		Resolver:    fetchEC2DHCPOptions,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.DhcpOptions{}, transformers.WithPrimaryKeys("DhcpOptionsId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEC2DHCPOptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	pag := ec2.NewDescribeDhcpOptionsPaginator(svc, &ec2.DescribeDhcpOptionsInput{})
	for pag.HasMorePages() {
		page, err := pag.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.DhcpOptions
	}
	return nil
}
