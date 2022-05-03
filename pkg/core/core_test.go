package core

import (
	"context"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setupDB(t *testing.T) (dsn string) {
	baseDSN := os.Getenv("CQ_CLIENT_TEST_DSN")
	if baseDSN == "" {
		baseDSN = "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}

	conn, err := pgx.Connect(context.Background(), baseDSN)
	if err != nil {
		assert.FailNow(t, "failed to create connection")
		return
	}

	newDB := "test_" + strconv.Itoa(rand.Int())

	_, err = conn.Exec(context.Background(), "CREATE DATABASE "+newDB)
	assert.NoError(t, err)

	t.Cleanup(func() {
		defer conn.Close(context.Background())

		if os.Getenv("CQ_TEST_DEBUG") != "" && t.Failed() {
			t.Log("Not dropping database", newDB)
			return
		}

		if _, err := conn.Exec(context.Background(), "DROP DATABASE "+newDB+" WITH(FORCE)"); err != nil {
			t.Logf("teardown: drop database failed: %v", err)
		}
	})

	return strings.Replace(baseDSN, "/postgres?", "/"+newDB+"?", 1)
}
