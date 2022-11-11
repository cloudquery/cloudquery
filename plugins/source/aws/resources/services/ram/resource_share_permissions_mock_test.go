package ram

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRamResourceSharePermissionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRamClient(ctrl)
	summary := types.ResourceSharePermissionSummary{}
	err := faker.FakeObject(&summary)
	if err != nil {
		t.Fatal(err)
	}

	var version int32
	err = faker.FakeObject(&version)
	if err != nil {
		t.Fatal(err)
	}
	verStr := fmt.Sprint(version)
	summary.Version = &verStr

	m.EXPECT().ListResourceSharePermissions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.ListResourceSharePermissionsOutput{Permissions: []types.ResourceSharePermissionSummary{summary}}, nil)

	detail := ""
	err = faker.FakeObject(&detail)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetPermission(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ram.GetPermissionOutput{Permission: &types.ResourceSharePermissionDetail{Permission: &detail}}, nil)

	return client.Services{Ram: m}
}

func TestRamResourceSharePermissions(t *testing.T) {
	client.AwsMockTestHelper(t, ResourceSharePermissions(), buildRamResourceSharePermissionsMock, client.TestOptions{})
}
