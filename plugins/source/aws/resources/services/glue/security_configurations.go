package glue

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SecurityConfigurations() *schema.Table {
	tableName := "aws_glue_security_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_SecurityConfiguration.html`,
		Resolver:    fetchGlueSecurityConfigurations,
		Transform:   transformers.TransformWithStruct(&types.SecurityConfiguration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				PrimaryKey: true,
			},
		},
	}
}

func fetchGlueSecurityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	paginator := glue.NewGetSecurityConfigurationsPaginator(svc, &glue.GetSecurityConfigurationsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SecurityConfigurations
	}
	return nil
}
