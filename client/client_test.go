package client_test

import (
	"fmt"
	"github.com/cloudquery/cloudquery/client"
	"github.com/ory/dockertest/v3"
	"log"
	"testing"
	"time"
)

func TestMigrationSQLServers(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	c, err := client.New("./testdata/config.yml", "", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Initialize(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		dockerName    string
		dockerVersion string
		env           []string
		driver        string
		dsn           string
		port          string
	}{
		{"postgres",
			"13",
			[]string{"POSTGRES_PASSWORD=pass"},
			"postgresql",
			"host=localhost user=postgres password=pass DB.name=postgres port=%s",
			"5432/tcp",
		},
	}

	for _, tc := range tests {
		t.Run(tc.driver, func(t *testing.T) {
			var resource *dockertest.Resource
			var port string
			resource, err = pool.Run(tc.dockerName, tc.dockerVersion, tc.env)
			if err != nil {
				log.Fatalf("Could not start resource: %s", err)
			}
			time.Sleep(20 * time.Second)
			port = resource.GetPort(tc.port)

			client, err := client.New("./testdata/config.yml", tc.driver, fmt.Sprintf(tc.dsn, port))
			if err != nil {
				t.Fatal(err)
			}

			testErr := client.Run("./testdata/config.yml")

			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s", err)
			}

			if testErr != nil {
				t.Fatal(testErr)
			}
		})
	}

}
