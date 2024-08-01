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

func createClusterRoleBindings(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.ClusterRoleBinding{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockClusterRoleBindingInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ClusterRoleBindingList{Items: []resource.ClusterRoleBinding{r}}, nil,
	)

	serviceClient := resourcemock.NewMockRbacV1Interface(ctrl)

	serviceClient.EXPECT().ClusterRoleBindings().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().RbacV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestClusterRoleBindings(t *testing.T) {
	client.MockTestHelper(t, ClusterRoleBindings(), createClusterRoleBindings)
}
