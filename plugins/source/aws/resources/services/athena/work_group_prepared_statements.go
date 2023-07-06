package athena

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func workGroupPreparedStatements() *schema.Table {
	tableName := "aws_athena_work_group_prepared_statements"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/athena/latest/APIReference/API_PreparedStatement.html`,
		Resolver:            fetchAthenaWorkGroupPreparedStatements,
		PreResourceResolver: getWorkGroupPreparedStatement,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "athena"),
		Transform:           transformers.TransformWithStruct(&types.PreparedStatement{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "work_group_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	wg := parent.Item.(types.WorkGroup)
	input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
	paginator := athena.NewListPreparedStatementsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *athena.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.PreparedStatements
	}
	return nil
}

func getWorkGroupPreparedStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	wg := resource.Parent.Item.(types.WorkGroup)

	d := resource.Item.(types.PreparedStatementSummary)
	dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
		WorkGroup:     wg.Name,
		StatementName: d.StatementName,
	}, func(options *athena.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.PreparedStatement
	return nil
}
