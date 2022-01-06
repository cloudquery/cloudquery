package getter

import (
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

func NormalizePath(src string) string {
	srcParts := strings.Split(src, "::")
	src = srcParts[0]
	if len(srcParts) > 1 {
		src = srcParts[1]
	}
	_url, err := url.Parse(src)
	if err == nil {
		src = filepath.Join(_url.Host, _url.Path)
	}
	srcParts = strings.Split(src, "@")
	if len(srcParts) > 1 {
		src = srcParts[0]
	}
	return filepath.ToSlash(strings.TrimRight(src, filepath.Ext(src)))
}

// ParseSourceSubPolicy takes a source URL and returns a tuple of the URL without
// the subdir and the subdir.
//
// ex:
//   dom.com/path/?q=p               => dom.com/path/?q=p, ""
//   proto://dom.com/path//*?q=p     => proto://dom.com/path?q=p, "*"
//   proto://dom.com/path//path2?q=p => proto://dom.com/path?q=p, "path2"
//   dom.com/path@q=p               => dom.com/path/?q=p, ""
//
func ParseSourceSubPolicy(src string) (string, string) {

	// URL might contains another url in query parameters
	stop := len(src)
	if idx := strings.Index(src, "?"); idx > -1 {
		stop = idx
	}
	// URL might contain query parameters
	if idx := strings.Index(src, "@"); idx > -1 {
		stop = idx
	}

	// Calculate an offset to avoid accidentally marking the scheme
	// as the dir.
	var offset int
	if idx := strings.Index(src[:stop], "://"); idx > -1 {
		offset = idx + 3
	}

	// First see if we even have an explicit subdir
	idx := strings.Index(src[offset:stop], "//")
	if idx == -1 {
		return src, ""
	}

	idx += offset
	subdir := src[idx+2:]
	src = src[:idx]

	// Next, check if we have query parameters and push them onto the
	// URL.
	if idx = strings.Index(subdir, "?"); idx > -1 {
		query := subdir[idx:]
		subdir = subdir[:idx]
		src += query
	}
	// Next, check if we have special parameters and push them onto the
	// URL.
	if idx = strings.Index(subdir, "@"); idx > -1 {
		query := subdir[idx:]
		subdir = subdir[:idx]
		src += query
	}

	if subdir != "" {
		subdir = path.Clean(subdir)
	}

	return src, subdir
}
