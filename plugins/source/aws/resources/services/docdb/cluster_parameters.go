package docdb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterParameters() *schema.Table {
	tableName := "aws_docdb_cluster_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html`,
		Resolver:    fetchDocdbClusterParameters,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "docdb"),
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchDocdbClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Docdb
	switch item := parent.Item.(type) {
	case types.DBClusterParameterGroup:
		return fetchParameterGroupParameters(ctx, meta, svc, item, res)
	case types.DBEngineVersion:
		return fetchEngineVersionParameters(ctx, meta, svc, item, res)
	}
	return fmt.Errorf("wrong parrent type to fetch cluster parameters")
}

func fetchParameterGroupParameters(ctx context.Context, meta schema.ClientMeta, svc services.DocdbClient, item types.DBClusterParameterGroup, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}
	p := docdb.NewDescribeDBClusterParametersPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *docdb.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Parameters
	}
	return nil
}

func fetchEngineVersionParameters(ctx context.Context, meta schema.ClientMeta, svc services.DocdbClient, item types.DBEngineVersion, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &docdb.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: item.DBParameterGroupFamily,
	}
	output, err := svc.DescribeEngineDefaultClusterParameters(ctx, input, func(options *docdb.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if output.EngineDefaults == nil || len(output.EngineDefaults.Parameters) == 0 {
		return nil
	}
	res <- output.EngineDefaults.Parameters
	return nil
}
