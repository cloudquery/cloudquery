// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/{{.Service}}"
	"github.com/aws/aws-sdk-go-v2/service/{{.Service}}/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func build{{.Service | ToCamel}}{{.SubService | ToCamel}}Mock(t *testing.T, ctrl *gomock.Controller) client.Services {
    m := mocks.NewMock{{.CloudqueryServiceName}}Client(ctrl)
    object := types.{{.StructName}}{}
    err := faker.FakeObject(&object)
    if err != nil {
		t.Fatal(err)
    }

    m.EXPECT().Get{{.StructName}}s(gomock.Any(), gomock.Any(), gomock.Any()).Return(
        &{{.Service}}.Get{{.StructName}}sOutput{ {{.StructName}}s: []types.{{.StructName}}{object}}, nil)

    return client.Services{ {{.CloudqueryServiceName}}: m}
}

func Test{{.Service | ToCamel}}{{.SubService | ToCamel}}(t *testing.T) {
    client.AwsMockTestHelper(t, {{.SubService | ToCamel}}(), build{{.Service | ToCamel}}{{.SubService | ToCamel}}Mock, client.TestOptions{})
}