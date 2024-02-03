package servicefabric

import (
	"context"
	"net/http"

	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Clusters() *schema.Table {
	return &schema.Table{
		Name:                 "azure_servicefabric_clusters",
		Resolver:             fetchServicefabricClusters,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/servicefabric/clusters/list?view=rest-servicefabric-2019-11-01-preview&tabs=HTTP",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_servicefabric_clusters", client.Namespacemicrosoft_servicefabric),
		Transform:            transformers.TransformWithStruct(&armservicefabric.Cluster{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchServicefabricClusters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armservicefabric.NewClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	// first request
	clusters, err := svc.List(ctx, nil)
	if err != nil {
		return err
	}
	for _, cluster := range clusters.Value {
		res <- cluster
	}

	// following next links
	nextLink := clusters.NextLink
	pl, err := armruntime.NewPipeline("", "", cl.Creds, runtime.PipelineOptions{}, cl.Options)
	if err != nil {
		return err
	}
	for nextLink != nil && *nextLink != "" {
		req, err := runtime.NewRequest(ctx, http.MethodGet, *nextLink)
		if err != nil {
			return err
		}
		resp, err := pl.Do(req)
		if err != nil {
			return err
		}
		if !runtime.HasStatusCode(resp, http.StatusOK) {
			return runtime.NewResponseError(resp)
		}
		result := &armservicefabric.ClustersClientListResponse{}
		if err := runtime.UnmarshalAsJSON(resp, result); err != nil {
			return err
		}
		for _, cluster := range result.Value {
			res <- cluster
		}
		nextLink = result.NextLink
	}

	return nil
}
