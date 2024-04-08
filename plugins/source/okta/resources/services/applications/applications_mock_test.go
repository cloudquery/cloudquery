package applications

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v4/okta"
)

func handleApplications(router *mux.Router) error {
	var a okta.AutoLoginApplication
	if err := faker.FakeObject(&a); err != nil {
		return err
	}
	al := "AUTO_LOGIN"
	a.SignOnMode = &al
	a.Profile = map[string]any{"top-key": "value"}
	a.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	a.Links = &okta.ApplicationLinks{
		Self: &okta.HrefObjectSelfLink{Href: "#"},
	}

	var o okta.OpenIdConnectApplication
	if err := faker.FakeObject(&o); err != nil {
		return err
	}
	oi := "OPENID_CONNECT"
	o.SignOnMode = &oi
	o.Profile = map[string]any{"top-key": "value"}
	o.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}

	router.HandleFunc("/api/v1/apps", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal([]okta.ListApplications200ResponseInner{
			okta.AutoLoginApplicationAsListApplications200ResponseInner(&a),
			okta.OpenIdConnectApplicationAsListApplications200ResponseInner(&o),
		})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var ag okta.ApplicationGroupAssignment
	if err := faker.FakeObject(&ag); err != nil {
		return err
	}
	ag.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	ag.Profile = map[string]any{"top-key": "value"}
	ag.AdditionalProperties = map[string]any{"key": "value"}
	ag.Links = &okta.LinksSelf{
		Self: &okta.HrefObjectSelfLink{Href: "#"},
	}

	router.HandleFunc("/api/v1/apps/"+*a.Id+"/groups", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal([]okta.ApplicationGroupAssignment{ag})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var au okta.AppUser
	if err := faker.FakeObject(&au); err != nil {
		return err
	}
	au.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	au.Profile = map[string]any{"top-key": "value"}
	au.AdditionalProperties = map[string]any{"key": "value"}
	t := time.Now()
	au.PasswordChanged = *okta.NewNullableTime(&t)
	au.Links = okta.LinksAppAndUser{
		App: &okta.HrefObjectAppLink{Href: "#"},
	}

	router.HandleFunc("/api/v1/apps/"+*a.Id+"/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal([]okta.AppUser{au})
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

func TestApplications(t *testing.T) {
	client.MockTestHelper(t, Applications(), handleApplications)
}
