package policy

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	pg "github.com/bbernays/pg-commands"
)

func StoreOutput(query string, outputLocation string) {
	exc := pg.NewExec(&pg.Postgres{
		Host:     "localhost",
		Port:     5432,
		DB:       "postgres",
		Username: "postgres",
		Password: "pass",
	})
	exc.Query = query
	dumpExec := exc.Exec(pg.ExecOptions{StreamPrint: false})
	if dumpExec.Error != nil {
		fmt.Println(query)
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Println(query)
	}
}

func StoreSnapshot(path string, tables []string) {
	if len(tables) == 0 {
		return
	}

	dump := pg.NewDump(&pg.Postgres{
		Host:     "localhost",
		Port:     5432,
		DB:       "postgres",
		Username: "postgres",
		Password: "pass",
	})
	dump.Options = []string{"-a", "--column-inserts"}
	for _, table := range tables {
		dump.Options = append(dump.Options, "-t", table)
	}
	log.Println(path)
	dump.SetFileName(path + "/pg-dump.sql")
	dump.SetupFormat("plain")
	dump.SetPath("./")

	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: false})
	if dumpExec.Error != nil {

		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)

	}
}
func RestoreSnapshot(fileName string) {
	// dumpExec.File
	file := "postgres_1640202245.sql"
	pgConnection := pg.Postgres{
		Host:     "localhost",
		Port:     5432,
		DB:       "postgres",
		Username: "postgres",
		Password: "pass",
	}
	cmd := exec.Command("psql", "-U", pgConnection.Username, "-h", pgConnection.Host, "-d", pgConnection.DB, "-a", "-f", file)
	fmt.Println(cmd.Env)
	fmt.Println("psql", "-U", pgConnection.Username, "-h", pgConnection.Host, "-d", pgConnection.DB, "-a", "-f", file)

	cmd.Env = append(cmd.Env, "PGPASSWORD=pass")
	cmd.Env = append(cmd.Env, os.Environ()...)

	var out, stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}

}

// func (c ConnectionManager) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
// 	if strings.LastIndex(query, ";") > 0 {
// 		query = query[:strings.LastIndex(query, ";")]
// 	}

// 	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", query)
// 	rows, err := c.pool.Query(ctx, explainQuery)
// 	if err != nil {
// 		log.Println(explainQuery)
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		var s string
// 		if err := rows.Scan(&s); err != nil {
// 			log.Fatal(err)
// 		}
// 		var arrayJsonMap []map[string](interface{})
// 		err := json.Unmarshal([]byte(s), &arrayJsonMap)
// 		if err != nil {
// 			log.Printf("ERROR: fail to unmarshal json, %s", err.Error())
// 		}

// 		flat, err := flatten.Flatten(arrayJsonMap[0], "", flatten.DotStyle)
// 		if err != nil {
// 			log.Printf("ERROR: fail to flatten json, %s", err.Error())
// 		}
// 		for key, val := range flat {
// 			if strings.HasSuffix(key, "Relation Name") {
// 				tableNames = append(tableNames, val.(string))
// 			}

// 		}
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Println(query)
// 		log.Fatal(err)
// 	}

// 	return tableNames, err
// }
