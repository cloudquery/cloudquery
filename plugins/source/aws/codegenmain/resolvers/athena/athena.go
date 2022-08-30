package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CreateDataCatalogArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "datacatalog", catalogName)
}

func ResolveDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.DataCatalog)
	return diag.WrapError(resource.Set(c.Name, CreateDataCatalogArn(cl, *dc.Name)))
}

func CreateWorkGroupArn(cl *client.Client, catalogName string) string {
	return cl.ARN(client.Athena, "workgroup", catalogName)
}

func ResolveWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(types.WorkGroup)
	return diag.WrapError(resource.Set(c.Name, CreateWorkGroupArn(cl, *dc.Name)))
}
