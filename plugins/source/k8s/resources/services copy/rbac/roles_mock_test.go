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
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return &r
}

func TestRbacRoles(t *testing.T) {
	client.K8sMockTestHelper(t, Roles(), createRbacRoles, client.TestOptions{})
}
