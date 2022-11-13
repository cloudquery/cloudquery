package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createReplicaSets(t *testing.T, ctrl *gomock.Controller) client.Services {
	setsClient := mocks.NewMockReplicaSetsClient(ctrl)
	setsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.ReplicaSetList{Items: []appsv1.ReplicaSet{fakeReplicaSet(t)}}, nil,
	)
	return client.Services{
		ReplicaSets: setsClient,
	}
}

func fakeReplicaSet(t *testing.T) appsv1.ReplicaSet {
	var rs appsv1.ReplicaSet
	if err := faker.FakeObject(&rs); err != nil {
		t.Fatal(err)
	}
	rs.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}

	return rs
}

func TestAppsReplicaSets(t *testing.T) {
	client.K8sMockTestHelper(t, ReplicaSets(), createReplicaSets, client.TestOptions{})
}
