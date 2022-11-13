package core

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/mocks"
	k8sTesting "github.com/cloudquery/cloudquery/plugins/source/k8s/resources/services/testing"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createCoreSecrets(t *testing.T, ctrl *gomock.Controller) client.Services {
	secrets := mocks.NewMockSecretsClient(ctrl)
	e := corev1.Secret{}

	if err := faker.FakeObject(&e); err != nil {
		t.Fatal(err)
	}
	e.ManagedFields = []metav1.ManagedFieldsEntry{k8sTesting.FakeManagedFields(t)}
	secrets.EXPECT().List(gomock.Any(), metav1.ListOptions{}).Return(
		&corev1.SecretList{Items: []corev1.Secret{e}}, nil,
	)
	return client.Services{
		Secrets: secrets,
	}
}

func TestCoreSecrets(t *testing.T) {
	client.K8sMockTestHelper(t, Secrets(), createCoreSecrets, client.TestOptions{})
}
