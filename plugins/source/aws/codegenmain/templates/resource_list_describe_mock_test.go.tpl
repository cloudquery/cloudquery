// Code generated by codegen; DO NOT EDIT.

package {{.AWSService | ToLower}}

import (
	"testing"

{{if .HasTags}}
  "github.com/aws/aws-sdk-go-v2/aws"
{{end}}
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"{{.TypesImport}}"
{{range .MockImports}}	{{.}}
{{end}}
)

func {{.MockFuncName}}(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMock{{.AWSService}}Client(ctrl)

	var item {{.PaginatorListType}}
	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().{{.ListMethod}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.ListMethod}}Input{},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.ListMethod}}Output{
		  {{.PaginatorListName}}: []{{.PaginatorListType}}{item},
    },
		nil,
	)

  var detail types.{{.AWSStructName}}
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}
{{range $v := .GetAndListOrder}}
	detail.{{$v}} = {{index $.MatchedGetAndListFields $v}}
{{end}}

	mock.EXPECT().{{.GetMethod}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.GetMethod}}Input{
{{range $v := .GetAndListOrder}}
	{{$v}}: {{index $.MatchedGetAndListFields $v}},
{{end}}
		},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.GetMethod}}Output{
		  {{.ItemName}}: &detail,
    },
		nil,
	)

{{if .HasTags}}
	mock.EXPECT().ListTagsFor{{.ItemName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
{{range $v := .GetAndListOrder}}
		{{$v}}: {{index $.MatchedGetAndListFields $v}},
{{end}}
    },
	).Return(
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Output{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
{{end}}
	return client.Services{
	  {{.AWSService}}: mock,
  }
}

func {{.TestFuncName}}(t *testing.T) {
	client.AwsMockTestHelper(t, {{.TableFuncName}}(), {{.MockFuncName}}, client.TestOptions{})
}
