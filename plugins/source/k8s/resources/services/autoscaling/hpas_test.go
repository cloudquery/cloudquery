package autoscaling

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resource "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/autoscaling/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
)

func createHpas(t *testing.T, ctrl *gomock.Controller) client.Services {
	r := resource.HorizontalPodAutoscaler{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockHorizontalPodAutoscalerInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.HorizontalPodAutoscalerList{Items: []resource.HorizontalPodAutoscaler{r}}, nil,
	)

	serviceClient := resourcemock.NewMockAutoscalingV1Interface(ctrl)

	serviceClient.EXPECT().HorizontalPodAutoscalers(metav1.NamespaceAll).Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().AutoscalingV1().Return(serviceClient)

	return client.Services{CoreAPI: cl}
}

func TestHpas(t *testing.T) {
	client.MockTestHelper(t, Hpas(), createHpas)
}
