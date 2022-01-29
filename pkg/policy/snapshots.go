package policy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	pg "github.com/bbernays/pg-commands"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jeremywohl/flatten"
)

type CliExecutor struct {
	// Connection to the database
	exec   *Executor
	config *pgx.ConnConfig
	log    hclog.Logger
}

func NewCliExecutor(exec *Executor, config *pgx.ConnConfig) *CliExecutor {
	return &CliExecutor{
		exec:   exec,
		log:    exec.log,
		config: config,
	}
}

func (ce *CliExecutor) StoreSnapshot(path string, tables []string, config *pgxpool.Config) error {
	if len(tables) == 0 {
		return errors.New("no tables to snapshot")
	}

	dump := pg.NewDump(&pg.Postgres{
		Host:     ce.config.Host,
		Port:     int(ce.config.Port),
		DB:       ce.config.Database,
		Username: ce.config.User,
		Password: ce.config.Password,
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

func (ce *CliExecutor) RestoreSnapshot(tx context.Context, source string, config *pgxpool.Config) error {

	args := []string{"-U", config.ConnConfig.User, "-h", config.ConnConfig.Host, "-d", config.ConnConfig.Database, "-p", fmt.Sprint(config.ConnConfig.Port)}
	ce.log.Debug("cli args", "args", args)
	// Execute psql command
	cmd := exec.Command("psql", args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		dat, err := os.ReadFile(source)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		io.WriteString(stdin, string(dat))
	}()

	cmd.Env = append(os.Environ(), fmt.Sprintf(`PGPASSWORD=%v`, config.ConnConfig.Password))

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		ce.log.Error("StandardError", "Output", string(stdoutStderr))
		return err
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
func (ce *CliExecutor) ExtractTableNames(ctx context.Context, query string) (tableNames []string, err error) {
	ce.log.Debug("extracting Table names-raw", "raw query", query)
	cleanedQuery := cleanQuery(query)
	ce.log.Debug("extracting Table names-cleaned", "cleaned query", cleanedQuery)
	explainQuery := fmt.Sprintf("EXPLAIN (FORMAT JSON) %s", cleanedQuery)

	rows, err := ce.exec.conn.Query(ctx, explainQuery)
	if err != nil {
		return tableNames, err
	}
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			ce.log.Error("error scanning into variable", "error", err)
		}
		var arrayJsonMap []map[string](interface{})
		err := json.Unmarshal([]byte(s), &arrayJsonMap)
		if err != nil {
			ce.log.Error("failed to unmarshal json", "err", err)
			return tableNames, err
		}

		flat, err := flatten.Flatten(arrayJsonMap[0], "", flatten.DotStyle)
		if err != nil {
			ce.log.Error("failed to flatten json", "err", err)
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
		ce.log.Error("Error fetching rows", "query", query, "error", err)
		return tableNames, err
	}

	return tableNames, err
}

func (ce *CliExecutor) StoreOutput(ctx context.Context, pol *Policy, destination string, config *pgxpool.Config) (err error) {
	if !pol.HasChecks() {
		return errors.New("no checks")
	}

	queries := []string{
		`SET client_encoding='UTF8'`,
	}

	for _, v := range pol.Views {
		queries = append(queries, fmt.Sprintf("CREATE OR REPLACE TEMPORARY VIEW %s AS %s", v.Name, v.Query))
	}

	queries = append(queries, fmt.Sprintf("\\COPY (SELECT COALESCE(JSON_AGG(FOO)::JSONB, '[]'::JSONB) FROM (%s) foo)  TO '%s'", cleanQuery(pol.Checks[0].Query), destination+"/"+"data.json"))

	ce.log.Debug("destination of query", "dest", destination+"/"+"data.json")
	// construct arguments
	args := []string{"-U", config.ConnConfig.User, "-h", config.ConnConfig.Host, "-d", config.ConnConfig.Database, "-p", fmt.Sprint(config.ConnConfig.Port)}
	for _, q := range queries {
		args = append(args, "-c", q)
	}

	// Execute psql command
	cmd := exec.Command("psql", args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf(`PGPASSWORD=%v`, config.ConnConfig.Password))

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		ce.log.Error("StandardError", "Output", string(stdoutStderr))
		return err
	}

	return nil
}
