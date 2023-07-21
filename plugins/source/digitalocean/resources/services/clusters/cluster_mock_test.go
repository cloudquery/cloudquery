package clusters

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createClusters(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockClustersService(ctrl)

	var data []godo.KubernetesCluster
	if err := faker.FakeObject(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().List(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)

	createNeighbors(t, m)

	return client.Services{
		Clusters: m,
	}
}

func TestClusters(t *testing.T) {
	client.MockTestHelper(t, Clusters(), createClusters, client.TestOptions{})
}
