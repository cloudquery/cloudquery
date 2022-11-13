mockgen --package=mocks --destination="client/mocks/mock.go" "k8s.io/client-go/kubernetes" Interface
mockgen --package=mocks --destination="client/mocks/batchv1.go" "k8s.io/client-go/kubernetes/typed/batch/v1" BatchV1Interface
mockgen --package=mocks --destination="client/mocks/batchv1_cronjob.go" "k8s.io/client-go/kubernetes/typed/batch/v1" CronJobInterface
mockgen --package=mocks --destination="client/mocks/appsv1.go" "k8s.io/client-go/kubernetes/typed/apps/v1" AppsV1Interface
mockgen --package=mocks --destination="client/mocks/appsv1_daemonset.go" "k8s.io/client-go/kubernetes/typed/apps/v1" DaemonSetInterface
# ~/go/bin/mockgen --package=client --destination="client/mock.go" "k8s.io/client-go/kubernetes" Interface