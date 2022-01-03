//go:build mock
// +build mock

package apps

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/cloudquery/faker/v3"
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
	if err := faker.FakeDataSkipFields(&rs, []string{"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&rs.Spec, []string{"Template"}); err != nil {
		t.Fatal(err)
	}
	rs.Spec.Template = testData.FakePodTemplateSpec(t)
	rs.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}

	return rs
}

func TestAppsReplicaSets(t *testing.T) {
	client.K8sMockTestHelper(t, ReplicaSets(), createReplicaSets, client.TestOptions{})
}
