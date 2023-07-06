package networkfirewall

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/networkfirewall/models"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Firewalls() *schema.Table {
	tableName := "aws_networkfirewall_firewalls"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_DescribeFirewall.html`,
		Resolver:            fetchFirewalls,
		PreResourceResolver: getFirewall,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.FirewallWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FirewallArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchFirewalls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListFirewallsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	p := networkfirewall.NewListFirewallsPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- response.Firewalls
	}
	return nil
}

func getFirewall(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	metadata := resource.Item.(types.FirewallMetadata)

	firewallDetails, err := svc.DescribeFirewall(ctx, &networkfirewall.DescribeFirewallInput{
		FirewallArn: metadata.FirewallArn,
	}, func(options *networkfirewall.Options) {
		options.Region = cl.Region
	})
	if err != nil && !cl.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.FirewallWrapper{
		FirewallStatus: firewallDetails.FirewallStatus,
		Firewall:       firewallDetails.Firewall,
	}
	return nil
}
