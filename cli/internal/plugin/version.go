package plugin

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/cli/internal/persistentdata"
	"github.com/cloudquery/cloudquery/cli/internal/versions"
	"github.com/hashicorp/go-version"
	"github.com/spf13/afero"
)

const DevelopmentVersion = "development"

const (
	// Timeout for http requests related to CloudQuery core version check.
	versionCheckHTTPTimeout = time.Second * 10
	cloudQueryGithubOrg     = "cloudquery"
	cloudQueryRepoName      = "cloudquery"
	lastUpdateCheckFileName = "last-update-check"

	// UpdateCheckPeriod specifies how much time in seconds must pass between
	// subsequent checks for CloudQuery core update availability.
	UpdateCheckPeriod = int64(23 * time.Hour / time.Second)
)

var (
	// Version variable is injected in build time
	Version = DevelopmentVersion
)

// unit tests helper
var getLatestVersion = doGetLatestVersion

func doGetLatestVersion(ctx context.Context, client *http.Client, owner, repo string) (string, error) {
	c := versions.NewClient()
	return c.GetLatestCLIRelease(ctx)
}

// CheckCoreUpdate checks if an update to CloudQuery core is available and returns its (new) version.
// To avoid making those network requests on each CLI invocation it stores last time and version seen on GitHub
// in a so-called "last update check" file. If there is an error then returned version is nil. To be specific,
// error is returned if:
//
// * core.Version is not a valid semantic version
// * if "last update check" file does not exist, and we fail to write to it
// * if "last update check" file has a single word "disable" in it
// * GitHub reports an error
// * last release version on GitHub is not a valid semantic version
//
// Otherwise, new available version is returned that is obtained from either:
//
// * "last update check" content and that version is newer than current
// * GitHub the latest release version, and it's newer than current
func CheckCoreUpdate(ctx context.Context, fs afero.Afero, nowUnix, period int64) (*version.Version, error) {
	if Version == DevelopmentVersion {
		// development version or something with local changes, skip update checks
		return nil, nil
	}

	currentVersion, err := version.NewSemver(Version)
	if err != nil {
		return nil, err
	}

	pd := persistentdata.New(fs, lastUpdateCheckFileName, func() string {
		return fmt.Sprintf("%d 0.0.0", nowUnix-period-1)
	})
	v, err := pd.Get()
	if err != nil {
		return nil, err
	}
	c := strings.TrimSpace(v.Content)
	if strings.HasPrefix(c, "disable") { // so both disable(d) work
		return nil, nil
	}
	fields := strings.Fields(c)
	var lastTime int64
	lastVersion, _ := version.NewSemver("0.0.0")
	if len(fields) >= 2 {
		lastTime, _ = strconv.ParseInt(fields[0], 10, 64)
		lastVersion, _ = version.NewSemver(fields[1])
	}
	if currentVersion.LessThan(lastVersion) {
		return lastVersion, nil
	}
	if nowUnix-lastTime > period {
		vn, err := getLatestVersion(ctx, &http.Client{Timeout: versionCheckHTTPTimeout}, cloudQueryGithubOrg, cloudQueryRepoName)
		if err != nil {
			return nil, err
		}
		newVersion, err := version.NewSemver(vn)
		if err != nil {
			// not really expected
			return nil, err
		}
		_ = v.Update(fmt.Sprintf("%d %s", nowUnix, newVersion))
		if currentVersion.LessThan(newVersion) {
			return newVersion, nil
		}
	}
	return nil, nil
}
