package publish

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/internal/hub"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"golang.org/x/sync/errgroup"
)

type imageReference struct {
	ref     string // image filename inc. all paths (to replace with URL)
	absFile string // absolute path to image file, to upload
	url     string // result of upload
	html    []byte // html section, superset of ref (only for html)

	startPos, endPos int // start,end of complete markdown tag. if html: start,end of actual ref
}

var htmlImageRe = regexp.MustCompile(`<img\s+(?:[^>"']|"[^"]*"|'[^']*')*\s*src="([^"]*)"`)

func processDocumentImages(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, docDir, contents string) (string, error) {
	ims, err := findMarkdownImages(contents, docDir)
	if err != nil {
		return "", err
	}
	if len(ims) == 0 {
		return contents, nil
	}

	fmt.Println("Preparing to upload images...")

	reqs := make([]cloudquery_api.TeamImageCreate, 0, len(ims))
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
		return "", fmt.Errorf("failed to prepare doc images: %w", err)
	}
	if resp.HTTPResponse.StatusCode > 299 {
		return "", fmt.Errorf("failed preparing: %w", hub.ErrorFromHTTPResponse(resp.HTTPResponse, resp))
	}

	for _, item := range resp.JSON201.Items {
		item := item

		key := item.Checksum + ":" + item.Name
		for i := range ims[key] {
			imItem := ims[key][i]
			imItem.url = item.URL
			ims[key][i] = imItem
		}
	}

	contents, err = replaceMarkdownImages(contents, ims)
	if err != nil {
		return "", fmt.Errorf("failed replacing markdown: %w", err)
	}

	fmt.Println("Uploading images...")
	eg, egCtx := errgroup.WithContext(ctx)
	eg.SetLimit(4)

	// Iterate again to upload
	for _, item := range resp.JSON201.Items {
		if item.UploadURL == nil {
			// Already exists in bucket
			continue
		}
		item := item
		fileref := ims[item.Checksum+":"+item.Name][0].absFile
		eg.Go(func() error {
			return uploadImage(egCtx, *item.UploadURL, fileref)
		})
	}

	if err := eg.Wait(); err != nil {
		return "", fmt.Errorf("failed to upload doc images: %w", err)
	}

	return contents, nil
}

func findMarkdownImages(contents, docDir string) (map[string][]imageReference, error) {
	imf := &imageFinder{
		docDir: docDir,
	}
	p := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(util.Prioritized(imf, 999999)),
		),
	)
	if err := p.Convert([]byte(contents), io.Discard); err != nil {
		return nil, fmt.Errorf("failed to parse markdown: %w", err)
	}
	if imf.err != nil {
		return nil, imf.err
	}

	return imf.images, nil
}

func convertMarkdownReferences(ims map[string][]imageReference) ([]imageReference, error) {
	type pos struct {
		at  int
		end bool
	}
	reps := make([]imageReference, 0, len(ims))
	posList := make([]pos, 0, 2*len(ims))
	for _, imList := range ims {
		for _, ir := range imList {
			if ir.endPos == 0 {
				// return "", fmt.Errorf("unknown range for image %q", ir.ref)
				continue // skip
			}
			if ir.startPos >= ir.endPos {
				return nil, fmt.Errorf("invalid range for image %q", ir.ref)
			}

			posList = append(posList, pos{at: ir.startPos}, pos{at: ir.endPos - 1, end: true})
			ir := ir
			reps = append(reps, ir)
		}
	}

	sort.Slice(posList, func(i, j int) bool {
		if posList[i].at == posList[j].at {
			return !posList[i].end
		}
		return posList[i].at < posList[j].at
	})

	// check for overlaps
	var (
		lastPos int
		open    int
	)
	for _, p := range posList {
		if p.end {
			open--
		} else {
			open++
		}
		if open > 1 {
			return nil, fmt.Errorf("found overlapping range: %d-%d", lastPos, p.at)
		}
		lastPos = p.at
	}

	return reps, nil
}

func replaceMarkdownImages(contents string, ims map[string][]imageReference) (string, error) {
	reps, err := convertMarkdownReferences(ims)
	if err != nil {
		return "", err
	}

	// sort by start position, descending
	sort.Slice(reps, func(i, j int) bool {
		return reps[i].startPos > reps[j].startPos
	})

	// replace in reverse order
	for _, ir := range reps {
		literalTag := contents[ir.startPos:ir.endPos]
		replacedTag := strings.Replace(literalTag, ir.ref, ir.url, 1)
		contents = contents[:ir.startPos] + replacedTag + contents[ir.endPos:]
	}

	return contents, nil
}

