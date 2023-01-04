package recipes

import (
	appconnections "cloud.google.com/go/beyondcorp/appconnections/apiv1"
	appconnectionspb "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb"

	appconnectors "cloud.google.com/go/beyondcorp/appconnectors/apiv1"
	appconnectorspb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"

	appgateways "cloud.google.com/go/beyondcorp/appgateways/apiv1"
	appgatewayspb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"

	clientconnectorservices "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1"
	clientconnectorservicespb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"

	clientgateways "cloud.google.com/go/beyondcorp/clientgateways/apiv1"
	clientgatewayspb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
)

func init() {
	resources := []*Resource{
		{
			SubService:          "app_connections",
			Struct:              &appconnectionspb.AppConnection{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#AppConnection",
			NewFunction:         appconnections.NewClient,
			RegisterServer:      appconnectionspb.RegisterAppConnectionsServiceServer,
			ProtobufImport:      "cloud.google.com/go/beyondcorp/appconnections/apiv1/appconnectionspb",
			MockImports:         []string{"cloud.google.com/go/beyondcorp/appconnections/apiv1"},
			ServiceAPIOverride:  "appconnections",
		},
		{
			SubService:          "app_connectors",
			Struct:              &appconnectorspb.AppConnector{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnectors#AppConnector",
			NewFunction:         appconnectors.NewClient,
			RegisterServer:      appconnectorspb.RegisterAppConnectorsServiceServer,
			ProtobufImport:      "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb",
			MockImports:         []string{"cloud.google.com/go/beyondcorp/appconnectors/apiv1"},
			ServiceAPIOverride:  "appconnectors",
		},
		{
			SubService:          "app_gateways",
			Struct:              &appgatewayspb.AppGateway{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appGateways#AppGateway",
			NewFunction:         appgateways.NewClient,
			RegisterServer:      appgatewayspb.RegisterAppGatewaysServiceServer,
			ProtobufImport:      "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb",
			MockImports:         []string{"cloud.google.com/go/beyondcorp/appgateways/apiv1"},
			ServiceAPIOverride:  "appgateways",
		},
		{
			SubService:          "client_connector_services",
			Struct:              &clientconnectorservicespb.ClientConnectorService{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientConnectorServices#ClientConnectorService",
			NewFunction:         clientconnectorservices.NewClient,
			RegisterServer:      clientconnectorservicespb.RegisterClientConnectorServicesServiceServer,
			ProtobufImport:      "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb",
			MockImports:         []string{"cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1"},
			ServiceAPIOverride:  "clientconnectorservices",
		},
		{
			SubService:          "client_gateways",
			Struct:              &clientgatewayspb.ClientGateway{},
			PrimaryKeys:         []string{ProjectIdColumn.Name, "name"},
			RequestStructFields: `Parent: "projects/" + c.ProjectId + "/locations/-",`,
			Description:         "https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientGateways#ClientGateway",
			NewFunction:         clientgateways.NewClient,
			RegisterServer:      clientgatewayspb.RegisterClientGatewaysServiceServer,
			ProtobufImport:      "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb",
			MockImports:         []string{"cloud.google.com/go/beyondcorp/clientgateways/apiv1"},
			ServiceAPIOverride:  "clientgateways",
		},
	}

	for _, resource := range resources {
		resource.Service = "beyondcorp"
		resource.Template = "newapi_list"
		resource.MockTemplate = "newapi_list_grpc_mock"
	}

	Resources = append(Resources, resources...)
}
