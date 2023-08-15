package ses

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ConfigurationSets() *schema.Table {
	tableName := "aws_ses_configuration_sets"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/ses/latest/APIReference-V2/API_GetConfigurationSet.html`,
		Resolver:            fetchSesConfigurationSets,
		PreResourceResolver: getConfigurationSet,
		Transform: transformers.TransformWithStruct(
			&sesv2.GetConfigurationSetOutput{},
			transformers.WithSkipFields("ResultMetadata"),
			transformers.WithNameTransformer(client.CreateTrimPrefixTransformer("configuration_set_")),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "email"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveConfigurationSetArn,
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			configurationSetEventDestinations(),
		},
	}
}

func fetchSesConfigurationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2

	p := sesv2.NewListConfigurationSetsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *sesv2.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ConfigurationSets
	}

	return nil
}

func getConfigurationSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sesv2
	csName := resource.Item.(string)

	getOutput, err := svc.GetConfigurationSet(ctx, &sesv2.GetConfigurationSetInput{ConfigurationSetName: &csName}, func(o *sesv2.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.SetItem(getOutput)

	return nil
}

func resolveConfigurationSetArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"configuration-set", *resource.Item.(*sesv2.GetConfigurationSetOutput).ConfigurationSetName}, nil
	})(ctx, meta, resource, c)
}
