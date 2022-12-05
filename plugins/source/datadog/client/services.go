// Code generated by codegen; DO NOT EDIT.
package client

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client/services"
)

type DatadogServices struct {
	DashboardListsAPI services.DashboardListsAPIClient
	DashboardsAPI     services.DashboardsAPIClient
	DowntimesAPI      services.DowntimesAPIClient
	HostsAPI          services.HostsAPIClient
	IncidentsAPI      services.IncidentsAPIClient
	MonitorsAPI       services.MonitorsAPIClient
	NotebooksAPI      services.NotebooksAPIClient
	RolesAPI          services.RolesAPIClient
	SyntheticsAPI     services.SyntheticsAPIClient
	UsersAPI          services.UsersAPIClient
}

func NewDatadogServices(apiClient *datadog.APIClient) DatadogServices {
	return DatadogServices{
		DashboardListsAPI: datadogV1.NewDashboardListsApi(apiClient),
		DashboardsAPI:     datadogV1.NewDashboardsApi(apiClient),
		DowntimesAPI:      datadogV1.NewDowntimesApi(apiClient),
		HostsAPI:          datadogV1.NewHostsApi(apiClient),
		IncidentsAPI:      datadogV2.NewIncidentsApi(apiClient),
		MonitorsAPI:       datadogV1.NewMonitorsApi(apiClient),
		NotebooksAPI:      datadogV1.NewNotebooksApi(apiClient),
		RolesAPI:          datadogV2.NewRolesApi(apiClient),
		SyntheticsAPI:     datadogV1.NewSyntheticsApi(apiClient),
		UsersAPI:          datadogV2.NewUsersApi(apiClient),
	}
}
