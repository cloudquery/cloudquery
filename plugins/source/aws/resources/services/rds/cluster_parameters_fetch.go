package rds

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRdsClusterParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Rds

	parentEngineVersion := parent.Item.(types.DBEngineVersion)

	if !strings.Contains(aws.ToString(parentEngineVersion.DBParameterGroupFamily), "aurora") {
		return nil
	}

	input := &rds.DescribeEngineDefaultClusterParametersInput{
		DBParameterGroupFamily: parentEngineVersion.DBParameterGroupFamily,
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
