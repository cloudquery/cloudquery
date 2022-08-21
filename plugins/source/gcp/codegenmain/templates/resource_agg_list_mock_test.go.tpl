// Code generated by codegen; DO NOT EDIT.

package {{.GCPService}}

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	{{range .MockImports}}
  "{{.}}"
  {{end}}
	"google.golang.org/api/option"
)

func create{{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}() (*client.Services, error) {
	var item {{.GCPService}}.{{.GCPStructName}}
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/{{.GCPSubService}}", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &{{.GCPService}}.{{.MockListStruct}}AggregatedList{
			Items: map[string]{{.GCPService}}.{{.GCPSubService | ToCamel}}ScopedList{
				"": {
					{{.GCPSubService | ToCamel}}: []*{{.GCPService}}.{{.GCPStructName}}{&item},
				},
			},
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
	svc, err := {{.GCPService}}.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		{{.GCPService|ToCamel}}: svc,
	}, nil
}

func Test{{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}(t *testing.T) {
	client.GcpMockTestHelper(t,  {{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}(), create{{.GCPService | ToCamel}}{{.GCPSubService | ToCamel}}, client.TestOptions{})
}
