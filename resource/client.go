package resource

type ClientInterface interface {
	CollectResource(resource string, config interface{}) error
}
