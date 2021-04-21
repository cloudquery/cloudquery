package mocks_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

var fakeResourceGroup = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test/providers/Microsoft.Storage/storageAccounts/cqprovidertest"

func buildComputeDiskMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockDisksClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{Disks: m},
	}
	l := compute.Disk{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	diskListPage := compute.NewDiskListPage(compute.DiskList{Value: &[]compute.Disk{l}}, func(ctx context.Context, list compute.DiskList) (compute.DiskList, error) {
		return compute.DiskList{}, nil
	},
	)
	m.EXPECT().List(gomock.Any()).Return(diskListPage, nil)
	return s
}

func buildResourceGroupMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockGroupsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{Groups: m},
	}
	l := resources.Group{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	groupPage := resources.NewGroupListResultPage(resources.GroupListResult{Value: &[]resources.Group{l}}, func(ctx context.Context, result resources.GroupListResult) (resources.GroupListResult, error) {
		return resources.GroupListResult{}, nil
	})
	m.EXPECT().List(gomock.Any(), "", nil).Return(groupPage, nil)
	return s
}

func buildKeyVaultMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	k := mocks.NewMockKeyClient(ctrl)
	v := mocks.NewMockVaultClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Vaults: v,
			Keys:   k,
		},
	}
	vault := keyvault.Vault{}
	err := faker.FakeData(&vault)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	vaultName := fakeResourceGroup
	vault.ID = &vaultName
	vaultPage := keyvault.NewVaultListResultPage(keyvault.VaultListResult{Value: &[]keyvault.Vault{vault}}, func(ctx context.Context, result keyvault.VaultListResult) (keyvault.VaultListResult, error) {
		return keyvault.VaultListResult{}, nil
	})
	v.EXPECT().ListBySubscription(gomock.Any(), gomock.Any()).Return(vaultPage, nil)

	key := keyvault.Key{}
	if err := faker.FakeData(&key); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	keyPage := keyvault.NewKeyListResultPage(keyvault.KeyListResult{Value: &[]keyvault.Key{key}}, func(ctx context.Context, result keyvault.KeyListResult) (keyvault.KeyListResult, error) {
		return keyvault.KeyListResult{}, nil
	})
	k.EXPECT().List(gomock.Any(), "test", *vault.Name).Return(keyPage, nil)
	return s
}

func buildStorageMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	acc := mocks.NewMockStorageAccountClient(ctrl)
	cont := mocks.NewMockStorageContainerClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Accounts:   acc,
			Containers: cont,
		},
	}
	account := storage.Account{}
	err := faker.FakeData(&account)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testAccount"
	account.Name = &name
	account.ID = &fakeResourceGroup
	page := storage.NewAccountListResultPage(storage.AccountListResult{Value: &[]storage.Account{account}}, func(ctx context.Context, result storage.AccountListResult) (storage.AccountListResult, error) {
		return storage.AccountListResult{}, nil
	})
	acc.EXPECT().List(gomock.Any()).Return(page, nil)

	container := storage.ListContainerItem{}
	if err := faker.FakeData(&container); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	containerPage := storage.NewListContainerItemsPage(storage.ListContainerItems{
		Value: &[]storage.ListContainerItem{container}}, func(ctx context.Context, items storage.ListContainerItems) (storage.ListContainerItems, error) {
		return storage.ListContainerItems{}, nil
	})
	cont.EXPECT().List(gomock.Any(), "test", *account.Name, "", "", gomock.Any()).Return(containerPage, nil)
	return s
}

func buildMySQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockMySQLServerClient(ctrl)
	configSvc := mocks.NewMockMySQLConfigurationClient(ctrl)
	s := services.Services{
		MySQL: services.MySQL{
			Servers:       serverSvc,
			Configuration: configSvc,
		},
	}
	server := mysql.Server{}
	err := faker.FakeData(&server)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	server.ID = &fakeResourceGroup
	page := mysql.ServerListResult{Value: &[]mysql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)

	config := mysql.Configuration{}
	if err := faker.FakeData(&config); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	configSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(mysql.ConfigurationListResult{Value: &[]mysql.Configuration{config}}, nil)
	return s
}

func buildPostgresServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockPostgresqlServerClient(ctrl)
	configSvc := mocks.NewMockPostgresqlConfigurationClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQL{
			Servers:       serverSvc,
			Configuration: configSvc,
		},
	}
	server := postgresql.Server{}
	err := faker.FakeData(&server)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	server.ID = &fakeResourceGroup
	page := postgresql.ServerListResult{Value: &[]postgresql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)

	config := postgresql.Configuration{}
	if err := faker.FakeData(&config); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	configSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(postgresql.ConfigurationListResult{Value: &[]postgresql.Configuration{config}}, nil)
	return s
}

func buildSQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockSqlServerClient(ctrl)
	databaseSvc := mocks.NewMockSqlDatabaseClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Servers:  serverSvc,
			Database: databaseSvc,
		},
	}
	server := sql.Server{}
	err := faker.FakeData(&server)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	server.ID = &fakeResourceGroup
	page := sql.ServerListResult{Value: &[]sql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)
	database := sql.Database{}
	if err := faker.FakeData(&database); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	databaseSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name, "true", "").Return(sql.DatabaseListResult{Value: &[]sql.Database{database}}, nil)
	return s
}
