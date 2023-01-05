package azparser

// Contains the information needed to generate CQ Table for specific Azure Client
type Table struct {
	// FullImportPath
	ImportPath string
	// Package in the plugin
	PackageName    string
	BaseImportPath string
	// Name of the table
	Name string
	// name of the function that creates a new azure client
	NewFuncName string
	// Namespace Parsed from the URL
	Namespace string
	// NamespaceConst
	NamespaceConst string
	// Pager name to use
	Pager string
	// ResponseStruct
	ResponseStruct          string
	ResponspeStructNextLink bool
	ResponseValueStruct     string
	// URL is the one set by NewListAll or NewList, depending on which one is available
	URL                      string
	NewFuncHasSubscriptionId bool
	// param names for NewXClient function
	NewClientParams []string
	Multiplex       string
	Skip            bool
}
