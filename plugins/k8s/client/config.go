package client

type Config struct {
	Contexts []string `hcl:"contexts,optional"`
}

func (Config) Example() string {
	return `configuration {
  // Optional. Set contexts that you want to fetch. If it is not given then all contexts from config are iterated over.
  // contexts = ["YOUR_CONTEXT_NAME1", "YOUR_CONTEXT_NAME2"]
}`
}
