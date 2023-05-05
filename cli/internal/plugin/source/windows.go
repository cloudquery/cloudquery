//go:build windows

package source

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		// launch as new process group so that signals are not sent to the child process
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, // windows
	}
}

func (c *Client) terminateProcess() error {
	c.logger.Debug().Msg("sending kill signal to destination plugin")
	if err := c.cmd.Process.Kill(); err != nil {
		c.logger.Error().Err(err).Msg("failed to kill source plugin")
	}
	c.logger.Info().Msg("waiting for source plugin to terminate")
	st, err := c.cmd.Process.Wait()
	if err != nil {
		return err
	}
	if !st.Success() {
		// on windows there is no way to shutdown gracefully via signal. Maybe we can do it via grpc api?
		// though it is a bit strange to expose api to shutdown a server :thinking?:
		c.logger.Info().Msgf("source plugin process exited with %s", st.String())
	}

	return nil
}

