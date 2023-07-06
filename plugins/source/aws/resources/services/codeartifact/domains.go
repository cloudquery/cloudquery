package codeartifact

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
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
		Transform:           transformers.TransformWithStruct(&types.DomainDescription{}, transformers.WithPrimaryKeys("Arn")),
		Columns: []schema.Column{
			{
				Name:       "request_account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "request_region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
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
	svc := cl.Services().Codeartifact
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
	svc := cl.Services().Codeartifact
	domainOutput, err := svc.DescribeDomain(ctx, &codeartifact.DescribeDomainInput{Domain: domain.Name}, func(options *codeartifact.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = domainOutput.Domain
	return nil
}
