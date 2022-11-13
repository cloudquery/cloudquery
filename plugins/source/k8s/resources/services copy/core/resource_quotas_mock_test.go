package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreResourceQuotas(t *testing.T, ctrl *gomock.Controller) client.Services {
	resourceQuotas := mocks.NewMockResourceQuotasClient(ctrl)
	e := corev1.ResourceQuota{}
	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	ss := corev1.ScopeSelector{}
	if err := faker.FakeObject(&ss); err != nil {
		t.Fatal(err)
	}
	rqsp := corev1.ResourceQuotaSpec{
		Hard:          *k8sTesting.FakeResourceList(t),
		Scopes:        []corev1.ResourceQuotaScope{corev1.ResourceQuotaScopeBestEffort},
		ScopeSelector: &ss,
	}
	rqst := corev1.ResourceQuotaStatus{
		Hard: *k8sTesting.FakeResourceList(t),
		Used: *k8sTesting.FakeResourceList(t),
	}
	e.Spec = rqsp
	e.Status = rqst
	e.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	resourceQuotas.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.ResourceQuotaList{Items: []corev1.ResourceQuota{e}}, nil,
	)
	return client.Services{
		ResourceQuotas: resourceQuotas,
	}
}

func TestCoreResourceQuotas(t *testing.T) {
	client.K8sMockTestHelper(t, ResourceQuotas(), createCoreResourceQuotas, client.TestOptions{})
}
