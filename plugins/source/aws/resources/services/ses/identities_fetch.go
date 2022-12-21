package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	p := sesv2.NewListEmailIdentitiesPaginator(svc, nil)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.EmailIdentities
	}

	return nil
}

func getIdentity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2
	ei := resource.Item.(types.IdentityInfo)

	getOutput, err := svc.GetEmailIdentity(ctx, &sesv2.GetEmailIdentityInput{EmailIdentity: ei.IdentityName})
	if err != nil {
		return err
	}

	resource.SetItem(&models.EmailIdentity{
		IdentityName:           ei.IdentityName,
		SendingEnabled:         ei.SendingEnabled,
		GetEmailIdentityOutput: getOutput,
	})

	return nil
}

func resolveIdentityArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return client.ResolveARN(client.SESService, func(resource *schema.Resource) ([]string, error) {
		return []string{"identity", *resource.Item.(*models.EmailIdentity).IdentityName}, nil
	})(ctx, meta, resource, c)
}
