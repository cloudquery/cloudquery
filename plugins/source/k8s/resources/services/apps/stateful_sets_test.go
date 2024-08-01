package apps

import (
	"testing"

	resource "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"
	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apps/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createStatefulSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.StatefulSet{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.Template = corev1.PodTemplateSpec{}
	i := intstr.FromInt(5000)
	r.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable = &i
	r.Spec.VolumeClaimTemplates[0].ObjectMeta.ManagedFields = nil

	resourceClient := resourcemock.NewMockStatefulSetInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.StatefulSetList{Items: []resource.StatefulSet{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAppsV1Interface(ctrl)

	serviceClient.EXPECT().StatefulSets(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AppsV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestStatefulSets(t *testing.T) {
	client.MockTestHelper(t, StatefulSets(), createStatefulSets)
}
