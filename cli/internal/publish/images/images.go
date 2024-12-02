package images

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"golang.org/x/exp/maps"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
)

type reference struct {
	ref         string // image filename inc. all paths (to replace with URL)
	absFile     string // absolute path to image file, to upload
	contentType string // content type of the file
	url         string // result of upload

	startPos int // start of complete markdown tag. if html: start of actual ref
	endPos   int // exclusive
}

type listKey struct {
	name, sum string
}

func DetectContentType(filename string) (string, error) {
	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType != "" {
		return contentType, nil
	}

	filebytes := make([]byte, 512)
	fp, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer fp.Close()
	if _, err := fp.Read(filebytes); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return http.DetectContentType(filebytes), nil
}

func ProcessDocument(ctx context.Context, c *cloudquery_api.ClientWithResponses, teamName, docDir, contents string) (string, error) {
	ims, err := findMarkdownImages(contents, docDir)
	if err != nil {
		return "", err
	}
	if len(ims) == 0 {
		return contents, nil
	}

	fmt.Println("Preparing to upload images...")

	reqs := make([]cloudquery_api.TeamImageCreate, 0, len(ims))
	for k, v := range ims {
		if len(v) == 0 {
			continue
		}
		absFile := v[0].absFile
		contentType, err := DetectContentType(absFile)
		if err != nil {
			return "", fmt.Errorf("failed to get content type for file %q: %w", absFile, err)
		}
		v[0].contentType = contentType
		reqs = append(reqs, cloudquery_api.TeamImageCreate{
			Name:        k.name,
			Checksum:    k.sum,
			ContentType: cloudquery_api.ContentType(contentType),
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

		key := listKey{name: item.Name, sum: item.Checksum}
		for i := range ims[key] {
			ims[key][i].url = item.URL
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
		absFile := ims[listKey{name: item.Name, sum: item.Checksum}][0].absFile
		eg.Go(func() error {
			return uploadImage(egCtx, item, absFile)
		})
	}

	if err := eg.Wait(); err != nil {
		return "", fmt.Errorf("failed to upload doc images: %w", err)
	}

	return contents, nil
}

func findMarkdownImages(contents, docDir string) (map[listKey][]reference, error) {
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

func convertMarkdownReferences(ims map[listKey][]reference) ([]reference, error) {
	type pos struct {
		at  int
		end bool
	}
	reps := make([]reference, 0, len(ims))
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

func replaceMarkdownImages(contents string, ims map[listKey][]reference) (string, error) {
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

func uploadImage(ctx context.Context, item cloudquery_api.TeamImage, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	req, err := http.NewRequest(http.MethodPut, *item.UploadURL, f)
	if err != nil {
		return err
	}
	for k, v := range item.RequiredHeaders {
		if headers, ok := v.([]any); ok {
			for _, h := range headers {
				if header, ok := h.(string); ok {
					req.Header.Add(k, header)
				}
			}
		}
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

func ensureValidFilename(filename, absDir string) (string, error) {
	u, err := url.Parse(filename)
	if err != nil {
		return "", nil // skip
	}

	if u.Scheme == "" {
		// it's a local file
		if filepath.IsAbs(filename) {
			return filename, nil
		}

		return filepath.Join(absDir, filename), nil
	} else if u.Scheme != "file" {
		return "", nil // skip
	}

	if u.Host != "" && u.Host != "localhost" {
		return "", fmt.Errorf("invalid file URL %s", filename)
	}
	p := u.Path
	if strings.HasPrefix(p, "/") && os.PathSeparator == '\\' {
		p = strings.TrimPrefix(p, "/")
	}
	filename = filepath.FromSlash(p)
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
	return hex.EncodeToString(result), nil
}

type imageFinder struct {
	images map[listKey][]reference
	docDir string
	err    error
}

func (f *imageFinder) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	if f.images == nil {
		f.images = make(map[listKey][]reference)
	}

	type refKeyType struct {
		dest, title string
	}

	refs := pc.References()
	refList := make(map[refKeyType][]parser.Reference, len(refs))
	for _, ref := range refs {
		key := refKeyType{dest: string(ref.Destination()), title: string(ref.Title())}
		refList[key] = append(refList[key], ref) // multiple refs can have the same dest/title (but different labels)
	}

	src := reader.Source()

	var allImgs []reference
	imageDestinations := make(map[refKeyType]struct{}) // referenced dests are put in here so that we can check them later
	f.err = ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		var (
			imgs []reference

			htmlBytes    []byte
			htmlStartPos int
		)
		switch el := n.(type) {
		case *ast.Image:
			imgRef := reference{
				ref: string(el.Destination),
			}
			refKey := refKeyType{dest: imgRef.ref, title: string(el.Title)}
			imageDestinations[refKey] = struct{}{} // mark this as an image so that we can cross-check with the reference list later
			if len(refList[refKey]) > 0 {
				// it's a reference, no need to check further as we won't find anything useful (regarding byte positions). Leave it to the reference handler
				return ast.WalkContinue, nil
			}

			if el.BaseNode.Parent() != nil && imgRef.endPos == 0 {
				if seg, found := handleImage(el, src); found {
					imgRef.startPos, imgRef.endPos = seg.Start, seg.Stop
				}
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
				htmlBytes = append(htmlBytes, src[a.Start:a.Stop]...)
				if i == 0 {
					htmlStartPos = a.Start
				}
			}
			// handle htmlBytes below
		case *ast.RawHTML:
			if el.Segments != nil {
				for i := 0; i < el.Segments.Len(); i++ { // should have 1 segment per tag?
					a := el.Segments.At(i)
					htmlBytes = append(htmlBytes, src[a.Start:a.Stop]...)
					if i == 0 {
						htmlStartPos = a.Start
					}
				}
			} else {
				sz := el.Lines().Len()
				for i := 0; i < sz; i++ {
					a := el.Lines().At(i)
					htmlBytes = append(htmlBytes, src[a.Start:a.Stop]...)
					if i == 0 {
						htmlStartPos = a.Start
					}
				}
			}
			// handle htmlBytes below
		default:
			return ast.WalkContinue, nil
		}

		if len(htmlBytes) > 0 {
			htmlImages, err := parseHTMLImages(htmlBytes, htmlStartPos)
			if err != nil {
				return ast.WalkStop, err
			}
			imgs = append(imgs, htmlImages...)
		}

		allImgs = append(allImgs, imgs...)
		return ast.WalkContinue, nil
	})
	if f.err != nil {
		return
	}

	f.err = func() error {
		refKeys := maps.Keys(refList)
		sort.Slice(refKeys, func(i, j int) bool {
			if refKeys[i].title == refKeys[j].title {
				return refKeys[i].dest < refKeys[j].dest
			}
			if refKeys[i].dest == refKeys[j].dest {
				return refKeys[i].title < refKeys[j].title
			}
			return refKeys[i].dest < refKeys[j].dest && refKeys[i].title < refKeys[j].title
		})

		for _, refKey := range refKeys {
			if _, isImage := imageDestinations[refKey]; !isImage {
				continue // only use this reference if it's a previously found image (ast.Image and dest/title match)
			}

			for _, pcRef := range refList[refKey] {
				// here we rebuild a `[image-id]: image.png "Optional Title Here"` line and replace it inside src. We're happy as long as there's a preceding space
				refLine := append([]byte{'['}, pcRef.Label()...)
				refLine = append(refLine, []byte("]: ")...)
				refLine = append(refLine, pcRef.Destination()...)
				if len(pcRef.Title()) > 0 {
					refLine = append(refLine, []byte(` "`)...)
					refLine = append(refLine, pcRef.Title()...)
					refLine = append(refLine, []byte{'"'}...)
				}

				if idx := bytes.Index(src, refLine); idx > 0 {
					check := bytes.TrimSpace(src[idx-1 : idx+len(refLine)]) // make sure it has some kind of space before
					if bytes.Equal(check, refLine) {
						allImgs = append(allImgs, reference{
							ref:      string(pcRef.Destination()),
							startPos: idx,
							endPos:   idx + len(refLine),
						})
					}
				}
			}
		}

		// process all matches here
		for _, img := range allImgs {
			absFile, err := ensureValidFilename(img.ref, f.docDir)
			if err != nil {
				return err
			}
			if absFile == "" {
				continue // skip
			}
			s, err := sha1sum(absFile)
			if err != nil {
				return fmt.Errorf("error processing image %q: %w", img.ref, err)
			}
			fileRef := filepath.Base(absFile)
			key := listKey{name: fileRef, sum: s}
			f.images[key] = append(f.images[key], reference{
				ref:      img.ref,
				absFile:  absFile,
				startPos: img.startPos,
				endPos:   img.endPos,
			})
		}
		return nil
	}()
}

func handleImage(el *ast.Image, source []byte) (seg text.Segment, found bool) {
	p := el.BaseNode.Parent()
	for p != nil {
		if p.Kind() == ast.KindParagraph {
			for i := 0; i < p.Lines().Len(); i++ {
				// we can have multiple lines in a paragraph, so we need to check each one to match destination/title, hoping there are no dupes
				lineLiteral := source[p.Lines().At(i).Start:p.Lines().At(i).Stop]

				// these checks are false negative if this image is a reference, but we handled them above
				if !bytes.Contains(lineLiteral, el.Destination) || (len(el.Title) > 0 && !bytes.Contains(lineLiteral, el.Title)) {
					continue
				}
				if len(el.Title) == 0 { // if no title, make sure the element ends with the link
					parts := bytes.SplitN(lineLiteral, el.Destination, 2)
					if len(parts) != 2 || !bytes.HasPrefix(bytes.TrimSpace(parts[1]), []byte(")")) { // HasPrefix because we can be inside a link
						continue
					}
				}
				return p.Lines().At(i), true
			}
			break
		}
		p = p.Parent()
	}
	return seg, false
}

func parseHTMLImages(htmlBytes []byte, htmlOffset int) ([]reference, error) {
	htmldoc, err := html.Parse(bytes.NewReader(htmlBytes))
	if err != nil {
		return nil, err
	}

	var imgs []reference
	var f func(*html.Node) int
	offset := 0

	f = func(n *html.Node) int {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key != "src" {
					continue
				}
				searchBytes := htmlBytes[offset:]
				startPos := bytes.Index(searchBytes, []byte(a.Val)) + offset
				endPos := startPos + len(a.Val)
				imgs = append(imgs, reference{
					ref:      a.Val,
					startPos: htmlOffset + startPos,
					endPos:   htmlOffset + endPos,
				})
				offset = endPos
				return offset
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			offset = f(c)
		}
		return offset
	}

	_ = f(htmldoc)
	return imgs, nil
}

var _ parser.ASTTransformer = &imageFinder{}
