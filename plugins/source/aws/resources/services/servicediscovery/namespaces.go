package servicediscovery

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Namespaces() *schema.Table {
	tableName := "aws_servicediscovery_namespaces"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloud-map/latest/api/API_Namespace.html`,
		Resolver:            fetchNamespaces,
		PreResourceResolver: getNamespace,
		Transform:           transformers.TransformWithStruct(&types.Namespace{}, transformers.WithPrimaryKeys("Arn")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "servicediscovery"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveServicediscoveryTags("Arn"),
			},
		},
	}
}
func fetchNamespaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicediscovery
	input := servicediscovery.ListNamespacesInput{MaxResults: aws.Int32(100)}
	paginator := servicediscovery.NewListNamespacesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *servicediscovery.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Namespaces
	}
	return nil
}

func getNamespace(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Servicediscovery
	namespace := resource.Item.(types.NamespaceSummary)

	desc, err := svc.GetNamespace(ctx, &servicediscovery.GetNamespaceInput{Id: namespace.Id}, func(o *servicediscovery.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = desc.Namespace
	return nil
}
