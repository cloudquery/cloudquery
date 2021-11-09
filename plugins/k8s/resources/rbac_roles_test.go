package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createRbacRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	roles := mocks.NewMockRolesClient(ctrl)
	roles.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&v1.RoleList{Items: []v1.Role{*fakeRole(t)}}, nil,
	)
	return client.Services{
		Roles: roles,
	}
}

func fakeRole(t *testing.T) *v1.Role {
	r := v1.Role{}
	if err := faker.FakeData(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{*fakeManagedFields(t)}
	return &r
}

func TestRbacRoles(t *testing.T) {
	k8sTestHelper(t, RbacRoles(), createRbacRoles)
}
