package pendo

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestHttpClient_prepareUrls(t *testing.T) {
	c := &HttpClient{}
	err := c.prepareUrls()
	if err != nil {
		t.Errorf("prepareUrls() error = %v", err)
	}
}

func TestHttpClient_validate(t *testing.T) {
	type fields struct {
		http   Doer
		logger *zerolog.Logger
		apiKey string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty api key",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "empty doer client",
			fields: fields{
				apiKey: "test",
			},
			wantErr: true,
		},
		{
			name: "empty logger",
			fields: fields{
				apiKey: "test",
				http:   &http.Client{},
			},
			wantErr: true,
		},
		{
			name: "valid doer client",
			fields: fields{
				apiKey: "test",
				http:   &http.Client{},
				logger: &zerolog.Logger{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HttpClient{
				doer:   tt.fields.http,
				logger: tt.fields.logger,
				apiKey: tt.fields.apiKey,
			}
			if err := c.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getter_features_list(t *testing.T) {
	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()
	got, err := getter[Feature]{}.list(context.Background(), &logger, "", u, mockDoer(t, "features.json"))
	if err != nil {
		t.Errorf("list() error = %v", err)
		return
	}

	if len(got) == 0 {
		t.Errorf("got unexpected empty features list ")
		return
	}
}

func Test_getter_guides_list(t *testing.T) {
	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()
	got, err := getter[Guide]{}.list(context.Background(), &logger, "", u, mockDoer(t, "guides.json"))
	if err != nil {
		t.Errorf("list() error = %v", err)
		return
	}

	if len(got) == 0 {
		t.Errorf("got unexpected empty guides list ")
		return
	}
}

func Test_getter_track_types_list(t *testing.T) {
	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()
	got, err := getter[TrackType]{}.list(context.Background(), &logger, "", u, mockDoer(t, "tracktypes.json"))
	if err != nil {
		t.Errorf("list() error = %v", err)
		return
	}

	if len(got) == 0 {
		t.Errorf("got unexpected empty track types list ")
		return
	}
}

func Test_getter_track_reports_list(t *testing.T) {
	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()
	got, err := getter[Report]{}.list(context.Background(), &logger, "", u, mockDoer(t, "reports.json"))
	if err != nil {
		t.Errorf("list() error = %v", err)
		return
	}

	if len(got) == 0 {
		t.Errorf("got unexpected empty reports list ")
		return
	}
}

func Test_getter_set_api_key(t *testing.T) {
	doer := doerFn(func(r *http.Request) (*http.Response, error) {
		if apiKey := r.Header.Get("X-Pendo-Integration-Key"); apiKey != "test" {
			t.Errorf("got unexpected api key %s", apiKey)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       http.NoBody,
		}, nil
	})

	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()
	_, _ = getter[Report]{}.list(context.Background(), &logger, "test", u, doer)
}

// tests for getter list function
func Test_getter_pages_list(t *testing.T) {
	type testCase[T any] struct {
		name    string
		doer    Doer
		ge      getter[T]
		wantErr bool
	}
	tests := []testCase[Page]{
		{
			name:    "empty list of pages",
			doer:    mockDoer(t, "empty_list.json"),
			ge:      getter[Page]{},
			wantErr: false,
		},
		{
			name:    "list of pages",
			doer:    mockDoer(t, "pages.json"),
			ge:      getter[Page]{},
			wantErr: false,
		},
		{
			name:    "list of pages with error",
			ge:      getter[Page]{},
			doer:    doerFn(func(r *http.Request) (*http.Response, error) { return nil, errors.New("test error") }),
			wantErr: true,
		},
		{
			name:    "list of pages with bad json",
			doer:    mockDoer(t, "bad_json.json"),
			ge:      getter[Page]{},
			wantErr: true,
		},
		{
			name: "list of pages with bad status code",
			doer: doerFn(func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       http.NoBody,
				}, nil
			}),
			ge:      getter[Page]{},
			wantErr: true,
		},
	}

	u, _ := url.Parse("doer://example.com")
	logger := zerolog.Nop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.ge.list(context.Background(), &logger, "", u, tt.doer)
			if (err != nil) != tt.wantErr {
				t.Errorf("list() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

type doerFn func(r *http.Request) (*http.Response, error)

func (d doerFn) Do(r *http.Request) (*http.Response, error) {
	return d(r)
}

func mockDoer(t *testing.T, fileName string) Doer {
	t.Helper()

	file, err := os.Open("testdata/" + fileName)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	return doerFn(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       file,
		}, nil
	})
}
