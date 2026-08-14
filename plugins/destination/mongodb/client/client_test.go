package client

import (
	"context"
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/destination/mongodb/v2/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/auth"
)

func TestOIDCCredential(t *testing.T) {
	cases := []struct {
		name      string
		give      *spec.WorkloadIdentityFederation
		wantProps map[string]string
		wantUser  string
	}{
		{
			name:      "k8s",
			give:      &spec.WorkloadIdentityFederation{Environment: "k8s"},
			wantProps: map[string]string{"ENVIRONMENT": "k8s"},
		},
		{
			name:      "azure",
			give:      &spec.WorkloadIdentityFederation{Environment: "azure", TokenResource: "aud", Username: "client-id"},
			wantProps: map[string]string{"ENVIRONMENT": "azure", "TOKEN_RESOURCE": "aud"},
			wantUser:  "client-id",
		},
		{
			name:      "gcp",
			give:      &spec.WorkloadIdentityFederation{Environment: "gcp", TokenResource: "aud"},
			wantProps: map[string]string{"ENVIRONMENT": "gcp", "TOKEN_RESOURCE": "aud"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cred := oidcCredential(tc.give)
			if cred.AuthMechanism != auth.MongoDBOIDC {
				t.Fatalf("AuthMechanism = %q, want %q", cred.AuthMechanism, auth.MongoDBOIDC)
			}
			if cred.Username != tc.wantUser {
				t.Fatalf("Username = %q, want %q", cred.Username, tc.wantUser)
			}
			if !reflect.DeepEqual(cred.AuthMechanismProperties, tc.wantProps) {
				t.Fatalf("AuthMechanismProperties = %v, want %v", cred.AuthMechanismProperties, tc.wantProps)
			}
		})
	}
}

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_MONGODB_TEST_CONN")
	if testConn == "" {
		return "mongodb://127.0.0.1:27017"
	}
	return testConn
}

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("mongodb", "development", New)
	s := &spec.Spec{
		ConnectionString: getTestConnection(),
		Database:         "destination_mongodb_test",
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Init(ctx, b, plugin.NewClientOptions{}); err != nil {
		t.Fatal(err)
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SkipMigrate:      true,
		},
		plugin.WithTestDataOptions(schema.TestSourceOptions{
			TimePrecision: time.Millisecond,
		}),
	)
}
