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

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	analyticsdata "google.golang.org/api/analyticsdata/v1beta"
)

type oauthSpec struct {
	AccessToken  string `json:"access_token,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

func (o *oauthSpec) validate() error {
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
func (o *oauthSpec) getTokenSource(ctx context.Context) (oauth2.TokenSource, error) {
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
