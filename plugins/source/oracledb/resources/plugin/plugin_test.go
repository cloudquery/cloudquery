package plugin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/oracledb/client"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"github.com/sijms/go-ora/v2/network"
)

func getTestDB(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("oracle", getTestConnectionString())
	if err != nil {
		return nil, err
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return db, err
}

func getTestConnectionString() string {
	testConn := os.Getenv("CQ_SOURCE_ORACLE_DB_TEST_CONNECTION_STRING")
	if testConn == "" {
		return "oracle://admin:xvg0agf3RPE9kyc*vmx@database-1.chdbvwauhrge.us-east-1.rds.amazonaws.com:1521/cloudquery"
	}
	return testConn
}


func isNotExistsError(err error) bool {
	var dbError *network.OracleError
	if errors.As(err, &dbError) {
		return dbError.ErrCode == 942
	}

	return false
}

func TestPlugin(t *testing.T) {
	p := Plugin()
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	p.SetLogger(l)
	spec := specs.Source{
		Name:         "test_oracledb_source",
		Path:         "cloudquery/oracledb",
		Version:      "vDevelopment",
		Destinations: []string{"test"},
		Tables:       []string{"test_oracledb_source"},
		Spec: client.Spec{
			ConnectionString: getTestConnectionString(),
		},
	}
	db, err := getTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, err := db.ExecContext(ctx, "DROP TABLE test_oracledb_source"); err != nil {
		if !isNotExistsError(err) {
			t.Fatal(err)
		}
	}
	if _, err := db.Exec(`CREATE TABLE test_oracledb_source (
		c1 char, 
		c2 nchar, 
		c3 varchar, 
		c4 nvarchar, 
		c5 clob,
		c6 nclob,
		c7 long,
		c8 number,
		c9 date,
		c10 blob
	)
	`); err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec("INSERT INTO test_oracledb_source VALUES ('a', 'b', 'c', 'd', 'e', 'f', 'g', 1, SYSDATE, EMPTY_BLOB())"); err != nil {
		t.Fatal(err)
	}

	if err := p.Init(ctx, spec); err != nil {
		t.Fatal(err)
	}
	res := make(chan *schema.Resource, 1)
	syncTime := time.Now()
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(res)
		return p.Sync(ctx, syncTime, res)
	})
	var resource *schema.Resource
	totalResources := 0
	for r := range res {
		resource = r
		totalResources++
	}
	err = g.Wait()
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if totalResources != 1 {
		t.Fatalf("expected 1 resource, got %d", totalResources)
	}
	fmt.Println(resource)
	// gotData := resource.GetValues()
	// actualStrings := make([]string, len(gotData))
	// for i, v := range gotData {
	// 	actualStrings[i] = v.String()
	// }
	// expectedStrings := make([]string, len(data))
	// for i, v := range data {
	// 	expectedStrings[i] = v.String()
	// }
	// require.Equal(t, expectedStrings, actualStrings)
}
