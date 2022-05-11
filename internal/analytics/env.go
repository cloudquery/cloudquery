package analytics

import (
	"net"
	"os"
	"runtime"
	"sort"
	"strings"
)

type Environment struct {
	// OS Client executing
	OS string `json:"os"`
	// Terminal is true if Client is executed inside a terminal environment
	Terminal bool `json:"terminal"`
	// CI is true if the Client is executed inside a CI i.e github actions etc'
	CI bool `json:"ci"`
	// FaaS is true if Client is executed inside a function i.e lambda etc'
	FaaS bool `json:"faas"`
	// Hashed hostname identifier
	Hostname string `json:"hostname"`
	MacAddr  string `json:"mac_addr"`
}

func getEnvironmentAttributes(terminal bool) *Environment {
	hn, _ := os.Hostname()
	if hn != "" {
		hn = HashAttribute(hn)
	}
	return &Environment{
		OS:       runtime.GOOS,
		Terminal: terminal,
		CI:       IsCI(),
		FaaS:     IsFaaS(),
		Hostname: hn,
		MacAddr:  macHost(),
	}
}

// macHost will extract MAC addresses, add the hostname and return a hash
func macHost() string {
	ifas, err := net.Interfaces()
	if err != nil {
		return ""
	}
	as := make([]string, 0, len(ifas)+1)
	for _, ifa := range ifas {
		if a := ifa.HardwareAddr.String(); a != "" {
			as = append(as, a)
		}
	}
	sort.Strings(as)
	if hn, err := os.Hostname(); err == nil && hn != "" {
		as = append(as, hn)
	}
	return HashAttribute(strings.Join(as, ","))
}

// IsCI determines if we're running under a CI env by checking CI-specific env vars
func IsCI() bool {
	for _, v := range []string{
		"CI", "BUILD_ID", "BUILDKITE", "CIRCLECI", "CIRCLE_CI", "CIRRUS_CI", "CODEBUILD_BUILD_ID", "GITHUB_ACTIONS", "GITLAB_CI", "HEROKU_TEST_RUN_ID", "TEAMCITY_VERSION", "TF_BUILD", "TRAVIS",
	} {
		if os.Getenv(v) != "" {
			return true
		}
	}

	return false
}

// IsFaaS determines if we're running under a Lambda env by checking Lambda-specific env vars
func IsFaaS() bool {
	for _, v := range []string{
		"LAMBDA_TASK_ROOT", "AWS_LAMBDA_FUNCTION_NAME", // AWS
		"FUNCTION_TARGET",             // GCP
		"AZURE_FUNCTIONS_ENVIRONMENT", // Azure
	} {
		if os.Getenv(v) != "" {
			return true
		}
	}

	return false
}
