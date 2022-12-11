package recipes

type Table struct {
	NewFunc        interface{}
	PkgPath        string
	URL            string
	Namespace      string
	Multiplex      string
	Pager          string
	ResponseStruct string
}

var Tables [][]*Table
