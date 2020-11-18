package provider

type Interface interface {
	Run(config interface{}) error
}
