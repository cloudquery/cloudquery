package client_test

import (
	"fmt"
	"github.com/cloudquery/cloudquery/client"
	"github.com/cloudquery/cloudquery/config"
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
	cfg, err := config.Parse("./testdata/config.yml")
	if err != nil {
		t.Fatal(err)
	}
	c, err := client.New("", "", cfg)
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
		{"sqlite",
			"5",
			[]string{"MYSQL_ROOT_PASSWORD=pass", "MYSQL_DATABASE=dbname"},
			"sqlite",
			"%s",
			"cloudquery.db",
		},
		// This is commented out because looks like there is a bug that only happens in github actions https://github.com/hashicorp/go-plugin/issues/149
		//{"mysql",
		//	"5",
		//	[]string{"MYSQL_ROOT_PASSWORD=pass", "MYSQL_DATABASE=dbname"},
		//	"mysql",
		//"root:pass@tcp(127.0.0.1:%s)/dbname",
		//"3306/tcp",
		//},
		//{"postgres",
		//	"13",
		//	[]string{"POSTGRES_PASSWORD=pass"},
		//	"postgresql",
		//	"host=localhost user=postgres password=pass DB.name=postgres port=%s",
		//	"5432/tcp",
		//},
		//{"mcr.microsoft.com/mssql/server",
		//	"2019-latest",
		//	[]string{"SA_PASSWORD=yourStrong(!)Password", "ACCEPT_EULA=Y"},
		//	"sqlserver",
		//	"sqlserver://sa:yourStrong(!)Password@localhost:%s?database=master",
		//	"1433/tcp",
		//},
	}

	for _, tc := range tests {
		t.Run(tc.driver, func(t *testing.T) {
			var resource *dockertest.Resource
			var port string
			if tc.dockerName != "sqlite" {
				resource, err = pool.Run(tc.dockerName, tc.dockerVersion, tc.env)
				if err != nil {
					log.Fatalf("Could not start resource: %s", err)
				}
				time.Sleep(20 * time.Second)
				port = resource.GetPort(tc.port)
			}
			cfg, err := config.Parse("./testdata/config.yml")
			if err != nil {
				t.Fatal(err)
			}
			client, err := client.New(tc.driver, fmt.Sprintf(tc.dsn, port), cfg)
			if err != nil {
				t.Fatal(err)
			}

			testErr := client.Run()

			if tc.dockerName != "sqlite" {
				if err := pool.Purge(resource); err != nil {
					log.Fatalf("Could not purge resource: %s", err)
				}
			}

			if testErr != nil {
				t.Fatal(testErr)
			}
		})
	}

}
