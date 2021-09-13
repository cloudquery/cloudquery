module github.com/cloudquery/cq-provider-k8s

go 1.16

require (
	github.com/cloudquery/cq-provider-sdk v0.4.3
	github.com/cloudquery/faker/v3 v3.7.4
	github.com/golang/mock v1.6.0
	github.com/hashicorp/go-hclog v0.16.1
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
)

require google.golang.org/genproto v0.0.0-20210202153253-cf70463f6119 // indirect
