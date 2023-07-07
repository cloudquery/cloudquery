package lightsail

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func databaseParameters() *schema.Table {
	tableName := "aws_lightsail_database_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_RelationalDatabaseParameter.html`,
		Resolver:    fetchLightsailDatabaseParameters,
		Transform:   transformers.TransformWithStruct(&types.RelationalDatabaseParameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "database_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchLightsailDatabaseParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.RelationalDatabase)
	input := lightsail.GetRelationalDatabaseParametersInput{
		RelationalDatabaseName: r.Name,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lightsail
	// No paginator available
	for {
		response, err := svc.GetRelationalDatabaseParameters(ctx, &input, func(options *lightsail.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Parameters
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}
