package publish

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMarkdownImages(t *testing.T) {
	cases := []struct {
		name        string
		contents    string
		expect      map[string][]imageReference
		expectError string
		skipCheck   bool
	}{
		{
			name: "no images",
			contents: `# Title
`,
			expect: nil,
		},
		{
			name: "basic",
			contents: `# Title
![](image.png)`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 22}}},
		},
		{
			name: "basic with alt",
			contents: `# Title
![Alt text](image.png)`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 30}}},
		},
		{
			name: "html with alt and width",
			contents: `# Title
<img src="image.png" alt="Alt text" width="100%">`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 57}}},
		},
		{
			name: "double html with alt and width",
			contents: `# Title
<img src="image.png" alt="Alt text" width="100%"> <img src="image.png" alt="Alt text2" width="50%">`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 57}, {ref: "image.png", startPos: 58, endPos: 107}}},
		},
		{
			name: "tricky html",
			contents: `# Title
<img src="image.png"
alt="Alt text" width="100%">`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 57}}},
		},
		{
			name: "quoted html",
			contents: `# Title
` + "```" + `
<img src="image.png" alt="Alt text" width="100%">
` + "```" + `
`,
			expect: nil,
		},
		{
			name: "quoted image",
			contents: `# Title
` + "```" + `
![Alt text](image.png "Title Here")
` + "```" + `
`,
			expect: nil,
		},
		{
			name: "single quoted html",
			contents: `# Title
` + "`" + `<img src="image.png" alt="Alt text" width="100%">` + "`" + `
`,
			expect: nil,
		},
		{
			name: "single quoted image",
			contents: `# Title
` + "`" + `![Alt text](image.png "Title Here")` + "`" + `
`,
			expect: nil,
		},
		{
			name: "basic with alt and title",
			contents: `# Title
![Alt text](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 43}}},
		},
		{
			name: "basic with title",
			contents: `# Title
![](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 35}}},
		},
		{
			name: "basic with title or alt, multiple",
			contents: `# Title
![](image.png "Title Here")
More test
![alt](image.png)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 36}, {ref: "image.png", startPos: 46, endPos: 63}}},
		},
		{
			name: "ref",
			contents: `# Title
![Alt text][image-id]

[image-id]: image.png "Optional Title Here"
`,
			expect:    map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png"}}}, // issue: startpos/endpos don't point to file
			skipCheck: true,
		},
		{
			name: "ref multiple",
			contents: `# Title
![Alt text][image-id]

text

![Same text][image-id]

[image-id]: image.png "Optional Title Here"
`,
			expect:    map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png"}, {ref: "image.png"}}}, // two references to same image. startpos/endpos issue as above
			skipCheck: true,
		},
		{
			name: "href",
			contents: `# Title
[![Alt text](image.png)](http://example.com/)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "image.png", startPos: 8, endPos: 53}}}, // includes full href
		},
		{
			name: "subdir",
			contents: `# Title
![Alt text](assets/images/image.png)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{ref: "assets/images/image.png", startPos: 8, endPos: 44}}},
		},
		{
			name: "basic file://",
			contents: `# Title
![](file://${ABS_IMAGE})`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:my special@image.png": {{ref: "file://${ABS_IMAGE}", startPos: 8, endPos: -12}}}, // 12 is the number of extra characters except the placeholder/filename
		},
	}

	tempdir := t.TempDir()
	data := []byte("test")
	require.NoError(t, os.WriteFile(filepath.Join(tempdir, "image.png"), data, 0644))
	require.NoError(t, os.MkdirAll(filepath.Join(tempdir, "assets", "images"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(tempdir, "assets", "images", "image.png"), data, 0644))

	special := filepath.Join(tempdir, "my special@image.png")
	require.NoError(t, os.WriteFile(special, data, 0644))
	specialAbs, err := filepath.Abs(special)
	require.NoError(t, err)
	specialAbsEscaped := strings.NewReplacer(" ", "%20", "@", "%40", string(filepath.Separator), "/").Replace(specialAbs)
	if filepath.Separator == '\\' {
		specialAbsEscaped = "/" + specialAbsEscaped // absolute paths in file:// mode start with "/", not C:\ (`file:///C:/`, not `file://C:/`)
	}

	const mockURL = "https://example.com/file.ext"
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.contents = strings.ReplaceAll(tc.contents, "${ABS_IMAGE}", specialAbsEscaped)
			out, err := findMarkdownImages(tc.contents, tempdir)
			if tc.expectError != "" {
				require.ErrorContains(t, err, tc.expectError)
				return
			}
			require.NoError(t, err)

			// clean up for successful comparison, also set mock URL
			for k, v := range out {
				for i := range v {
					//v[i].ref = strings.ReplaceAll(v[i].ref, specialAbsEscaped, "${ABS_IMAGE}")
					v[i].absFile = ""
					v[i].url = mockURL
				}
				out[k] = v
			}
			if len(out) == 0 {
				out = nil
			}
			// adjust end position for special case
			for k, v := range tc.expect {
				for i := range v {
					if strings.Contains(v[i].ref, "${ABS_IMAGE}") {
						v[i].ref = strings.ReplaceAll(v[i].ref, "${ABS_IMAGE}", specialAbsEscaped)
						if v[i].endPos < 0 {
							v[i].endPos = v[i].startPos + len(specialAbsEscaped) - v[i].endPos
						}
					}
					v[i].url = mockURL
				}
				tc.expect[k] = v
			}

			require.EqualValues(t, tc.expect, out)

			replaced, err := replaceMarkdownImages(tc.contents, out)
			require.NoError(t, err)

			if len(tc.expect) > 0 && !tc.skipCheck {
				require.Contains(t, replaced, mockURL)
			}
		})
	}
}
