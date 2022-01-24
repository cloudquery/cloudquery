package policy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	pg "github.com/bbernays/pg-commands"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jeremywohl/flatten"
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

func StoreSnapshot(path string, tables []string, dsn string) error {
	if len(tables) == 0 {
		return errors.New("no tables to snapshot")
	}
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}
	// config, err := ParseConfig(dsn)
	dump := pg.NewDump(&pg.Postgres{
		Host:     config.ConnConfig.Host,
		Port:     int(config.ConnConfig.Port),
		DB:       config.ConnConfig.Database,
		Username: config.ConnConfig.User,
		Password: config.ConnConfig.Password,
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
		return errors.New("error dumping tables")
	}
	return nil
}
func RestoreSnapshot(fileName, dsn string) error {
	// dumpExec.File
	file := "postgres_1640202245.sql"
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}
	// config, err := ParseConfig(dsn)
	pgConnection := pg.NewDump(&pg.Postgres{
		Host:     config.ConnConfig.Host,
		Port:     int(config.ConnConfig.Port),
		DB:       config.ConnConfig.Database,
		Username: config.ConnConfig.User,
		Password: config.ConnConfig.Password,
	})
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

func cleanQuery(query string) string {
	var re = regexp.MustCompile(`(?s)\/\*.*?\*\/|--.*?\n`)
	query = re.ReplaceAllString(query, "")
	query = strings.TrimSuffix(query, ";")
	query = strings.TrimSpace(query)
	query = strings.TrimSuffix(query, ";")

	return strings.TrimSpace(query)
}
func (e *Executor) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
	e.log.Debug("extracting Table names-raw", "raw query", query)
	cleanedQuery := cleanQuery(query)
	e.log.Debug("extracting Table names-cleaned", "cleaned query", cleanedQuery)
	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", cleanedQuery)

	rows, err := e.conn.Query(ctx, explainQuery)
	if err != nil {
		return tableNames, err
	}

	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}
		var arrayJsonMap []map[string](interface{})
		err := json.Unmarshal([]byte(s), &arrayJsonMap)
		if err != nil {
			e.log.Error("failed to unmarshal json", "err", err)
			return tableNames, err
		}

		flat, err := flatten.Flatten(arrayJsonMap[0], "", flatten.DotStyle)
		if err != nil {
			e.log.Error("failed to flatten json", "err", err)
			return tableNames, err
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
