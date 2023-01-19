package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/scaffold/cmd"
)

func main() {
	if err := cmd.NewCmdRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
