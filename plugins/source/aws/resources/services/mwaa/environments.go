package mwaa

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Environments() *schema.Table {
	tableName := "aws_mwaa_environments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html`,
		Resolver:            fetchMwaaEnvironments,
		Transform:           transformers.TransformWithStruct(&types.Environment{}),
		PreResourceResolver: getEnvironment,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "airflow"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchMwaaEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := mwaa.ListEnvironmentsInput{}
	c := meta.(*client.Client)
	svc := c.Services().Mwaa
	p := mwaa.NewListEnvironmentsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *mwaa.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Environments
	}
	return nil
}

func getEnvironment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Mwaa
	name := resource.Item.(string)

	output, err := svc.GetEnvironment(ctx, &mwaa.GetEnvironmentInput{Name: &name}, func(options *mwaa.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = output.Environment
	return nil
}
