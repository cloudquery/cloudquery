//go:build darwin || linux
// +build darwin linux

package cmd

import (
	"fmt"
	"syscall"

	zerolog "github.com/rs/zerolog/log"
)

const ulimitUnix uint64 = 16384

func init() {
	fileDescriptorF = checkAndSetUlimitUnix
}

func checkAndSetUlimitUnix() {
	logger := zerolog.Logger
	if err := setUlimit(ulimitUnix); err != nil {
		logger.Err(fmt.Errorf("error setting ulimit: %w", err))
	}
}

func setUlimit(ulimit uint64) error {
	logger := zerolog.Logger
	rLimit, err := getUlimit()
	if err != nil {
		logger.Err(fmt.Errorf("error getting ulimit: %w", err))
	}
	if rLimit.Max < ulimit {
		logger.Debug().Uint64("previous_ulimit", rLimit.Max).Uint64("new_ulimit", ulimit).Msg("adjusting max ulimit")
		rLimit.Max = ulimit
	}
	logger.Debug().Uint64("previous_ulimit", rLimit.Cur).Uint64("new_ulimit", ulimit).Msg("adjusting current ulimit")
	rLimit.Cur = ulimit
	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}

func getUlimit() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return rLimit, err
}
