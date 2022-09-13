package {{.Service}}

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

{{$func := .SubService | ToCamel}}
{{$func =  printf "%v%v" "List" $func}}
{{if ne .FunctionName ""}}
{{$func = .FunctionName}}
{{end}}


func create{{.SubServiceName | ToCamel}}(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMock{{.Service | ToCamel}}Service(ctrl)
{{if .MockWrapper}}
    data := &godo.{{.MockStructName}}{}
	if err := faker.FakeData(data); err != nil {
		t.Fatal(err)
	}
	data.Links = nil
{{else}}
    var data  []{{if .IsStructPointer}}*{{end}}godo.{{.MockStructName}}
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}
{{end}}

	m.EXPECT().{{$func}}(gomock.Any(),gomock.Any(){{if ne .ParentStructName ""}},gomock.Any(){{end}}).Return(data, &godo.Response{}, nil)

    {{range $val := .Relations -}}
    create{{$val}}(t, m)
    {{end}}

	return client.Services{
		{{.Service | ToCamel}}: m,
	}
}

func Test{{.SubServiceName | ToCamel}}(t *testing.T) {
	client.MockTestHelper(t, {{.SubServiceName | ToCamel}}(), create{{.SubServiceName | ToCamel}}, client.TestOptions{})
}