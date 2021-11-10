package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createAppsDaemonSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	daemonSetsClient := mocks.NewMockDaemonSetsClient(ctrl)
	daemonSetsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.DaemonSetList{Items: []appsv1.DaemonSet{fakeDaemonSet(t)}}, nil,
	)
	return client.Services{
		DaemonSets: daemonSetsClient,
	}
}

func TestAppsDaemonSets(t *testing.T) {
	k8sTestHelper(t, AppsDaemonSets(), createAppsDaemonSets)
}
