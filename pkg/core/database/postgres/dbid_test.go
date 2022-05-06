package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-sdk/database"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseId(t *testing.T) {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}

	pool, err := database.New(context.Background(), hclog.NewNullLogger(), dbUrl)
	assert.NoError(t, err)

	dbId1, err := GetDatabaseId(context.Background(), pool)
	assert.NoError(t, err)
	assert.NotEmpty(t, dbId1)

	dbId2, err := GetDatabaseId(context.Background(), pool)
	assert.NoError(t, err)
	assert.Equal(t, dbId1, dbId2)

	pool.Close()

	// new conn
	pool, err = database.New(context.Background(), hclog.NewNullLogger(), dbUrl)
	assert.NoError(t, err)

	dbId3, err := GetDatabaseId(context.Background(), pool)
	assert.NoError(t, err)
	assert.Equal(t, dbId1, dbId3)

	pool.Close()
}
