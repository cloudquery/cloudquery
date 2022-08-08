//go:build darwin || linux

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

// In this entire file, by 'ulimit' we mean the "num open files ulimit" (RLIMIT_NOFILE)
func checkAndSetUlimitUnix() {
	logger := zerolog.Logger

	rlimit, err := getUlimit()
	if err != nil {
		logger.Err(err).Msg("checkAndSetUlimitUnix: failed getting ulimit")
		return
	}
	logger.Info().Uint64("hard_ulimit", rlimit.Max).Uint64("soft_ulimit", rlimit.Cur).Msg("limits (before adjustment)")

	if err := setUlimit(); err != nil {
		logger.Err(err).Msg("failed setting ulimit")
	}

	rlimit, err = getUlimit()
	if err != nil {
		logger.Err(err).Msg("checkAndSetUlimitUnix: failed getting ulimit")
		return
	}
	logger.Info().Uint64("hard_ulimit", rlimit.Max).Uint64("soft_ulimit", rlimit.Cur).Msg("limits (after adjustment)")
}

func setUlimit() error {
	logger := zerolog.Logger

	// Setting the hard ulimit is very likely to fail (usually requires root or even more arcane permissions).
	// If it fails we just log at "INFO" level.
	if err := setHardUlimit(); err != nil {
		// Just an info log here, since it's really not actionable.
		logger.Info().AnErr("err", err).Msg("failed setting hard ulimit")
	}

	return setSoftUlimit()
}

func setHardUlimit() error {
	logger := zerolog.Logger
	rLimit, err := getUlimit()
	if err != nil {
		return fmt.Errorf("setHardUlimit: error getting ulimit: %w", err)
	}

	if rLimit.Max < ulimitUnix {
		logger.Info().Uint64("previous_ulimit", rLimit.Max).Uint64("new_ulimit", ulimitUnix).Msg("adjusting hard ulimit")
		rLimit.Max = ulimitUnix
		return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	}

	return nil
}

func setSoftUlimit() error {
	logger := zerolog.Logger

	rlimit, err := getUlimit()
	if err != nil {
		return fmt.Errorf("setSoftUlimit: error getting ulimit: %w", err)
	}

	desiredUlimit := min(rlimit.Max, ulimitUnix)
	if rlimit.Cur < desiredUlimit {
		logger.Info().Uint64("previous_ulimit", rlimit.Cur).Uint64("new_ulimit", desiredUlimit).Msg("adjusting soft ulimit")
		rlimit.Cur = desiredUlimit
		return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	}

	return nil
}

func getUlimit() (syscall.Rlimit, error) {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return rLimit, err
}

// Golang doesn't have an out-of-the-box min function xd
func min(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}
