package {{.Service}}

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func create{{.Service | ToCamel}}(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMock{{.Service | ToCamel}}Service(ctrl)

	var data godo.{{.Service | ToCamel}}
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().Get(gomock.Any()).Return(&data, nil, nil)

    //add children mocks
    {{range $val := .Relations}}create{{$val}}(t, m){{end}}

	return client.Services{
		{{.Service | ToCamel}}: m,
	}
}

func Test{{.Service | ToCamel}}(t *testing.T) {
	client.MockTestHelper(t, {{.Service | ToCamel}}(), create{{.Service | ToCamel}}, client.TestOptions{})
}