package getter

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/go-getter"
	"github.com/rs/zerolog/log"
)

var (
	detectors = []getter.Detector{
		new(GitHubDetector),
		new(getter.GitDetector),
		new(getter.S3Detector),
		new(getter.GCSDetector),
		new(HubDetector),
		new(fileDetector),
	}

	detectorsMap = map[string]getter.Detector{
		"github": new(getter.GitHubDetector),
		"git":    new(getter.GitDetector),
		"s3":     new(getter.S3Detector),
		"gcs":    new(getter.GCSDetector),
		"hub":    new(HubDetector),
		"local":  new(fileDetector),
	}

	decompressors = map[string]getter.Decompressor{
		"bz2": new(getter.Bzip2Decompressor),
		"gz":  new(getter.GzipDecompressor),
		"xz":  new(getter.XzDecompressor),
		"zip": new(getter.ZipDecompressor),

		"tar.bz2":  new(getter.TarBzip2Decompressor),
		"tar.tbz2": new(getter.TarBzip2Decompressor),

		"tar.gz": new(getter.TarGzipDecompressor),
		"tgz":    new(getter.TarGzipDecompressor),

		"tar.xz": new(getter.TarXzDecompressor),
		"txz":    new(getter.TarXzDecompressor),
	}

	getters = map[string]getter.Getter{
		"file": new(getter.FileGetter),
		"gcs":  new(getter.GCSGetter),
		"git":  new(getter.GitGetter),
		"hg":   new(getter.HgGetter),
		"s3":   new(getter.S3Getter),
	}
)

func Get(ctx context.Context, installPath, url string, options ...getter.ClientOption) error {
	log.Debug().Str("url", url).Msg("getting source from url")
	pwd, _ := os.Getwd()
	client := getter.Client{
		Src:           url,
		Dst:           installPath,
		Pwd:           pwd,
		Mode:          getter.ClientModeDir,
		Detectors:     detectors,
		Decompressors: decompressors,
		Getters:       getters,
		Ctx:           ctx,
		// Extra options provided by caller to overwrite default behavior
		Options: options,
	}

	if err := client.Get(); err != nil {
		return err
	}
	// If we get down here then we've either downloaded the package or
	// copied a previous tree we downloaded, and so either way we should
	// have got the full module package structure written into instPath.
	return nil
}

func DetectType(src string) (string, string, bool, error) {
	forcedProtocol := strings.Split(src, "::")
	if len(forcedProtocol) > 1 {
		if _, ok := detectorsMap[forcedProtocol[0]]; ok {
			return forcedProtocol[0], src, true, nil
		}
	}

	pwd, _ := os.Getwd()
	for t, d := range detectorsMap {
		source, found, err := d.Detect(src, pwd)
		if err != nil {
			return "", source, false, fmt.Errorf("failed to detect url %s: %w", src, err)
		}
		if found {
			return t, source, true, nil
		}
	}
	return "", src, false, nil
}
