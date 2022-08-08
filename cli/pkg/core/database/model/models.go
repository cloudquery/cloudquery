package model

import "time"

type DatabaseInfo struct {
	Version     string
	Uptime      time.Duration
	FullVersion string
}
