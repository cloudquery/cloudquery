package applications

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func handleApplications(router *mux.Router) error {
	var a okta.AutoLoginApplication
	if err := faker.FakeObject(&a); err != nil {
		return err
	}
	a.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	a.Profile = map[string]any{"top-key": "value"}

	as := okta.ApplicationSignOnMode("AUTO_LOGIN")
	a.SignOnMode = &as

	a.Credentials.Password.Hash.Algorithm = &okta.AllowedPasswordCredentialHashAlgorithmEnumValues[0]
	a.Credentials.Scheme = &okta.AllowedApplicationCredentialsSchemeEnumValues[0]
	a.Credentials.Signing.Use = &okta.AllowedApplicationCredentialsSigningUseEnumValues[0]
	a.Credentials.Password.Hash.DigestAlgorithm = &okta.AllowedDigestAlgorithmEnumValues[0]
	lcs := okta.APPLICATIONLIFECYCLESTATUS_ACTIVE
	a.Status = &lcs
	a.Links = &okta.ApplicationLinks{
		Self: &okta.HrefObjectSelfLink{Href: "#"},
	}

	router.HandleFunc("/api/v1/apps", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal([]okta.AutoLoginApplication{a})
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
	ag.Profile = map[string]map[string]any{"top-key": {"key": "value"}}
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
	au.Profile = map[string]map[string]any{"top-key": {"key": "value"}}
	au.AdditionalProperties = map[string]any{"key": "value"}
	au.Status = okta.APPUSERSTATUS_ACTIVE
	au.SyncState = &okta.AllowedAppUserSyncStateEnumValues[0]
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
