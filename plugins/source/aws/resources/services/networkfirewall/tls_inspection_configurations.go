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

func TLSInspectionConfigurations() *schema.Table {
	tableName := "aws_networkfirewall_tls_inspection_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_DescribeTLSInspectionConfiguration.html`,
		Resolver:            fetchTLSInspectionConfigurations,
		PreResourceResolver: getTLSInspectionConfigurations,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.TLSInspectionConfigurationWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("TLSInspectionConfigurationArn"),
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

func fetchTLSInspectionConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListTLSInspectionConfigurationsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	p := networkfirewall.NewListTLSInspectionConfigurationsPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- response.TLSInspectionConfigurations
	}
	return nil
}

func getTLSInspectionConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Networkfirewall
	metadata := resource.Item.(types.TLSInspectionConfigurationMetadata)

	tlsInspectionConfigurationDetails, err := svc.DescribeTLSInspectionConfiguration(ctx, &networkfirewall.DescribeTLSInspectionConfigurationInput{
		TLSInspectionConfigurationArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = cl.Region
	})
	if err != nil && !cl.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.TLSInspectionConfigurationWrapper{
		TLSInspectionConfiguration:         tlsInspectionConfigurationDetails.TLSInspectionConfiguration,
		TLSInspectionConfigurationResponse: tlsInspectionConfigurationDetails.TLSInspectionConfigurationResponse,
	}
	return nil
}
