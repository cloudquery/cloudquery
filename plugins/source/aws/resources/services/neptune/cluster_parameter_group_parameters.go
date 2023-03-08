package neptune

import (
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ClusterParameterGroupParameters() *schema.Table {
	tableName := "aws_neptune_cluster_parameter_group_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameterGroups`,
		Resolver:    fetchNeptuneClusterParameterGroupParameters,
		Transform:   transformers.TransformWithStruct(&types.Parameter{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "cluster_parameter_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
