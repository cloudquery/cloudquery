// Code generated by codegen; DO NOT EDIT.

package {{.Service}}

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	{{range .MockImports}}
  "{{.}}"
  {{end}}
	"google.golang.org/api/option"
)

type Mock{{.SubService | ToCamel}}Result struct {
  {{.OutputField}} []*{{.Service}}.{{.StructName}} `json:"{{.OutputField | ToLower}},omitempty"`
}

func create{{.SubService | ToCamel}}() (*client.Services, error) {
	var item {{.Service}}.{{.StructName}}
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
  {{.MockPostFaker}}
	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &Mock{{.SubService | ToCamel}}Result{
			{{.OutputField}}: []*{{.Service}}.{{.StructName}}{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := {{.Service}}.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		{{.Service|ToCamel}}: svc,
	}, nil
}

func Test{{.SubService | ToCamel}}(t *testing.T) {
	client.MockTestHelper(t, {{.SubService | ToCamel}}(), create{{.SubService | ToCamel}}, client.TestOptions{})
}