package docdb

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func clusterParameters() *schema.Table {
	tableName := "aws_docdb_cluster_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Parameter.html`,
		Resolver:    fetchDocdbClusterParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}, transformers.WithPrimaryKeyComponents("ParameterName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "engine",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("engine"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "engine_version",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("engine_version"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchDocdbClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDocdb).Docdb
	input := &docdb.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: parent.Item.(types.DBEngineVersion).DBParameterGroupFamily,
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
