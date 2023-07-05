package certificates

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/certificates/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createSigningRequests(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.CertificateSigningRequest{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockCertificateSigningRequestInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.CertificateSigningRequestList{Items: []resource.CertificateSigningRequest{r}}, nil,
	)

	serviceClient := resourcemock.NewMockCertificatesV1Interface(ctrl)

	serviceClient.EXPECT().CertificateSigningRequests().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().CertificatesV1().Return(serviceClient)

	return cl
}

func TestSigningRequests(t *testing.T) {
	client.K8sMockTestHelper(t, SigningRequests(), createSigningRequests)
}
