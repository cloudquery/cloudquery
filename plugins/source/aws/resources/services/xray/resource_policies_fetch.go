package xray

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchXrayResourcePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	paginator := xray.NewListResourcePoliciesPaginator(meta.(*client.Client).Services().Xray, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.ResourcePolicies
	}
	return nil
}
func createXrayResourcePolicyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rp := resource.Item.(types.ResourcePolicy)
	cl := meta.(*client.Client)

	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.XRayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/resource-policy/%s", aws.ToString(rp.PolicyName)),
	}.String())

}
