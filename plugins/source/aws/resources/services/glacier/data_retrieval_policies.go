package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DataRetrievalPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_glacier_data_retrieval_policies",
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html`,
		Resolver:    fetchGlacierDataRetrievalPolicies,
		Transform:   transformers.TransformWithStruct(&types.DataRetrievalPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glacier"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
