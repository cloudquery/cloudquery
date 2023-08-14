package containerservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/gorilla/mux"
)

func clusterUpgradeProfiles() *schema.Table {
	return &schema.Table{
		Name:                 "azure_containerservice_managed_cluster_upgrade_profiles",
		Resolver:             fetchManagedClusterUpgradeProfile,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/aks/managed-clusters/get-upgrade-profile?tabs=HTTP#managedclusterupgradeprofile",
		Transform:            transformers.TransformWithStruct(&armcontainerservice.ManagedClusterUpgradeProfile{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchManagedClusterUpgradeProfile(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	cluster := parent.Item.(*armcontainerservice.ManagedCluster)
	svc, err := armcontainerservice.NewManagedClustersClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}

	group, err := client.ParseResourceGroup(*cluster.ID)
	if err != nil {
		return err
	}
	resp, err := svc.GetUpgradeProfile(ctx, group, *cluster.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.ManagedClusterUpgradeProfile
	return nil
}

func createManagedClusterUpgradeProfile(router *mux.Router) error {
	var item armcontainerservice.ManagedClustersClientGetUpgradeProfileResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{group}/providers/Microsoft.ContainerService/managedClusters/{cluster}/upgradeProfiles/default", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
