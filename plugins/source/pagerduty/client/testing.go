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
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	spec := Spec{
		TeamIds:              nil,
		MaxRequestsPerSecond: nil,
		Concurrency:          0,
	}

	pagerdutyClient := pagerduty.NewClient("test_auth_token")
	pagerdutyClient.HTTPClient = buildMockHttpClient()

	schedulerClient, err := New(l, spec, WithClient(pagerdutyClient))
	if err != nil {
		t.Fatal(err)
	}

	tables := schema.Tables{table}
	if err := transformers.TransformTables(tables); err != nil {
		t.Fatal(err)
	}
	sc := scheduler.NewScheduler(scheduler.WithLogger(l))
	messages, err := sc.SyncAll(context.Background(), schedulerClient, tables)
	if err != nil {
		t.Fatal(err)
	}
	inserts := messages.GetInserts()
	records := inserts.GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
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
