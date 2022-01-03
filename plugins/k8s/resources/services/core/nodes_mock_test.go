//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreNodes(t *testing.T, ctrl *gomock.Controller) client.Services {
	nodes := mocks.NewMockNodesClient(ctrl)
	nodes.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.NodeList{Items: []corev1.Node{testData.FakeNode(t)}}, nil,
	)
	return client.Services{
		Nodes: nodes,
	}
}

func TestCoreNodes(t *testing.T) {
	client.K8sMockTestHelper(t, Nodes(), createCoreNodes, client.TestOptions{})
}
