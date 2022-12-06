package recipes

type Table struct {
	NewFunc interface{}
	PkgPath string
	URL string
}

var Tables [][]*Table