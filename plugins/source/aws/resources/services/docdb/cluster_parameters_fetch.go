package docdb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDocdbClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Docdb
	switch item := parent.Item.(type) {
	case types.DBClusterParameterGroup:
		return fetchParameterGroupParameters(ctx, svc, item, res)
	case types.DBEngineVersion:
		return fetchEngineVersionParameters(ctx, svc, item, res)
	}
	return fmt.Errorf("wrong parrent type to fetch cluster parameters")
}

func fetchParameterGroupParameters(ctx context.Context, svc services.DocdbClient, item types.DBClusterParameterGroup, res chan<- interface{}) error {
	input := &docdb.DescribeDBClusterParametersInput{
		DBClusterParameterGroupName: item.DBClusterParameterGroupName,
	}
	p := docdb.NewDescribeDBClusterParametersPaginator(svc, input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Parameters
	}
	return nil
}

func fetchEngineVersionParameters(ctx context.Context, svc services.DocdbClient, item types.DBEngineVersion, res chan<- interface{}) error {
	input := &docdb.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: item.DBParameterGroupFamily,
	}
	output, err := svc.DescribeEngineDefaultClusterParameters(ctx, input)
	if err != nil {
		return err
	}
	if output.EngineDefaults == nil || len(output.EngineDefaults.Parameters) == 0 {
		return nil
	}
	res <- output.EngineDefaults.Parameters
	return nil
}
