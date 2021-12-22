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
			Max:                  *fakeResourceList(t),
			Min:                  *fakeResourceList(t),
			Default:              *fakeResourceList(t),
			DefaultRequest:       *fakeResourceList(t),
			MaxLimitRequestRatio: *fakeResourceList(t),
		},
	}
	lr.ManagedFields = []metav1.ManagedFieldsEntry{fakeManagedFields(t)}
	limitRanges.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.LimitRangeList{Items: []corev1.LimitRange{lr}}, nil,
	)
	return client.Services{
		LimitRanges: limitRanges,
	}
}

func TestCoreLimitRanges(t *testing.T) {
	k8sTestHelper(t, CoreLimitRanges(), createCoreLimitRanges)
}
