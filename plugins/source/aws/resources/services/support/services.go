package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var servicesSupportedLanguageCodes = []string{"en", "ja"}

func Services() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_services",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeServices.html`,
		Resolver:    fetchServices,
		Transform:   transformers.TransformWithStruct(&types.Service{}, transformers.WithPrimaryKeys("Code")),
		Multiplex:   client.ServiceAccountRegionsLanguageCodeMultiplex("support", servicesSupportedLanguageCodes),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Support
	input := support.DescribeServicesInput{Language: aws.String(c.LanguageCode)}

	response, err := svc.DescribeServices(ctx, &input)
	if err != nil {
		return err
	}

	res <- response.Services

	return nil
}
