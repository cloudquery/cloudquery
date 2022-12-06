package azparser

// Contains the information needed to generate CQ Table for specific Azure Client
type Table struct {
	// name of the function that creates a new azure client
	NewFuncName string
	// Rest URL parsed for specific API
	URL string
	// Namespace Parsed from the URL
	Namespace string
	// Does the client has standard NewListPager which we currently support autogeneration for
	HasListPager bool
	// param names for NewListPager
	NewListPagerParams []string
	// param names for NewXClient function
	NewClientParams []string
	Multiplex string
}
