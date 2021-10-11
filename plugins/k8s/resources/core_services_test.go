package resources

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreServices(t *testing.T, ctrl *gomock.Controller) client.Services {
	s := mocks.NewMockServicesClient(ctrl)
	var service corev1.Service
	if err := faker.FakeData(&service); err != nil {
		t.Fatal(err)
	}
	service.Spec.ClusterIP = "192.168.1.1"
	service.Spec.ClusterIPs = []string{"192.168.1.1", "fd00::1"}
	service.Spec.ExternalIPs = []string{"192.168.2.1", "fd00:1::1"}
	s.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.ServiceList{Items: []corev1.Service{service}}, nil,
	)
	return client.Services{
		Services: s,
	}
}

func TestCoreServices(t *testing.T) {
	k8sTestHelper(t, CoreServices(), createCoreServices)
}
