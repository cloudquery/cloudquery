package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreLimitRanges(t *testing.T, ctrl *gomock.Controller) client.Services {
	limitRanges := mocks.NewMockLimitRangesClient(ctrl)

	lr := corev1.LimitRange{}
	if err := faker.FakeObject(&lr); err != nil {
		t.Fatal(err)
	}
	lr.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	limitRanges.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.LimitRangeList{Items: []corev1.LimitRange{lr}}, nil,
	)
	return client.Services{
		LimitRanges: limitRanges,
	}
}

func TestCoreLimitRanges(t *testing.T) {
	client.K8sMockTestHelper(t, LimitRanges(), createCoreLimitRanges, client.TestOptions{})
}
