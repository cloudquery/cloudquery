//go:build mock
// +build mock

package rbac

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
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
	if err := faker.FakeData(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{testing.FakeManagedFields(t)}
	return &r
}

func TestRbacRoleBindings(t *testing.T) {
	client.K8sMockTestHelper(t, RoleBindings(), createRbacRoleBindings, client.TestOptions{})

}
