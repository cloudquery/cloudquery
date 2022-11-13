package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createAppsDaemonSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	daemonSetsClient := client.NewMockInterface(ctrl)
	daemonSetsClient.EXPECT().AppsV1().Return()
	daemonSetsClient.AppsV1().DaemonSets("")
	daemonSetsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.DaemonSetList{Items: []appsv1.DaemonSet{k8sTesting.FakeDaemonSet(t)}}, nil,
	)
	return client.Services{
		DaemonSets: daemonSetsClient,
	}
}

func TestDaemonSets(t *testing.T) {
	client.K8sMockTestHelper(t, DaemonSets(), createAppsDaemonSets, client.TestOptions{})
}
