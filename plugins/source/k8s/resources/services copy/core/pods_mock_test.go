package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCorePods(t *testing.T, ctrl *gomock.Controller) client.Services {
	pods := mocks.NewMockPodsClient(ctrl)
	pods.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.PodList{Items: []corev1.Pod{k8sTesting.FakePod(t)}}, nil,
	)
	return client.Services{
		Pods: pods,
	}
}

func TestCorePods(t *testing.T) {
	client.K8sMockTestHelper(t, Pods(), createCorePods, client.TestOptions{})
}
