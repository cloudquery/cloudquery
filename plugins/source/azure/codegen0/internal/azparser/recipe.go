package azparser

// Contains the information needed to generate CQ Table for specific Azure Client
type Table struct {
	// name of the function that creates a new azure client
	NewFuncName string
	// Namespace Parsed from the URL
	Namespace string
	// Pager name to use
	Pager string
	// ResponseStruct
	ResponseStruct string
	// URL is the one set by NewListAll or NewList, depending on which one is available
	URL string
	// param names for NewXClient function
	NewClientParams []string
	Multiplex       string
}
