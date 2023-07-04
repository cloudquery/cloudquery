package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	compute "cloud.google.com/go/compute/apiv1"
)

func machineTypes() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_machine_types",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/machineTypes/list#response-body`,
		Resolver:    fetchMachineTypes,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.MachineType{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchMachineTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	zone := parent.Item.(*pb.Zone)
	var maxResults = uint32(500)
	req := &pb.ListMachineTypesRequest{
		Project:    c.ProjectId,
		Zone:       *zone.Name,
		MaxResults: &maxResults,
	}
	gcpClient, err := compute.NewMachineTypesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.List(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp
	}
	return nil
}

func createMachineTypes(mux *httprouter.Router, zones *pb.ZoneList) error {
	var machineTypes pb.MachineTypeList
	if err := faker.FakeObject(&machineTypes); err != nil {
		return err
	}
	emptyStr := ""
	machineTypes.NextPageToken = &emptyStr
	mux.GET("/compute/v1/projects/testProject/zones/"+*zones.Items[0].Name+"/machineTypes", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&machineTypes)
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
