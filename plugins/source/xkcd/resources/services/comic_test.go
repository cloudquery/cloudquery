package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/hermanschaaf/cq-source-xkcd/client"
	"github.com/hermanschaaf/cq-source-xkcd/internal/xkcd"
)

func TestComicsTable(t *testing.T) {
	var comic xkcd.Comic
	if err := faker.FakeObject(&comic); err != nil {
		t.Fatal(err)
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		d, _ := json.Marshal(comic)
		_, _ = w.Write(d)
	}))
	defer ts.Close()

	client.TestHelper(t, ComicsTable(), ts)
}
