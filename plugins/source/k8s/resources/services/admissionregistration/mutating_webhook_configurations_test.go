package admissionregistration

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/admissionregistration/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createMutatingWebhookConfigurations(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.MutatingWebhookConfiguration{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockMutatingWebhookConfigurationInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.MutatingWebhookConfigurationList{Items: []resource.MutatingWebhookConfiguration{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAdmissionregistrationV1Interface(ctrl)

	serviceClient.EXPECT().MutatingWebhookConfigurations().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AdmissionregistrationV1().Return(serviceClient)

	return cl
}

func TestMutatingWebhookConfigurations(t *testing.T) {
	client.K8sMockTestHelper(t, MutatingWebhookConfigurations(), createMutatingWebhookConfigurations)
}
