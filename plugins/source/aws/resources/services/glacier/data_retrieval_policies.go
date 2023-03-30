package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DataRetrievalPolicies() *schema.Table {
	tableName := "aws_glacier_data_retrieval_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html`,
		Resolver:    fetchGlacierDataRetrievalPolicies,
		Transform:   client.TransformWithStruct(&types.DataRetrievalPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
