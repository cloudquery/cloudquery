package azparser

// Contains the information needed to generate CQ Table for specific Azure Client
type Table struct {
	// name of the function that creates a new azure client
	NewFuncName string
	// Rest URL parsed for specific API
	URL string
	// Does the client has standard NewListPager which we currently support autogeneration for
	HasListPager bool
	// number of params for NewListPager
	HasListPagerParams int
}


