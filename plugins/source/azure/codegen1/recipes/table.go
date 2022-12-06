package recipes

type Table struct {
	NewFunc interface{}
	PkgPath string
	URL     string
	Namespace string
	Multiplex string
}

var Tables [][]*Table
