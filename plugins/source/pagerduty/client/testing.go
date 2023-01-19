package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"
)

type MockHttpClient struct {
	// A map from request path to response object()
	// e.g. "/users" -> []User
	mockResponses map[string]any
}

func (mockHttpClient *MockHttpClient) AddMockResponse(url string, object any) {
	if mockHttpClient.mockResponses == nil {
		mockHttpClient.mockResponses = make(map[string]any)
	}

	mockHttpClient.mockResponses[url] = object
}

func PagerdutyMockTestHelper(t *testing.T, table *schema.Table, buildMockHttpClient func() *MockHttpClient) {
	t.Helper()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	table.IgnoreInTests = false

	version := "vDev"

	configureTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		pagerdutyClient := pagerduty.NewClient("test_auth_token")
		pagerdutyClient.HTTPClient = buildMockHttpClient()

		cqClient := Client{
			PagerdutyClient: pagerdutyClient,
			logger:          logger,
		}

		return &cqClient, nil
	}

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{table},
		configureTestExecutionClient,
	)
	p.SetLogger(logger)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}

func (mockHttpClient *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	mockResponseObject := mockHttpClient.mockResponses[strings.TrimRight(req.URL.Path, "/")]
	marshaledMockResponse, err := json.Marshal(mockResponseObject)

	if err != nil {
		panic(err)
	}

	httpResponse := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Body:       io.NopCloser(bytes.NewReader(marshaledMockResponse)),
	}

	return &httpResponse, nil
}

// Timestamp fields such as `created_at` arrive as `string`s from the API.
// this function pusts valid `RFC3339` timestamps into fields like `DeletedAt`, `CreatedAt`...
// Mostly copy-paste from the `plugin-sdk` faker.
// Receives an interface that is a pointer to a struct, and only looks at fields one level deep.
// Pionter-to-pointer structs are supported.
func FakeStringTimestamps(ptrObj any) error {
	timestampFieldNames := []string{
		"CreateAt", "CreatedAt", "DeletedAt", "LastStatusChangeAt", "StartTime", "EndTime", "LastIncidentTimestamp",
	}

	ptrType := reflect.TypeOf(ptrObj) // reflection-type is a pointer
	ptrKind := ptrType.Kind()
	ptrValue := reflect.ValueOf(ptrObj) // a reflection-value that is a pointer

	if ptrKind != reflect.Ptr {
		return fmt.Errorf("object must be a pointer")
	}

	if ptrValue.IsNil() {
		return fmt.Errorf("object must not be nil")
	}

	if ptrValue.Elem().Kind() == reflect.Ptr {
		return FakeStringTimestamps(ptrValue.Elem().Interface())
	}

	if ptrValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("object must be a pointer to a struct")
	}

	structValue := ptrValue.Elem()
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldType := field.Type
		fieldKind := fieldType.Kind()
		fieldValue := structValue.Field(i)

		if fieldKind == reflect.String {
			if !fieldValue.CanSet() { // don't panic on unexported fields
				continue
			}

			if slices.Contains(timestampFieldNames, field.Name) {
				fieldValue.Set(reflect.ValueOf(time.Now().Format(time.RFC3339)))
			}
		}
	}

	return nil
}
