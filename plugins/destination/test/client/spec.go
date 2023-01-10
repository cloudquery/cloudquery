package client

type Spec struct {
	// ErrorOnWrite: if true, will return an error on write rather than consume from the channel
	ErrorOnWrite bool `json:"error_on_write,omitempty"`
}
