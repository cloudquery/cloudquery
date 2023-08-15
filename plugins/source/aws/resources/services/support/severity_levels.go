package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

var severitySupportedLanguageCodes = []string{"en", "ja"}

func SeverityLevels() *schema.Table {
	tableName := "aws_support_severity_levels"

	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeSeverityLevels.html`,
		Resolver:    fetchSeverityLevels,
		Transform:   transformers.TransformWithStruct(&types.SeverityLevel{}, transformers.WithPrimaryKeys("Code")),
		Multiplex:   client.ServiceAccountRegionsLanguageCodeMultiplex(tableName, "support", severitySupportedLanguageCodes),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			client.LanguageCodeColumn(true),
		},
	}
}

func fetchSeverityLevels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Support
	input := support.DescribeSeverityLevelsInput{Language: aws.String(cl.LanguageCode)}

	response, err := svc.DescribeSeverityLevels(ctx, &input, func(o *support.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.SeverityLevels

	return nil
}
