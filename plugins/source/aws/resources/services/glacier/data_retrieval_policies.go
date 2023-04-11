package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func DataRetrievalPolicies() *schema.Table {
	tableName := "aws_glacier_data_retrieval_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html`,
		Resolver:    fetchGlacierDataRetrievalPolicies,
		Transform:   transformers.TransformWithStruct(&types.DataRetrievalPolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchGlacierDataRetrievalPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier

	response, err := svc.GetDataRetrievalPolicy(ctx, &glacier.GetDataRetrievalPolicyInput{})
	if err != nil {
		return err
	}
	res <- response.Policy
	return nil
}
