package client

import (
	"context"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/internal/test/providertest"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/fsnotify/fsnotify"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
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

func setupTestPlugin(t *testing.T) context.CancelFunc {
	debugCtx, cancelServe := context.WithCancel(context.Background())
	dir, _ := os.Getwd()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatal(err)
	}
	if err := watcher.Add(dir); err != nil {
		t.Fatal(err)
	}
	defer watcher.Close()

	go providertest.ServeTestPlugin(debugCtx)
	_ = os.Setenv("CQ_REATTACH_PROVIDERS", filepath.Join(dir, ".cq_reattach"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CQ")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	<-watcher.Events

	unmanaged, err := serve.ParseReattachProviders(os.Getenv("CQ_REATTACH_PROVIDERS"))
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range unmanaged {
		_, err := net.DialTimeout(u.Addr.Network(), u.Addr.String(), time.Second*5)
		if err != nil {
			t.Fatal(err)
		}
	}

	return cancelServe
}
