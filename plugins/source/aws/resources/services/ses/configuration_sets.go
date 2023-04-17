package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveConfigurationSetArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			configurationSetEventDestinations(),
		},
	}
}

func fetchSesConfigurationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	p := sesv2.NewListConfigurationSetsPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ConfigurationSets
	}

	return nil
}

func getConfigurationSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2
	csName := resource.Item.(string)

	getOutput, err := svc.GetConfigurationSet(ctx, &sesv2.GetConfigurationSetInput{ConfigurationSetName: &csName})
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
