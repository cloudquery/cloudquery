package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesConfigurationSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	config := sesv2.ListConfigurationSetsInput{}
	p := sesv2.NewListConfigurationSetsPaginator(svc, &config)
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
	resource.Item = getOutput
	return nil
}
