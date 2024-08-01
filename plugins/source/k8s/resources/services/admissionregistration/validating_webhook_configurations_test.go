package admissionregistration

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/admissionregistration/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createValidatingWebhookConfigurations(t *testing.T, ctrl *gomock.Controller) client.Services {
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

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AdmissionregistrationV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestValidatingWebhookConfigurations(t *testing.T) {
	client.MockTestHelper(t, ValidatingWebhookConfigurations(), createValidatingWebhookConfigurations)
}
