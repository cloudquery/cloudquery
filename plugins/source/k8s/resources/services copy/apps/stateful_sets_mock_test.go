package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func createStatefulSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	setsClient := mocks.NewMockStatefulSetsClient(ctrl)

	setsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.StatefulSetList{Items: []appsv1.StatefulSet{fakeStatefulSet(t)}}, nil,
	)
	return client.Services{
		StatefulSets: setsClient,
	}
}

func fakeStatefulSet(t *testing.T) appsv1.StatefulSet {
	var rs appsv1.StatefulSet
	if err := faker.FakeObject(&rs); err != nil {
		t.Fatal(err)
	}
	intOrStr := intstr.FromInt(100)
	rs.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable = &intOrStr
	rs.Spec.VolumeClaimTemplates = []corev1.PersistentVolumeClaim{*k8sTesting.FakePersistentVolumeClaim(t)}
	rs.Spec.Selector = k8sTesting.FakeSelector(t)
	rs.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return rs
}

func TestAppsStatefulSets(t *testing.T) {
	client.K8sMockTestHelper(t, StatefulSets(), createStatefulSets, client.TestOptions{})
}
