package recipes

type Resource struct {
	NewFunc interface{}
	PkgPath string
}

var Resources [][]*Resource