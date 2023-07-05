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

func createVolumeAttachments(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.VolumeAttachment{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}

	resourceClient := resourcemock.NewMockVolumeAttachmentInterface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.VolumeAttachmentList{Items: []resource.VolumeAttachment{r}}, nil,
	)

	serviceClient := resourcemock.NewMockStorageV1Interface(ctrl)

	serviceClient.EXPECT().VolumeAttachments().Return(resourceClient)

	cl := mocks.NewMockInterface(ctrl)
	cl.EXPECT().StorageV1().Return(serviceClient)

	return cl
}

func TestVolumeAttachments(t *testing.T) {
	client.K8sMockTestHelper(t, VolumeAttachments(), createVolumeAttachments)
}
