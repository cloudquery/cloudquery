package apps

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/apps/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"
)

func createDeployments(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.Deployment{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	r.Spec.Template = corev1.PodTemplateSpec{}
	r.Spec.Strategy = resource.DeploymentStrategy{}

	resourceClient := resourcemock.NewMockDeploymentInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.DeploymentList{Items: []resource.Deployment{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAppsV1Interface(ctrl)

	serviceClient.EXPECT().Deployments("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AppsV1().Return(serviceClient)

	return cl
}

func TestDeployments(t *testing.T) {
	client.K8sMockTestHelper(t, Deployments(), createDeployments)
}
