package client

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"strings"

	"github.com/invopop/jsonschema"
	orderedmap "github.com/wk8/go-ordered-map/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type OAuthSpec struct {
	AccessToken  string `json:"access_token,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func (OAuthSpec) JSONSchemaExtend(sc *jsonschema.Schema) {
	sc.If = &jsonschema.Schema{
		Not: &jsonschema.Schema{
			// access_token is present & not empty
			Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
				properties := jsonschema.NewProperties()
				// if access_token is empty
				accessToken := *sc.Properties.Value("access_token")
				one := uint64(1)
				accessToken.MinLength = &one
				properties.Set("access_token", &accessToken)
				return properties
			}(),
			Required: []string{"access_token"},
		},
	}
	sc.Then = &jsonschema.Schema{
		Properties: func() *orderedmap.OrderedMap[string, *jsonschema.Schema] {
			properties := jsonschema.NewProperties()
			// then client_id & client_secret are required & mustn't be empty
			clientID := *sc.Properties.Value("client_id")
			clientSecret := *sc.Properties.Value("client_secret")
			one := uint64(1)
			clientID.MinLength = &one
			clientSecret.MinLength = &one
			properties.Set("client_id", &clientID)
			properties.Set("client_secret", &clientSecret)
			return properties
		}(),
		Required: []string{"client_id", "client_secret"},
	}
}

func (o *OAuthSpec) validate() error {
	if o == nil {
		// OAuth is optional
		return nil
	}

	switch {
	case len(o.AccessToken) > 0:
		return nil
	case len(o.ClientID) == 0:
		return fmt.Errorf("empty client_id in oauth spec")
	case len(o.ClientSecret) == 0:
		return fmt.Errorf("empty client_secret in oauth spec")
	default:
		return nil
	}
}
func (o *OAuthSpec) getTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	if len(o.AccessToken) > 0 {
		return oauth2.StaticTokenSource(&oauth2.Token{AccessToken: o.AccessToken}), nil
	}

	lst, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	config := &oauth2.Config{
		ClientID:     o.ClientID,
		ClientSecret: o.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://" + lst.Addr().String(),
		Scopes:       []string{analyticsdata.AnalyticsReadonlyScope},
	}

	b := make([]byte, 16)
	rand.Read(b)
	state := strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")

	handler := &oauthHandler{
		state: state,
		err:   make(chan error),
	}

	srv := http.Server{Handler: handler}

	go func() {
		defer srv.Close()
		_ = exec.CommandContext(ctx, "open", config.AuthCodeURL(state, oauth2.AccessTypeOffline)).Run()
		err = <-handler.err
	}()

	if serveErr := srv.Serve(lst); serveErr != http.ErrServerClosed {
		return nil, serveErr
	}

	if err != nil {
		return nil, err
	}

	// we have exchange token now
	token, err := config.Exchange(ctx, handler.code, oauth2.AccessTypeOffline)
	if err != nil {
		return nil, err
	}

	return config.TokenSource(context.Background(), token), nil
}

type oauthHandler struct {
	state string
	code  string
	err   chan error
}

func (o *oauthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer close(o.err)

	if state := r.FormValue("state"); state != o.state {
		w.WriteHeader(http.StatusBadRequest)
		err := fmt.Errorf("incorrect \"state\" value: expected %q, got %q", o.state, state)
		fmt.Fprint(w, err.Error())
		o.err <- err
		return
	}

	o.code = r.FormValue("code")
	o.err <- nil

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Authorization successful. You may close the window.")
}
