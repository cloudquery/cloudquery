package annotations

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createAnnotations(router *mux.Router) error {
	var o mixpanel.Annotation
	if err := faker.FakeObject(&o); err != nil {
		return err
	}

	router.HandleFunc("/api/app/projects/"+strconv.FormatInt(client.TestProjectID, 10)+"/annotations", func(w http.ResponseWriter, r *http.Request) {
		list := mixpanel.AnnotationList{
			CommonResponse: mixpanel.CommonResponse{
				Status: mixpanel.StatusOK,
			},
			Results: []mixpanel.Annotation{o},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestAnnotations(t *testing.T) {
	client.MockTestHelper(t, Annotations(), createAnnotations, client.TestOptions{})
}
