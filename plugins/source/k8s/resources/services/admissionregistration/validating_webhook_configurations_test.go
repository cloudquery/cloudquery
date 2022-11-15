// Code generated by codegen; DO NOT EDIT.

package admissionregistration

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/admissionregistration/v1"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createValidatingWebhookConfigurations(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.ValidatingWebhookConfiguration{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockValidatingWebhookConfigurationInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.ValidatingWebhookConfigurationList{Items: []resource.ValidatingWebhookConfiguration{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAdmissionregistrationV1Interface(ctrl)

	serviceClient.EXPECT().ValidatingWebhookConfigurations().Return(resourceClient)

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().AdmissionregistrationV1().Return(serviceClient)

	return client
}

func TestValidatingWebhookConfigurations(t *testing.T) {
	client.K8sMockTestHelper(t, ValidatingWebhookConfigurations(), createValidatingWebhookConfigurations)
}