func uploadImage(ctx context.Context, uploadURL, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	first512 := make([]byte, 512)
	if _, err := f.Read(first512); err != nil {
		return err
	}
	ct := http.DetectContentType(first512)
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, uploadURL, f)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", ct)

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

func ensureValidFilename(filename, absDir string) (string, error) {
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
		return filename, nil
	}

	if filepath.IsAbs(filename) {
		return filename, nil
	}

	return filepath.Join(absDir, filename), nil
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

type imageFinder struct {
	images map[string][]imageReference
	docDir string
	err    error
}

func (f *imageFinder) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	if f.images == nil {
		f.images = make(map[string][]imageReference)
	}
	if f.err != nil {
		return
	}
	src := reader.Source()
	f.err = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		var (
			imgs []imageReference

			html         []byte
			htmlStartPos int
		)
		switch el := n.(type) {
		case *ast.Image:
			imgRef := imageReference{
				ref: string(el.Destination),
			}
			p := el.BaseNode.Parent()
			for p != nil {
				if p.Kind() == ast.KindParagraph {
					for i := 0; i < p.Lines().Len(); i++ {
						// we can have multiple lines in a paragraph, so we need to check each one to match destination/title, hoping there are no dupes
						lineLiteral := src[p.Lines().At(i).Start:p.Lines().At(i).Stop]
						if !bytes.Contains(lineLiteral, el.Destination) {
							continue
						}
						if len(el.Title) > 0 && !bytes.Contains(lineLiteral, el.Title) {
							continue
						}
						if len(el.Title) == 0 { // if no title, make sure the element ends with the link
							parts := bytes.SplitN(lineLiteral, el.Destination, 2)
							if len(parts) != 2 || !bytes.HasPrefix(bytes.TrimSpace(parts[1]), []byte(")")) { // HasPrefix because we can be inside a link
								continue
							}
						}
						imgRef.startPos, imgRef.endPos = p.Lines().At(i).Start, p.Lines().At(i).Stop
						break
					}
					break
				}
				p = p.Parent()
			}
			imgs = append(imgs, imgRef)
		case *ast.CodeBlock:
			return ast.WalkSkipChildren, nil
		case *ast.FencedCodeBlock:
			return ast.WalkSkipChildren, nil
		case *ast.HTMLBlock:
			sz := el.Lines().Len()
			for i := 0; i < sz; i++ {
				a := el.Lines().At(i)
				html = append(html, src[a.Start:a.Stop]...)
				if i == 0 {
					htmlStartPos = a.Start
				}
			}
		case *ast.RawHTML:
			if el.Segments != nil {
				for i := 0; i < el.Segments.Len(); i++ { // should have 1 segment per tag?
					a := el.Segments.At(i)
					html = append(html, src[a.Start:a.Stop]...)
					if i == 0 {
						htmlStartPos = a.Start
					}
				}
			} else {
				sz := el.Lines().Len()
				for i := 0; i < sz; i++ {
					a := el.Lines().At(i)
					html = append(html, src[a.Start:a.Stop]...)
					if i == 0 {
						htmlStartPos = a.Start
					}
				}
			}
		default:
			return ast.WalkContinue, nil
		}

		if len(html) > 0 {
			matchPositions := htmlImageRe.FindAllStringSubmatchIndex(string(html), -1)
			if len(matchPositions) == 0 {
				return ast.WalkContinue, nil // not an image
			}
			for _, m := range matchPositions {
				imgs = append(imgs, imageReference{
					ref:      string(html[m[2]:m[3]]),
					startPos: htmlStartPos + m[2],
					endPos:   htmlStartPos + m[3],
				})
			}
		}

		if len(imgs) == 0 {
			return ast.WalkContinue, nil
		}

		for _, img := range imgs {
			absFile, err := ensureValidFilename(img.ref, f.docDir)
			if err != nil {
				return ast.WalkStop, err
			}
			if absFile == "" {
				return ast.WalkContinue, nil // skip
			}
			s, err := sha1sum(absFile)
			if err != nil {
				return ast.WalkStop, fmt.Errorf("error processing image %q: %w", img.ref, err)
			}
			fileRef := filepath.Base(absFile)
			f.images[s+":"+fileRef] = append(f.images[s+":"+fileRef], imageReference{
				ref:      img.ref,
				absFile:  absFile,
				startPos: img.startPos,
				endPos:   img.endPos,
			})
		}

		return ast.WalkContinue, nil
	})
}

var _ parser.ASTTransformer = &imageFinder{}
