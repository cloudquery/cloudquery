package storage

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/storage/v1"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createCsiStorageCapacities(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.CSIStorageCapacity{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockCSIStorageCapacityInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.CSIStorageCapacityList{Items: []resource.CSIStorageCapacity{r}}, nil,
	)

	serviceClient := resourcemock.NewMockStorageV1Interface(ctrl)

	serviceClient.EXPECT().CSIStorageCapacities("").Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().StorageV1().Return(serviceClient)

	return cl
}

func TestCsiStorageCapacities(t *testing.T) {
	client.K8sMockTestHelper(t, CsiStorageCapacities(), createCsiStorageCapacities)
}
