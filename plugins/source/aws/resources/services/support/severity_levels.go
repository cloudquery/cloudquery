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

func SeverityLevels() *schema.Table {
	return &schema.Table{
		Name:        "aws_support_severity_levels",
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeSeverityLevels.html`,
		Resolver:    fetchSeverityLevels,
		Transform:   transformers.TransformWithStruct(&types.SeverityLevel{}, transformers.WithPrimaryKeys("Code")),
		Multiplex:   client.ServiceAccountRegionsLanguageCodeMultiplex("support", []string{"en", "ja"}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
	}
}

func fetchSeverityLevels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	client := meta.(*client.Client)
	svc := client.Services().Support
	input := support.DescribeSeverityLevelsInput{Language: aws.String(client.LanguageCode)}

	response, err := svc.DescribeSeverityLevels(ctx, &input)
	if err != nil {
		return err
	}

	res <- response.SeverityLevels

	return nil
}
