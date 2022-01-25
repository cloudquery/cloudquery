package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	pg "github.com/bbernays/pg-commands"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jeremywohl/flatten"
)

func StoreSnapshot(path string, tables []string, config *pgxpool.Config) error {
	if len(tables) == 0 {
		return errors.New("no tables to snapshot")
	}

	dump := pg.NewDump(&pg.Postgres{
		Host:     config.ConnConfig.Host,
		Port:     int(config.ConnConfig.Port),
		DB:       config.ConnConfig.Database,
		Username: config.ConnConfig.User,
		Password: config.ConnConfig.Password,
	})
	dump.Options = []string{}
	for _, table := range tables {
		dump.Options = append(dump.Options, "-t", table)
	}

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
			e.log.Error("error scanning into variable", "error", err)
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
			if strings.HasSuffix(key, "Alias") {
				tableNames = append(tableNames, val.(string))
			}

		}
	}
	if err := rows.Err(); err != nil {
		e.log.Error("Error fetching rows", "query", query, "error", err)
		return tableNames, err
	}

	return tableNames, err
}

func (e *Executor) StoreOutput(ctx context.Context, pol *Policy, destination string, config *pgxpool.Config) (err error) {

	queries := []string{
		`SET client_encoding='UTF8'`,
	}

	for _, v := range pol.Views {
		queries = append(queries, fmt.Sprintf("CREATE OR REPLACE TEMPORARY VIEW %s AS %s", v.Name, v.Query))
	}

	queries = append(queries, fmt.Sprintf("\\COPY (SELECT json_agg(foo)::jsonb FROM (%s) foo ) TO '%s'", cleanQuery(pol.Checks[0].Query), destination+"/"+"data.json"))
	pgConnection := pg.NewDump(&pg.Postgres{
		Host:     config.ConnConfig.Host,
		Port:     int(config.ConnConfig.Port),
		DB:       config.ConnConfig.Database,
		Username: config.ConnConfig.User,
		Password: config.ConnConfig.Password,
	})
	// construct arguments
	args := []string{"-U", pgConnection.Username, "-h", pgConnection.Host, "-d", pgConnection.DB}
	for _, q := range queries {
		args = append(args, "-c", q)
	}

	// Execute psql command
	cmd := exec.Command("psql", args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf(`PGPASSWORD=%v`, pgConnection.Password))

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		e.log.Error("StandardError", "Output", stdoutStderr)
		return err
	}

	return nil
}
