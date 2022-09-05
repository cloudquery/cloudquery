// Code generated by codegen using template resource_list_describe.go.tpl; DO NOT EDIT.

package athena

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/athena/types"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

func AthenaWorkGroupPreparedStatements() *schema.Table {
	return &schema.Table{
		Name:      "aws_athena_work_group_prepared_statements",
		Resolver:  fetchAthenaWorkGroupPreparedStatements,
		Multiplex: client.ServiceAccountRegionMultiplexer("athena"),
		Columns: []schema.Column{
			{
				Name:     "workgroup_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIdResolver,
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "query_statement",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("QueryStatement"),
			},
			{
				Name:     "statement_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StatementName"),
			},
			{
				Name:     "work_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WorkGroupName"),
			},
		},
	}
}

func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena

	r1 := parent.Item.(types.WorkGroup)

	input := athena.ListPreparedStatementsInput{
		WorkGroup: r1.Name,
	}
	paginator := athena.NewListPreparedStatementsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {

			return diag.WrapError(err)
		}

		for _, item := range output.PreparedStatements {

			do, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
				WorkGroup: r1.Name,

				StatementName: item.StatementName,
			})
			if err != nil {

				if cl.IsNotFoundError(err) {
					continue
				}
				return diag.WrapError(err)
			}
			res <- do.PreparedStatement
		}
	}
	return nil
}
