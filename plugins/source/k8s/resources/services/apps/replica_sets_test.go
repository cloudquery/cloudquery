package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apps/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createReplicaSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.ReplicaSet{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.Template = corev1.PodTemplateSpec{}

	resourceClient := resourcemock.NewMockReplicaSetInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ReplicaSetList{Items: []resource.ReplicaSet{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAppsV1Interface(ctrl)

	serviceClient.EXPECT().ReplicaSets(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AppsV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestReplicaSets(t *testing.T) {
	client.MockTestHelper(t, ReplicaSets(), createReplicaSets)
}
