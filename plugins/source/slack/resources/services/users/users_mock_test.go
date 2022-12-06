package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/slack-go/slack"
)

func createUsers() (*slack.Client, error) {
	items := make([]slack.User, 1)
	if err := faker.FakeObject(&items); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.POST("/users.list", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(items)
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
	opts := []slack.Option{
		slack.OptionAPIURL(ts.URL + "/"),
	}
	c := slack.New("test", opts...)
	return c, nil
}

func TestUsers(t *testing.T) {
	client.MockTestHelper(t, Users(), createUsers, client.TestOptions{})
}
