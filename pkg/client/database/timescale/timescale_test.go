//go:build history
// +build history

package timescale

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client/history"
	pgsdk "github.com/cloudquery/cq-provider-sdk/database/postgres"
	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

type mockConn struct {
	row pgx.Row
}

func (m mockConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return m.row
}

type mockScanner struct {
	t   *testing.T
	val string
	err error
}

func (m mockScanner) Scan(dst ...interface{}) error {
	if len(dst) != 1 {
		m.t.Fatalf("called with %d args, want exactly one", len(dst))
	}
	ptr, ok := dst[0].(*string)
	if !ok {
		m.t.Fatalf("received %T, expected *string", dst[0])
	}
	*ptr = m.val
	return m.err
}

var testTable = &schema.Table{
	Name: "test_table",
	Columns: []schema.Column{
		{
			Name: "id",
			Type: schema.TypeString,
		},
	},
	Relations: []*schema.Table{
		{
			Name: "test_rel_table",
			Columns: []schema.Column{
				{
					Name:     "parent_cq_id",
					Type:     schema.TypeUUID,
					Resolver: schema.ParentIdResolver,
				},
				{
					Name: "test",
					Type: schema.TypeString,
				},
			},
		},
	},
	Options: schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
}

func getDSN() string {
	dbDSN := os.Getenv("CQ_TIMESCALE_TEST_DSN")
	if dbDSN == "" {
		dbDSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable" // timescale
	}
	return dbDSN
}

func TestSetupHistory(t *testing.T) {
	ctx := context.TODO()
	ts, err := New(hclog.L(), getDSN(), &history.Config{
		Retention:      1,
		TimeInterval:   1,
		TimeTruncation: 24,
	})
	assert.NoError(t, err)

	ok, err := ts.Validate(ctx)
	assert.NoError(t, err)
	assert.True(t, ok)

	migrationDSN, err := ts.Setup(ctx)
	assert.NoError(t, err)

	err = ts.Prepare(ctx)
	assert.NoError(t, err)

	{
		pool, err := pgsdk.Connect(ctx, migrationDSN)
		assert.NoError(t, err)
		defer pool.Close()

		conn, err := pool.Acquire(ctx)
		assert.NoError(t, err)
		defer conn.Release()

		tc := migration.NewTableCreator(hclog.L(), schema.TSDBDialect{})
		ups, downs, err := tc.CreateTableDefinitions(ctx, testTable, nil)
		assert.NoError(t, err)

		newDowns := make([]string, len(downs))
		for i, sql := range downs {
			if strings.HasPrefix(sql, "DROP TABLE ") {
				sql = strings.TrimSuffix(sql, ";") + " CASCADE"
			}
			newDowns[i] = sql
		}
		defer func() {
			for _, sql := range newDowns {
				_, err = conn.Exec(ctx, sql)
				assert.NoError(t, err)
			}
		}()

		for _, sql := range append(newDowns, ups...) { // DROP old tables first, if they exist
			_, err = conn.Exec(ctx, sql)
			assert.NoError(t, err)
		}
	}

	err = ts.Finalize(ctx, err)
	assert.NoError(t, err)

	t.Run("FinalizeSecondTime", func(t *testing.T) {
		// Finalize() again (after setup, to connect) shouldn't create any errors
		_, err := ts.Setup(ctx)
		assert.NoError(t, err)

		err = ts.Finalize(ctx, nil)
		assert.NoError(t, err)
	})

	pool, err := pgsdk.Connect(ctx, getDSN())
	assert.NoError(t, err)
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	assert.NoError(t, err)
	defer conn.Release()

	t.Run("QueryView", func(t *testing.T) {
		_, err = conn.Exec(ctx, "select cq_fetch_date from test_table")
		assert.Nil(t, err)
	})

	t.Run("QueryHistoryTable", func(t *testing.T) {
		_, err = conn.Exec(ctx, "select cq_fetch_date from history.test_table")
		assert.Nil(t, err)
	})

	partitionDate := time.Now().Format("2006/01/02")

	t.Run("Insert", func(t *testing.T) {
		const (
			sqlInsertMainTable = `INSERT INTO public.test_table(cq_id, cq_meta, cq_fetch_date, id)
	VALUES ('0d0bf7c6-c87d-4b3c-a270-60246dcb6ab1', NULL, TO_DATE('%s', 'YYYY/MM/DD'), 'test_id')`
			sqlInsertRelTable = `INSERT INTO public.test_rel_table(cq_id, cq_meta, cq_fetch_date, parent_cq_id, test)
	VALUES (gen_random_uuid(), null, TO_DATE('%s', 'YYYY/MM/DD'), '0d0bf7c6-c87d-4b3c-a270-60246dcb6ab1', 'test2')`
		)

		_, err = conn.Exec(ctx, fmt.Sprintf(sqlInsertMainTable, partitionDate))
		assert.NoError(t, err)
		_, err = conn.Exec(ctx, fmt.Sprintf(sqlInsertRelTable, partitionDate))
		assert.NoError(t, err)
	})

	t.Run("Select", func(t *testing.T) {
		res, err := conn.Exec(ctx, "select * from test_rel_table")
		assert.NoError(t, err)
		assert.Equal(t, res.RowsAffected(), int64(1))
	})

	t.Run("DeleteCascadeTrigger", func(t *testing.T) {
		res, err := conn.Exec(ctx, fmt.Sprintf(`DELETE FROM test_table WHERE cq_fetch_date = TO_DATE('%s', 'YYYY/MM/DD')`, partitionDate))
		assert.NoError(t, err)
		assert.Equal(t, res.RowsAffected(), int64(1))
		res, err = conn.Exec(ctx, "select * from test_rel_table")
		assert.NoError(t, err)
		assert.Equal(t, res.RowsAffected(), int64(0))
	})
}

func Test_doValidateTimescaleVersion(t *testing.T) {
	tests := []struct {
		name       string
		q          mockConn
		minVersion string
		wantErr    error
	}{
		{
			"lower than needed",
			mockConn{row: mockScanner{t, "1.7.5", nil}},
			"2.0",
			errors.New("unsupported Timescale version: 1.7.5. (should be >= 2.0.0)"),
		},
		{
			"equal",
			mockConn{row: mockScanner{t, "2.0.0", nil}},
			"2.0",
			nil,
		},
		{
			"greater than needed",
			mockConn{row: mockScanner{t, "2.6.0", nil}},
			"2.0",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := version.Must(version.NewVersion(tt.minVersion))
			err := doValidateTimescaleVersion(context.Background(), tt.q, want)
			if (tt.wantErr == nil) != (err == nil) {
				t.Errorf("wantErr is %v, returned error is %v", tt.wantErr, err)
			}
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
