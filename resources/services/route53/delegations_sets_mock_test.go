//go:build mock
// +build mock

package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRoute53DelegationSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	ds := route53Types.DelegationSet{}
	if err := faker.FakeData(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListReusableDelegationSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListReusableDelegationSetsOutput{
			DelegationSets: []route53Types.DelegationSet{ds},
		}, nil)
	return client.Services{
		Route53: m,
	}
}
func TestRoute53DelegationSets(t *testing.T) {
	client.AwsMockTestHelper(t, Route53ReusableDelegationSets(), buildRoute53DelegationSetsMock, client.TestOptions{})
}
