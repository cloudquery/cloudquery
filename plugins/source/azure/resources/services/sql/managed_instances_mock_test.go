package sql

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSQLManagedInstancesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	instancesClient := mocks.NewMockManagedInstancesClient(ctrl)
	databasesClient := mocks.NewMockManagedDatabasesClient(ctrl)
	instanceVulnerabilityAssessmentsClient := mocks.NewMockManagedInstanceVulnerabilityAssessmentsClient(ctrl)
	instanceEncryptionProtectorsClient := mocks.NewMockManagedInstanceEncryptionProtectorsClient(ctrl)
	managedDatabaseVulnerabilityAssessmentsClient := mocks.NewMockManagedDatabaseVulnerabilityAssessmentsClient(ctrl)
	managedDatabaseVulnerabilityAssessmentsScansClient := mocks.NewMockManagedDatabaseVulnerabilityAssessmentScansClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			ManagedDatabases:                            databasesClient,
			ManagedDatabaseVulnerabilityAssessments:     managedDatabaseVulnerabilityAssessmentsClient,
			ManagedInstances:                            instancesClient,
			ManagedInstanceVulnerabilityAssessments:     instanceVulnerabilityAssessmentsClient,
			ManagedDatabaseVulnerabilityAssessmentScans: managedDatabaseVulnerabilityAssessmentsScansClient,
			ManagedInstanceEncryptionProtectors:         instanceEncryptionProtectorsClient,
		},
	}
	instance := sql.ManagedInstance{}
	if err := faker.FakeData(&instance); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	instance.Name = &name
	rg := client.FakeResourceGroup
	instance.ID = &rg
	instancesClient.EXPECT().List(gomock.Any()).Return(
		sql.NewManagedInstanceListResultPage(
			sql.ManagedInstanceListResult{Value: &[]sql.ManagedInstance{instance}},
			func(context.Context, sql.ManagedInstanceListResult) (sql.ManagedInstanceListResult, error) {
				return sql.ManagedInstanceListResult{}, nil
			},
		),
		nil,
	)

	database := sql.ManagedDatabase{}
	if err := faker.FakeData(&database); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	database.ID = &rg
	databasesClient.EXPECT().ListByInstance(gomock.Any(), "test", *instance.Name).Return(
		sql.NewManagedDatabaseListResultPage(
			sql.ManagedDatabaseListResult{Value: &[]sql.ManagedDatabase{database}},
			func(context.Context, sql.ManagedDatabaseListResult) (sql.ManagedDatabaseListResult, error) {
				return sql.ManagedDatabaseListResult{}, nil
			},
		), nil,
	)

	var instanceVuln sql.ManagedInstanceVulnerabilityAssessment
	if err := faker.FakeData(&instanceVuln); err != nil {
		t.Fatal(err)
	}
	instanceVulnerabilityAssessmentsClient.EXPECT().ListByInstance(gomock.Any(), "test", *instance.Name).Return(
		sql.NewManagedInstanceVulnerabilityAssessmentListResultPage(
			sql.ManagedInstanceVulnerabilityAssessmentListResult{Value: &[]sql.ManagedInstanceVulnerabilityAssessment{instanceVuln}},
			func(context.Context, sql.ManagedInstanceVulnerabilityAssessmentListResult) (sql.ManagedInstanceVulnerabilityAssessmentListResult, error) {
				return sql.ManagedInstanceVulnerabilityAssessmentListResult{}, nil
			},
		), nil,
	)

	var protector sql.ManagedInstanceEncryptionProtector
	if err := faker.FakeData(&protector); err != nil {
		t.Fatal(err)
	}
	instanceEncryptionProtectorsClient.EXPECT().ListByInstance(gomock.Any(), "test", *instance.Name).Return(
		sql.NewManagedInstanceEncryptionProtectorListResultPage(
			sql.ManagedInstanceEncryptionProtectorListResult{Value: &[]sql.ManagedInstanceEncryptionProtector{protector}},
			func(context.Context, sql.ManagedInstanceEncryptionProtectorListResult) (sql.ManagedInstanceEncryptionProtectorListResult, error) {
				return sql.ManagedInstanceEncryptionProtectorListResult{}, nil
			},
		), nil,
	)

	var dbVuln sql.DatabaseVulnerabilityAssessment
	if err := faker.FakeData(&dbVuln); err != nil {
		t.Fatal(err)
	}
	managedDatabaseVulnerabilityAssessmentsClient.EXPECT().ListByDatabase(gomock.Any(), "test", *instance.Name, *database.Name).Return(
		sql.NewDatabaseVulnerabilityAssessmentListResultPage(
			sql.DatabaseVulnerabilityAssessmentListResult{Value: &[]sql.DatabaseVulnerabilityAssessment{dbVuln}},
			func(context.Context, sql.DatabaseVulnerabilityAssessmentListResult) (sql.DatabaseVulnerabilityAssessmentListResult, error) {
				return sql.DatabaseVulnerabilityAssessmentListResult{}, nil
			},
		), nil,
	)

	var record sql.VulnerabilityAssessmentScanRecord
	if err := faker.FakeData(&record); err != nil {
		t.Fatal(err)
	}
	managedDatabaseVulnerabilityAssessmentsScansClient.EXPECT().ListByDatabase(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		sql.NewVulnerabilityAssessmentScanRecordListResultPage(
			sql.VulnerabilityAssessmentScanRecordListResult{Value: &[]sql.VulnerabilityAssessmentScanRecord{record}},
			func(context.Context, sql.VulnerabilityAssessmentScanRecordListResult) (sql.VulnerabilityAssessmentScanRecordListResult, error) {
				return sql.VulnerabilityAssessmentScanRecordListResult{}, nil
			},
		), nil,
	)

	return s
}

func TestSQLManagedInstances(t *testing.T) {
	client.AzureMockTestHelper(t, SqlManagedInstances(), buildSQLManagedInstancesMock, client.TestOptions{})
}
