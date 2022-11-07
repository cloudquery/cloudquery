package ram

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRamResourceShareAssociatedResourcesMock(t *testing.T, m *mocks.MockRamClient) {
	object := types.ResourceShareAssociation{}
	err := faker.FakeObject(&object)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetResourceShareAssociations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetResourceShareAssociationsOutput{ResourceShareAssociations: []types.ResourceShareAssociation{object}}, nil)
}
