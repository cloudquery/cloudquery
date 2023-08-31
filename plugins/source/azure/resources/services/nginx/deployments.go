package nginx

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Deployments() *schema.Table {
	return &schema.Table{
		Name:                 "azure_nginx_deployments",
		Resolver:             fetchDeployments,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v2@v2.0.0#Deployment",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_nginx_deployments", client.Namespacenginx_nginxplus),
		Transform:            transformers.TransformWithStruct(&armnginx.Deployment{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
