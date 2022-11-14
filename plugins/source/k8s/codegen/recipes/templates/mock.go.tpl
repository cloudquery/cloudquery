// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"testing"

	client "github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	mocks "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks"

	resourcemock "github.com/cloudquery/cloudquery/plugins/source/k8s/mocks/{{.ImportPath}}"
	// k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/golang/mock/gomock"
	resource "k8s.io/api/{{.ImportPath}}"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"github.com/cloudquery/plugin-sdk/faker"
	{{range .MockImports}}
	{{.}}
	{{end}}
)

func create{{.SubService | ToCamel}}(t *testing.T, ctrl *gomock.Controller) kubernetes.Interface {
	r := resource.{{.StructName}}{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	{{.FakerOverride}}

	resourceClient := resourcemock.NewMock{{.StructName}}Interface(ctrl)
	resourceClient.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&resource.{{.StructName}}List{Items: []resource.{{.StructName}}{r}}, nil,
	)

	serviceClient := resourcemock.NewMock{{.ServiceFuncName}}Interface(ctrl)
	{{if .GlobalResource}}
		serviceClient.EXPECT().{{.ResourceFuncName}}().Return(resourceClient)
	{{else}}
		serviceClient.EXPECT().{{.ResourceFuncName}}("").Return(resourceClient)
	{{end}}
	

	client := mocks.NewMockInterface(ctrl)
	client.EXPECT().{{.ServiceFuncName}}().Return(serviceClient)

	return client
}

func Test{{.SubService | ToCamel}}(t *testing.T) {
	client.K8sMockTestHelper(t, {{.SubService | ToCamel}}(), create{{.SubService | ToCamel}})
}
