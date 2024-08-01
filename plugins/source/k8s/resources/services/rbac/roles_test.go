package rbac

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/rbac/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.Role{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockRoleInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.RoleList{Items: []resource.Role{r}}, nil,
	)

	serviceClient := resourcemock.NewMockRbacV1Interface(ctrl)

	serviceClient.EXPECT().Roles(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().RbacV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestRoles(t *testing.T) {
	client.MockTestHelper(t, Roles(), createRoles)
}
