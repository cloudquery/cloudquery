package apps

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func createDeployments(t *testing.T, ctrl *gomock.Controller) client.Services {
	deploymentsClient := mocks.NewMockDeploymentsClient(ctrl)
	deploymentsClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&appsv1.DeploymentList{Items: []appsv1.Deployment{fakeAppsDeployment(t)}}, nil,
	)
	return client.Services{
		Deployments: deploymentsClient,
	}
}

func fakeAppsDeployment(t *testing.T) appsv1.Deployment {
	var deployment appsv1.Deployment
	if err := faker.FakeObject(&deployment); err != nil {
		t.Fatal(err)
	}

	intOrStr1 := intstr.FromInt(100)
	intOrStr2 := intstr.FromInt(100)
	deployment.Spec.Strategy.RollingUpdate.MaxSurge = &intOrStr1
	deployment.Spec.Strategy.RollingUpdate.MaxUnavailable = &intOrStr2

	deployment.Spec.Template = k8sTesting.FakePodTemplateSpec(t)
	deployment.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	return deployment
}

func TestDeployments(t *testing.T) {
	client.K8sMockTestHelper(t, Deployments(), createDeployments, client.TestOptions{})
}
