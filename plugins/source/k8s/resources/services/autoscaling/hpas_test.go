package autoscaling

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/autoscaling/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createHpas(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.HorizontalPodAutoscaler{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockHorizontalPodAutoscalerInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.HorizontalPodAutoscalerList{Items: []resource.HorizontalPodAutoscaler{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAutoscalingV1Interface(ctrl)

	serviceClient.EXPECT().HorizontalPodAutoscalers("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AutoscalingV1().Return(serviceClient)

	return cl
}

func TestHpas(t *testing.T) {
	client.K8sMockTestHelper(t, Hpas(), createHpas)
}
