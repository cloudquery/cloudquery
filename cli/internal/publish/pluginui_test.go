package publish

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickContentType(t *testing.T) {
	tests := []struct {
		filename string
		data     []byte
		want     string
	}{
		{
			filename: "octetstream.svg",
			data:     []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="2500" height="330" viewBox="0 0 719.713 94.943"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M17.32,16.76l-.53.53-1.09-1.1a4.33,4.33,0,0,0,.53-.52Z" style="fill:#4285f4"/></g></svg>`),
			want:     "image/svg+xml",
		},
		{
			filename: "textplain.svg",
			data:     []byte(`<svg id="Artwork" xmlns="http://www.w3.org/2000/svg" width="24px" height="24px" viewBox="0 0 24 24"><path d="M8.36,7.8a4.4,4.4,0,0,0-.54.52l-.23-.23-.88-.88.53-.53.88.88Z" style="fill:#4285f4"/><path d="M17.32,16.76l-.53.53-1.09-1.1a4.33,4.33,0,0,0,.53-.52Z" style="fill:#4285f4"/><path d="M6.58,11.6a2.61,2.61,0,0,0,0,.4,2,2,0,0,0,0,.35H4.87V11.6Z" style="fill:#4285f4"/><path d="M19.13,11.6v.75H17.47c0-.11,0-.23,0-.35s0-.27,0-.4Z" style="fill:#4285f4"/><path d="M7.79,15.65a4.42,4.42,0,0,0,.53.53L7.21,17.29l-.53-.53Z" style="fill:#4285f4"/><path d="M15.9,7.54l.86-.86.53.53L16.2,8.3a4.42,4.42,0,0,0-.53-.53Z" style="fill:#4285f4"/><path d="M12,17.47l.38,0v1.67h-.75V17.45Z" style="fill:#4285f4"/><path d="M12.4,4.87V6.56a1.94,1.94,0,0,0-.34,0h0a2.68,2.68,0,0,0-.4,0V4.87Z" style="fill:#4285f4"/><path d="M17.92,11.6a5.88,5.88,0,0,0-1.49-3.53,5.29,5.29,0,0,0-.54-.52A5.81,5.81,0,0,0,12.4,6.1H12l-.37,0A5.84,5.84,0,0,0,8.12,7.56a6.62,6.62,0,0,0-.53.53A5.89,5.89,0,0,0,6.12,11.6c0,.13,0,.27,0,.4s0,.24,0,.35a5.89,5.89,0,0,0,1.44,3.53,5.3,5.3,0,0,0,.53.53,5.91,5.91,0,0,0,3.56,1.49h.75a5.89,5.89,0,0,0,3.52-1.47,4.42,4.42,0,0,0,.53-.53,5.83,5.83,0,0,0,1.47-3.55c0-.11,0-.23,0-.35S17.93,11.73,17.92,11.6ZM12,17.16A5.16,5.16,0,1,1,17.16,12,5.16,5.16,0,0,1,12,17.16Zm0-9.09A3.93,3.93,0,1,0,16,12,3.94,3.94,0,0,0,12,8.07Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M12,5a1.4,1.4,0,0,1,0-2.79V3a.65.65,0,0,0,0,1.29Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M12,5V4.29A.65.65,0,0,0,12,3V2.25A1.4,1.4,0,0,1,12,5Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M12,21.75A1.4,1.4,0,0,1,12,19v.75A.65.65,0,0,0,12,21Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M12,21.75V21a.65.65,0,0,0,0-1.29V19a1.4,1.4,0,0,1,0,2.79Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M6.09,7.48a1.39,1.39,0,1,1,0-2.78v.75a.64.64,0,0,0-.64.64.64.64,0,0,0,.64.64Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M6.09,7.48V6.73a.64.64,0,0,0,0-1.28V4.7a1.39,1.39,0,1,1,0,2.78Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M17.91,19.3a1.39,1.39,0,0,1,0-2.78v.75a.64.64,0,0,0,0,1.28Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M17.91,19.3v-.75a.64.64,0,0,0,.64-.64.64.64,0,0,0-.64-.64v-.75a1.39,1.39,0,0,1,0,2.78Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M3.64,13.39a1.39,1.39,0,0,1,0-2.78v.75a.64.64,0,1,0,0,1.28Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M3.64,13.39v-.75a.64.64,0,1,0,0-1.28v-.75a1.39,1.39,0,1,1,0,2.78Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M20.36,13.39a1.38,1.38,0,0,1-1-.4A1.43,1.43,0,0,1,19,12a1.4,1.4,0,0,1,1.4-1.39v.75a.65.65,0,0,0-.65.64.64.64,0,0,0,.19.45.6.6,0,0,0,.46.19Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M20.36,13.39v-.75a.64.64,0,0,0,0-1.28v-.75a1.39,1.39,0,0,1,0,2.78Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M6.06,19.3a1.39,1.39,0,0,1-1-.42,1.33,1.33,0,0,1-.38-1,1.44,1.44,0,0,1,1.45-1.38h0v.75h0a.69.69,0,0,0-.7.65.57.57,0,0,0,.17.44.67.67,0,0,0,.47.19Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M6.06,19.3h0v-.75h0a.69.69,0,0,0,.7-.65.56.56,0,0,0-.17-.44.67.67,0,0,0-.47-.19v-.75a1.39,1.39,0,0,1,1,.42,1.32,1.32,0,0,1,.38,1A1.44,1.44,0,0,1,6.06,19.3Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M17.91,7.48a1.39,1.39,0,0,1,0-2.78v.75a.64.64,0,1,0,0,1.28Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M17.91,7.48V6.73a.64.64,0,0,0,.64-.64.64.64,0,0,0-.64-.64V4.7a1.39,1.39,0,1,1,0,2.78Z" style="fill:#669df6;fill-rule:evenodd"/><path d="M12,8.07v7.85a3.93,3.93,0,1,1,0-7.85Z" style="fill:#aecbfa;fill-rule:evenodd"/><path d="M16,12A3.92,3.92,0,0,1,12,15.92V8.07A3.93,3.93,0,0,1,16,12Z" style="fill:#669df6;fill-rule:evenodd"/><rect x="12.02" y="13.04" width="2.11" height="0.75" style="fill:#aecbfa"/><rect x="12.02" y="11.67" width="2.11" height="0.75" style="fill:#aecbfa"/><rect x="12.02" y="10.31" width="2.11" height="0.75" style="fill:#aecbfa"/><rect x="9.91" y="13.04" width="2.11" height="0.75" style="fill:#fff"/><rect x="9.91" y="11.67" width="2.11" height="0.75" style="fill:#fff"/><rect x="9.91" y="10.31" width="2.11" height="0.75" style="fill:#fff"/></svg>`),
			want:     "image/svg+xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			fullpath, err := writeTempFile(t, tt.filename, tt.data)
			require.NoError(t, err)

			qt, err := quickContentType(fullpath)
			require.NoError(t, err)
			require.Equal(t, tt.want, qt)
		})
	}
}

func writeTempFile(t *testing.T, filename string, data []byte) (string, error) {
	t.Helper()

	dir := t.TempDir()
	fullpath := filepath.Join(dir, filename)

	fp, err := os.Create(fullpath)
	require.NoError(t, err)
	defer fp.Close()

	_, err = fp.Write(data)
	require.NoError(t, err)
	return fullpath, nil
}
