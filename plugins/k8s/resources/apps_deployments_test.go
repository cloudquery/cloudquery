package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createAppsDeployments(t *testing.T, ctrl *gomock.Controller) client.Services {
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
	if err := faker.FakeDataSkipFields(&deployment, []string{"Spec"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(&deployment.Spec, []string{"Template"}); err != nil {
		t.Fatal(err)
	}
	deployment.Spec.Template = fakePodTemplateSpec(t)
	deployment.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	return deployment
}

func TestAppsDeployments(t *testing.T) {
	k8sTestHelper(t, AppsDeployments(), createAppsDeployments)
}
