package settings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/xanzy/go-gitlab"
)

func buildSettings(mux *httprouter.Router) error {
	var settings gitlab.Settings
	if err := faker.FakeObject(&settings); err != nil {
		return err
	}
	resp, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	mux.GET("/api/v4/application/settings", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, string(resp))
	})
	return nil
}

func TestSettings(t *testing.T) {
	client.GitlabMockTestHelper(t, Settings(), buildSettings, client.TestOptions{})
}
