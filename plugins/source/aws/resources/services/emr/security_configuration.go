package emr

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SecurityConfigurations() *schema.Table {
	tableName := "aws_emr_security_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_DescribeSecurityConfiguration.html`,
		Resolver:    fetchSecurityConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:   transformers.TransformWithStruct(&types.SecurityConfigurationSummary{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "account_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSAccount,
				PrimaryKey: true,
			},
			{
				Name:       "region",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveAWSRegion,
				PrimaryKey: true,
			},
			{
				Name:     "security_configuration",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveConfiguration,
			},
		},
	}
}

func fetchSecurityConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Emr
	paginator := emr.NewListSecurityConfigurationsPaginator(svc, &emr.ListSecurityConfigurationsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.SecurityConfigurations
	}
	return nil
}

func resolveConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.SecurityConfigurationSummary)
	svc := cl.Services().Emr
	response, err := svc.DescribeSecurityConfiguration(ctx, &emr.DescribeSecurityConfigurationInput{Name: item.Name}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(col.Name, response.SecurityConfiguration)
}
