// +build integration

package cloudqueryclient

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

func TestMigrationSQLite(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "*.cloudquery.db")
	if err != nil {
		t.Fatal(err)
	}
	if err = tmpFile.Close(); err != nil {
		t.Fatal(err)
	}
	client, err := New("sqlite", tmpFile.Name(), true)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Run("./testdata/config.yml")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMigrationSQLServers(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	tests := []struct {
		dockerName string
		dockerVersion string
		env []string
		driver string
		dsn string
		port string
	}{
		{"mysql",
			"5",
			[]string{"MYSQL_ROOT_PASSWORD=pass", "MYSQL_DATABASE=dbname"},
			"mysql",
		"root:pass@tcp(127.0.0.1:%s)/dbname",
		"3306/tcp",
		},
		{"postgres",
			"13",
			[]string{"POSTGRES_PASSWORD=pass"},
			"postgresql",
			"host=localhost user=postgres password=pass DB.name=postgres port=%s",
			"5432/tcp",
		},
		{"mcr.microsoft.com/mssql/server",
			"2019-latest",
			[]string{"SA_PASSWORD=yourStrong(!)Password", "ACCEPT_EULA=Y"},
			"sqlserver",
			"sqlserver://sa:yourStrong(!)Password@localhost:%s?database=master",
			"1433/tcp",
		},
	}

	for _, tc := range tests {
		t.Run(tc.driver, func(t *testing.T) {
			resource, err := pool.Run(tc.dockerName, tc.dockerVersion, tc.env)
			if err != nil {
				log.Fatalf("Could not start resource: %s", err)
			}
			time.Sleep(20 * time.Second)
			port := resource.GetPort(tc.port)
			client, err := New(tc.driver, fmt.Sprintf(tc.dsn, port), true)
			if err != nil {
				t.Fatal(err)
			}
			testErr := client.Run("./testdata/config.yml")

			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s", err)
			}

			if testErr != nil {
				t.Fatal(err)
			}
		})
	}

}
