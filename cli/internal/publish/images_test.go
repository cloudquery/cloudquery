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
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with alt",
			contents: `# Title
![Alt text](image.png)`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with alt and title",
			contents: `# Title
![Alt text](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png \"Title Here\")", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with title",
			contents: `# Title
![](image.png "Title Here")
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png \"Title Here\")", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "basic with title or alt, multiple",
			contents: `# Title
![](image.png "Title Here")
More test
![alt](image.png)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![](image.png \"Title Here\")", mdPartial: "image.png", url: ""}, {mdFull: "![alt](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "ref",
			contents: `# Title
![Alt text][image-id]

[image-id]: image.png "Optional Title Here"
`,
			expect: nil, // unsupported
		},
		{
			name: "href",
			contents: `# Title
[![Alt text](image.png)](http://example.com/)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](image.png)", mdPartial: "image.png", url: ""}}},
		},
		{
			name: "subdir",
			contents: `# Title
![Alt text](assets/images/image.png)
`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:image.png": {{mdFull: "![Alt text](assets/images/image.png)", mdPartial: "assets/images/image.png", url: ""}}},
		},
		{
			name: "basic file://",
			contents: `# Title
![](file://${ABS_IMAGE})`,
			expect: map[string][]imageReference{"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3:my special@image.png": {{mdFull: "![](file://${ABS_IMAGE})", mdPartial: "file://${ABS_IMAGE}", url: ""}}},
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
		specialAbsEscaped = strings.Replace(specialAbsEscaped, "file://", "file:///", 1)
	}

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

			// reset absFile to empty string for comparison, also put back ${ABS_IMAGE} placeholder
			for k, v := range out {
				for i := range v {
					v[i].mdFull = strings.ReplaceAll(v[i].mdFull, specialAbsEscaped, "${ABS_IMAGE}")
					v[i].mdPartial = strings.ReplaceAll(v[i].mdPartial, specialAbsEscaped, "${ABS_IMAGE}")
					v[i].absFile = ""
				}
				out[k] = v
			}

			require.EqualValues(t, tc.expect, out)
		})
	}
}
