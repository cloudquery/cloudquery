package testing

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
