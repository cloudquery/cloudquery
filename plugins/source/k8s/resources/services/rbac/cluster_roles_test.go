package rbac

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/rbac/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createClusterRoles(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.ClusterRole{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockClusterRoleInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ClusterRoleList{Items: []resource.ClusterRole{r}}, nil,
	)

	serviceClient := resourcemock.NewMockRbacV1Interface(ctrl)

	serviceClient.EXPECT().ClusterRoles().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().RbacV1().Return(serviceClient)

	return cl
}

func TestClusterRoles(t *testing.T) {
	client.K8sMockTestHelper(t, ClusterRoles(), createClusterRoles)
}
