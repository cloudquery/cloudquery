package rbac

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createRbacRoleBindings(t *testing.T, ctrl *gomock.Controller) client.Services {
	roles := mocks.NewMockRoleBindingsClient(ctrl)
	roles.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&v1.RoleBindingList{Items: []v1.RoleBinding{*fakeRoleBinding(t)}}, nil,
	)
	return client.Services{
		RoleBindings: roles,
	}
}

func fakeRoleBinding(t *testing.T) *v1.RoleBinding {
	r := v1.RoleBinding{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return &r
}

func TestRbacRoleBindings(t *testing.T) {
	client.K8sMockTestHelper(t, RoleBindings(), createRbacRoleBindings, client.TestOptions{})
}
