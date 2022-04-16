package client

import (
	"context"
	"errors"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/internal/test/providertest"
	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/serve"
	"github.com/fsnotify/fsnotify"
	"github.com/hashicorp/go-version"
	"github.com/jackc/pgx/v4"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_normalizeResources(t *testing.T) {
	tests := []struct {
		name      string
		requested []string
		all       map[string]*schema.Table
		want      []string
		wantErr   bool
	}{
		{
			"wilcard",
			[]string{"*"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2", "3"},
			false,
		},
		{
			"wilcard with explicit",
			[]string{"*", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"unknown resource",
			[]string{"1", "2", "x"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"duplicate resource",
			[]string{"1", "2", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			nil,
			true,
		},
		{
			"ok, all explicit",
			[]string{"2", "1"},
			map[string]*schema.Table{"3": nil, "2": nil, "1": nil},
			[]string{"1", "2"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeResources(tt.requested, tt.all)
			if (err != nil) != tt.wantErr {
				t.Errorf("doInterpolate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doInterpolate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collectProviderVersions(t *testing.T) {
	tests := []struct {
		name       string
		providers  []*config.RequiredProvider
		getVersion func(providerName string) (string, error)
		want       map[string]*version.Version
		wantErr    bool
	}{
		{
			"no required providers",
			nil,
			func(providerName string) (string, error) { panic("test") },
			map[string]*version.Version{},
			false,
		},
		{
			"failed to get a version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "1.0", errors.New("test") },
			nil,
			true,
		},
		{
			"failed to parse version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "xyz", nil },
			nil,
			true,
		},
		{
			"failed to parse version",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) { return "xyz", nil },
			nil,
			true,
		},
		{
			"ok",
			[]*config.RequiredProvider{{Name: "aws"}, {Name: "gcp"}},
			func(providerName string) (string, error) {
				if providerName == "aws" {
					return "1.2.3", nil
				} else {
					return "v4.5.6", nil
				}
			},
			map[string]*version.Version{
				"aws": version.Must(version.NewVersion("1.2.3")),
				"gcp": version.Must(version.NewVersion("v4.5.6")),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collectProviderVersions(tt.providers, tt.getVersion)
			require.Equal(t, tt.wantErr, err != nil, "collectProviderVersions() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got, "collectProviderVersions() = %v, want %v", got, tt.want)
		})
	}
}
