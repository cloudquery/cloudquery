//go:build !windows

package manageddestination

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		// launch as new process group so that signals (ex: SIGINT) are not sent to the child process
		Setpgid: true, // UNIX systems
	}
}

func (c *Client) terminateProcess() error {
	c.logger.Debug().Msg("sending interrupt signal to source plugin")
	if err := c.cmd.Process.Signal(os.Interrupt); err != nil {
		c.logger.Error().Err(err).Msg("failed to send interrupt signal to source plugin")
	}
	timer := time.AfterFunc(5*time.Second, func() {
		c.logger.Info().Msg("sending kill signal to source plugin")
		if err := c.cmd.Process.Kill(); err != nil {
			c.logger.Error().Err(err).Msg("failed to kill source plugin")
		}
	})
	c.logger.Info().Msg("waiting for source plugin to terminate")
	st, err := c.cmd.Process.Wait()
	timer.Stop()
	if err != nil {
		return err
	}
	if !st.Success() {
		var additionalInfo string
		status := st.Sys().(syscall.WaitStatus)
		if status.Signaled() && st.ExitCode() != -1 {
			additionalInfo += fmt.Sprintf(" (exit code: %d)", st.ExitCode())
		}
		if st.ExitCode() == 137 {
			additionalInfo = " (Out of Memory)"
		}
		return fmt.Errorf("source plugin process failed with %s%s", st.String(), additionalInfo)
	}

	return nil
}
