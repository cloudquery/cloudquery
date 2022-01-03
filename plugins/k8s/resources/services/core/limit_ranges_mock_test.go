//go:build mock
// +build mock

package core

import (
	"testing"

	"github.com/cloudquery/cq-provider-k8s/client"
	"github.com/cloudquery/cq-provider-k8s/client/mocks"
	"github.com/cloudquery/cq-provider-k8s/resources/services/testData"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreLimitRanges(t *testing.T, ctrl *gomock.Controller) client.Services {
	limitRanges := mocks.NewMockLimitRangesClient(ctrl)
	lr := corev1.LimitRange{}
	if err := faker.FakeDataSkipFields(&lr, []string{
		"Spec"}); err != nil {
		t.Fatal(err)
	}
	lr.Spec.Limits = []corev1.LimitRangeItem{
		{
			Type:                 corev1.LimitTypePod,
			Max:                  *testData.FakeResourceList(t),
			Min:                  *testData.FakeResourceList(t),
			Default:              *testData.FakeResourceList(t),
			DefaultRequest:       *testData.FakeResourceList(t),
			MaxLimitRequestRatio: *testData.FakeResourceList(t),
		},
	}
	lr.ManagedFields = []metav1.ManagedFieldsEntry{testData.FakeManagedFields(t)}
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
