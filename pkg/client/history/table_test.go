//go:build history
// +build history

package history_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/client/history"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

const (
	testDBConnection   = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	sqlInsertMainTable = `INSERT INTO public.test_table(
	cq_id, meta, cq_fetch_date, test)
	VALUES ('0d0bf7c6-c87d-4b3c-a270-60246dcb6ab1', NULL, TO_DATE('%s', 'YYYY/MM/DD'), 'test');
	`
	sqlInsertRelTable = `INSERT INTO public.test_rel_table(
	cq_id, meta, cq_fetch_date, parent_cq_id, test)
	VALUES (gen_random_uuid(), null, TO_DATE('%s', 'YYYY/MM/DD'), '0d0bf7c6-c87d-4b3c-a270-60246dcb6ab1', 'test2');
	`
)

var testTable = &schema.Table{
	Name: "test_table",
	Columns: []schema.Column{
		{
			Name: "test",
			Type: schema.TypeString,
		},
	},
	Relations: []*schema.Table{
		{
			Name: "test_rel_table",
			Columns: []schema.Column{
				{
					Name: "parent_cq_id",
					Type: schema.TypeUUID,
				},
				{
					Name: "test",
					Type: schema.TypeString,
				},
			},
		},
	},
	Options: schema.TableCreationOptions{PrimaryKeys: []string{"test"}},
}

func TestHistory_SetupHistory(t *testing.T) {
	pool, err := client.CreateDatabase(context.Background(), testDBConnection)
	assert.NoError(t, err)
	defer pool.Close()
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
	defer conn.Release()
	assert.NoError(t, history.SetupHistory(context.Background(), conn))
}

func TestHistoryTableCreator_CreateTables(t *testing.T) {
	m, err := history.NewHistoryTableCreator(&history.Config{Retention: 1,
		TimeInterval:   1,
		TimeTruncation: 24,
	}, hclog.L())
	assert.NoError(t, err)
	assert.NotNil(t, m)

	pool, err := client.CreateDatabase(context.Background(), testDBConnection)
	assert.NoError(t, err)
	defer pool.Close()
	conn, err := pool.Acquire(context.Background())
	assert.NoError(t, err)
	defer conn.Release()
	// Call setup history as previous test can execute before
	assert.NoError(t, history.SetupHistory(context.Background(), conn))

	assert.NoError(t, m.CreateTable(context.Background(), conn, testTable, nil))
	// creating tables again shouldn't create any errors
	assert.NoError(t, m.CreateTable(context.Background(), conn, testTable, nil))
	// query the view
	_, err = conn.Exec(context.Background(), "select cq_fetch_date from test_table")
	assert.Nil(t, err)
	// query the history table itself
	_, err = conn.Exec(context.Background(), "select cq_fetch_date from history.test_table")
	assert.Nil(t, err)
	partitionDate := time.Now().Format("2006/01/02")
	_, err = conn.Exec(context.Background(), fmt.Sprintf(sqlInsertMainTable, partitionDate))
	assert.Nil(t, err)
	_, err = conn.Exec(context.Background(), fmt.Sprintf(sqlInsertRelTable, partitionDate))
	// Check data was inserted
	res, err := conn.Exec(context.Background(), "select * from test_rel_table")
	assert.Nil(t, err)
	assert.Equal(t, res.RowsAffected(), int64(1))
	// Test that delete cascade trigger works
	res, err = conn.Exec(context.Background(), fmt.Sprintf(`DELETE FROM test_table WHERE cq_fetch_date = TO_DATE('%s', 'YYYY/MM/DD')`, partitionDate))
	assert.Nil(t, err)
	assert.Equal(t, res.RowsAffected(), int64(1))
	res, err = conn.Exec(context.Background(), "select * from test_rel_table")
	assert.Nil(t, err)
	assert.Equal(t, res.RowsAffected(), int64(0))
}
