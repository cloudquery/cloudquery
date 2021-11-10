package resources

import (
	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
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
	if err := faker.FakeDataSkipFields(&rs, []string{"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&rs.Spec, []string{"Template"}); err != nil {
		t.Fatal(err)
	}
	rs.Spec.Template = fakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}

	return rs
}

func TestAppsReplicaSets(t *testing.T) {
	k8sTestHelper(t, AppsReplicaSets(), createReplicaSets)
}
