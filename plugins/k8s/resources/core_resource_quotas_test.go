package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreResourceQuotas(t *testing.T, ctrl *gomock.Controller) client.Services {
	resourceQuotas := mocks.NewMockResourceQuotasClient(ctrl)
	e := corev1.ResourceQuota{}
	if err := faker.FakeDataSkipFields(&e, []string{"Spec", "Status"}); err != nil {
		t.Fatal(err)
	}
	ss := corev1.ScopeSelector{}
	if err := faker.FakeData(&ss); err != nil {
		t.Fatal(err)
	}
	rqsp := corev1.ResourceQuotaSpec{
		Hard:          *fakeResourceList(t),
		Scopes:        []corev1.ResourceQuotaScope{corev1.ResourceQuotaScopeBestEffort},
		ScopeSelector: &ss,
	}
	rqst := corev1.ResourceQuotaStatus{
		Hard: *fakeResourceList(t),
		Used: *fakeResourceList(t),
	}
	e.Spec = rqsp
	e.Status = rqst
	e.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	resourceQuotas.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.ResourceQuotaList{Items: []corev1.ResourceQuota{e}}, nil,
	)
	return client.Services{
		ResourceQuotas: resourceQuotas,
	}
}

func TestCoreResourceQuotas(t *testing.T) {
	k8sTestHelper(t, CoreResourceQuotas(), createCoreResourceQuotas)
}
