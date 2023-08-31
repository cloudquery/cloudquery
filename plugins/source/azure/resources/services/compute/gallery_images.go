package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/gorilla/mux"
)

func galleryImages() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_gallery_images",
		Resolver:             fetchGalleryImages,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/gallery-images/list-by-gallery?tabs=HTTP#galleryimage",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_gallery_images", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.GalleryImage{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations:            []*schema.Table{galleryImageVersions()},
	}
}

func fetchGalleryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	gallery := parent.Item.(*armcompute.Gallery)
	svc, err := armcompute.NewGalleryImagesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	group, err := client.ParseResourceGroup(*gallery.ID)
	if err != nil {
		return err
	}

	pager := svc.NewListByGalleryPager(group, *gallery.Name, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}

func createMockGalleryImages(router *mux.Router) error {
	var item armcompute.GalleryImageVersionsClientListByGalleryImageResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images", func(w http.ResponseWriter, r *http.Request) {
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
