package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Domains() *schema.Table {
	tableName := "aws_codeartifact_domains"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:            fetchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codeartifact"),
		Transform:           transformers.TransformWithStruct(&types.DomainDescription{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCodeartifactTags("Arn"),
			},
		},
		Relations: []*schema.Table{},
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	paginator := codeartifact.NewListDomainsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codeartifact.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Domains
	}
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	domain := resource.Item.(types.DomainSummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	domainOutput, err := svc.DescribeDomain(ctx, &codeartifact.DescribeDomainInput{Domain: domain.Name}, func(options *codeartifact.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = domainOutput.Domain
	return nil
}
