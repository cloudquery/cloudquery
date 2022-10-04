package cmd

import (
	"github.com/spf13/cobra"
	"strings"
)

type prettyError struct {
	err error
}

type command func(*cobra.Command, []string) error

func commandWithPrettyErrors(original command) command {
	return func(cmd *cobra.Command, args []string) error {
		err := original(cmd, args)
		return prettifyError(err)
	}
}

func (p prettyError) Error() string {
	msg := p.err.Error()
	return strings.Join(strings.Split(msg, ": "), ":\n  ")
}

func prettifyError(err error) error {
	if err == nil {
		return err
	}
	return prettyError{err: err}
}
