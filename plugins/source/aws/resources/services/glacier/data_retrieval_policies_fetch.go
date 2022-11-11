package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlacierDataRetrievalPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier

	response, err := svc.GetDataRetrievalPolicy(ctx, &glacier.GetDataRetrievalPolicyInput{})
	if err != nil {
		return err
	}
	res <- response.Policy
	return nil
}
