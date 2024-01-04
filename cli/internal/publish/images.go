package publish

import (
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/internal/hub"
	"golang.org/x/sync/errgroup"
)

type imageReference struct {
	mdFull    string
	mdPartial string
	absFile   string
	url       string
}

func processDocumentImages(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, docDir, contents string) (string, error) {
	ims, err := findMarkdownImages(contents, docDir)
	if err != nil {
		return "", err
	}
	if len(ims) == 0 {
		return contents, nil
	}

	fmt.Println("Preparing to upload images...")

	reqs := make([]cloudquery_api.TeamImageCreate, len(ims))
	for k := range ims {
		refParts := strings.SplitN(k, ":", 2)
		checksum, name := refParts[0], refParts[1]

		reqs = append(reqs, cloudquery_api.TeamImageCreate{
			Name:     name,
			Checksum: checksum,
		})
	}

	resp, err := c.CreateTeamImagesWithResponse(ctx, teamName, cloudquery_api.CreateTeamImagesJSONRequestBody{Images: reqs})
	if err != nil {
		return "", fmt.Errorf("failed to upload doc images: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return "", hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp)
	}

	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(4)

	fmt.Println("Uploading images...")

	for _, item := range resp.JSON201.Items {
		item := item

		key := item.Checksum + ":" + item.Name
		for i := range ims[key] {
			imItem := ims[key][i]
			imItem.url = item.URL
			ims[key][i] = imItem
		}

		if item.UploadURL == nil {
			// Already exists in bucket
			continue
		}
		fileref := ims[item.Checksum+":"+item.Name][0].absFile
		eg.Go(func() error {
			return uploadImage(egCtx, *item.UploadURL, fileref)
		})
	}

	if err := eg.Wait(); err != nil {
		return "", fmt.Errorf("failed to upload doc images: %w", err)
	}

	return replaceMarkdownImages(contents, ims), nil
}

func findMarkdownImages(contents, docDir string) (map[string][]imageReference, error) {
	re := regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	matches := re.FindAllStringSubmatch(contents, -1)

	if l := len(matches); l == 0 {
		return nil, nil
	} else if l > 100 {
		return nil, fmt.Errorf("too many images found in document. Maximum allowed is 100")
	}

	matchMap := make(map[string][]imageReference, len(matches))
	for _, match := range matches {
		fn, err := ensureValidFilename(strings.SplitN(match[2], " ", 2)[0])
		if err != nil {
			return nil, err
		}
		if fn == "" {
			continue // skip
		}

		absFile := filepath.Join(docDir, fn)

		s, err := sha1sum(absFile)
		if err != nil {
			return nil, fmt.Errorf("error processing image %q: %w", fn, err)
		}

		fileRef := filepath.Base(fn)
		matchMap[s+":"+fileRef] = append(matchMap[s+":"+fileRef], imageReference{
			mdFull:    match[0],
			mdPartial: fn,
			absFile:   absFile,
		})
	}

	return matchMap, nil
}

func replaceMarkdownImages(contents string, ims map[string][]imageReference) string {
	for _, refs := range ims {
		for _, ref := range refs {
			startPos := strings.Index(contents, ref.mdFull)
			for startPos > -1 {
				endPos := startPos + len(ref.mdFull)
				target := contents[startPos:endPos]                         // match full image tag
				target = strings.Replace(target, ref.mdPartial, ref.url, 1) // replace ref to image within full tag
				contents = contents[:startPos] + target + contents[endPos:] // put full tag back into document
				startPos = strings.Index(contents, ref.mdFull)              // try again
			}
		}
	}

	return contents
}

func uploadImage(ctx context.Context, uploadURL, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodPut, uploadURL, f)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload image: %s", resp.Status)
	}
	return nil
}

func ensureValidFilename(filename string) (string, error) {
	if strings.HasPrefix(filename, "https://") || strings.HasPrefix(filename, "http://") {
		return "", nil // skip
	}

	if strings.HasPrefix(filename, "file://") {
		u, err := url.Parse(filename)
		if err != nil {
			return "", err
		}
		if u.Host != "" && u.Host != "localhost" {
			return "", fmt.Errorf("invalid file URL %s", filename)
		}
		p := u.Path
		if strings.HasPrefix(p, "/") && os.PathSeparator == '\\' {
			p = strings.TrimPrefix(p, "/")
		}
		filename = strings.ReplaceAll(p, "/", string(os.PathSeparator))
	}

	return filename, nil
}

func sha1sum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s := sha1.New()
	_, err = io.Copy(s, f)
	if err != nil {
		return "", err
	}
	result := s.Sum(nil)
	return fmt.Sprintf("%x", result), nil
}
