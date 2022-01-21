package testing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jeremywohl/flatten"
)

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strings"

// 	"github.com/cloudquery/cloudquery/pkg/policy"

// 	"github.com/jackc/pgx/v4/pgxpool"
// 	"github.com/jeremywohl/flatten"
// 	uuid "github.com/satori/go.uuid"
// )

// type ConnectionManager struct {
// 	pool *pgxpool.Pool
// }

// func New(psqlInfo string) *ConnectionManager {
// 	// pool, err := cq.Client.CreateDatabase(context.Background(), psqlInfo)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	var pool *pgxpool.Pool
// 	return &ConnectionManager{pool: pool}
// }

// func (c ConnectionManager) DumpAllPolicies(ctx context.Context, policies policy.Policies, path string) error {
// 	return c.TraversePolicies(ctx, policies, path)
// }

// func AggregatePolicies(aggregatedPolicies map[string]string, policies policy.Policies, path string) map[string]string {
// 	for _, policy := range policies {
// 		pathLocal := fmt.Sprintf("%s/%s", path, policy.Name)

// 		for _, query := range policy.Checks {
// 			queryPath := fmt.Sprintf("%s/query/%s", pathLocal, query.Name)
// 			if _, ok := aggregatedPolicies[queryPath]; !ok {
// 				aggregatedPolicies[queryPath] = query.Query
// 			}

// 		}
// 		AggregatePolicies(aggregatedPolicies, policy.Policies, pathLocal)
// 	}
// 	return aggregatedPolicies
// }

// func (c ConnectionManager) TraversePolicies(ctx context.Context, policies policy.Policies, path string) error {
// 	for _, policy := range policies {
// 		pathLocal := fmt.Sprintf("%s/%s", path, policy.Name)
// 		for _, view := range policy.Views {
// 			c.CreateView(context.Background(), view.Name, view.Query)
// 		}

// 		for _, query := range policy.Checks {
// 			c.HandleQuery(context.Background(), query, fmt.Sprintf("%s/query/%s", pathLocal, query.Name))
// 		}
// 		c.TraversePolicies(ctx, policy.Policies, pathLocal)
// 	}
// 	return nil
// }

// func (c ConnectionManager) HandleQuery(ctx context.Context, query *policy.Check, path string) error {

// 	path = createPath(path, query.Name)
// 	// tables, _ := c.ExtractTableNames(context.Background(), query.Query)
// 	// StoreSnapshot(path, tables)
// 	fullPath := fmt.Sprintf("%s/queryResults", path)
// 	create(fullPath)

// 	q := fmt.Sprintf("\\copy (%s) TO '%s' with csv", cleanQuery(query.Query), fullPath)
// 	StoreOutput(q, path)
// 	return nil
// }

// func (c ConnectionManager) CreateView(ctx context.Context, name, query string) (err error) {
// 	// log.Printf("creating view: %s", name)
// 	fullQuery := fmt.Sprintf("CREATE OR REPLACE VIEW %s AS %s", name, query)

// 	_, err = c.pool.Query(ctx, fullQuery)
// 	if err != nil {
// 		log.Println(fullQuery)
// 		log.Fatal(err)
// 	}

// 	return err
// }

// func (c ConnectionManager) GetOutput(ctx context.Context, query, resultsPath string) (err error) {
// 	copyQuery := fmt.Sprintf("COPY (%s) TO './PolicyTesting/database-data%s';", strings.TrimSuffix(query, ";"), strings.ReplaceAll(resultsPath, "/query/", "/")+"results")
// 	log.Println(copyQuery)
// 	resp, err := c.pool.Exec(ctx, copyQuery)
// 	if err != nil {
// 		log.Println(copyQuery)
// 		log.Fatal(err)
// 	}
// 	log.Println(resp)
// 	return err

// }

func (c ConnectionManager) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
	if strings.LastIndex(query, ";") > 0 {
		query = query[:strings.LastIndex(query, ";")]
	}

	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", query)
	rows, err := c.pool.Query(ctx, explainQuery)
	if err != nil {
		log.Println(explainQuery)
		log.Fatal(err)
	}

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}
		var arrayJsonMap []map[string](interface{})
		err := json.Unmarshal([]byte(s), &arrayJsonMap)
		if err != nil {
			log.Printf("ERROR: fail to unmarshal json, %s", err.Error())
		}

		flat, err := flatten.Flatten(arrayJsonMap[0], "", flatten.DotStyle)
		if err != nil {
			log.Printf("ERROR: fail to flatten json, %s", err.Error())
		}
		for key, val := range flat {
			if strings.HasSuffix(key, "Relation Name") {
				tableNames = append(tableNames, val.(string))
			}

		}
	}
	if err := rows.Err(); err != nil {
		log.Println(query)
		log.Fatal(err)
	}

	return tableNames, err
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func cleanQuery(query string) string {

	var re = regexp.MustCompile(`(?s)\/\*.*?\*\/|--.*?\n`)
	query = re.ReplaceAllString(query, "")

	// query = strings.ReplaceAll(query, "\n", " ")
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")

	return strings.TrimSpace(query)
}

func createPath(path, queryName string) string {
	if !strings.HasPrefix(path, "/") {
		path = fmt.Sprintf("/%s", path)
	}
	path = strings.TrimSuffix(path, "/")

	path = strings.TrimSuffix(path, "/query/"+queryName)

	u2 := uuid.NewV4()

	cleanedPath := filepath.Join("./database-data/", path, "/query"+"-"+queryName+"/", "tests", u2.String())

	// generate test name (uuid)
	err := os.MkdirAll(cleanedPath, os.ModePerm)
	if err != nil {
		log.Printf("%+v\n", err)
	}
	return cleanedPath
}
