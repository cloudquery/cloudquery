package recipes

type Resource struct {
	NewFunc interface{}
	PkgPath string
	URL string
}

var Resources [][]*Resource