// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func Armmaintenance() []*Table {
	tables := []*Table{
		{
			NewFunc:        armmaintenance.NewConfigurationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/maintenanceConfigurations",
			Namespace:      "Microsoft.Maintenance",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Maintenance)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ConfigurationsClientListResponse",
		},
		{
			NewFunc:        armmaintenance.NewPublicMaintenanceConfigurationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
			Namespace:      "Microsoft.Maintenance",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Maintenance)`,
			Pager:          `NewListPager`,
			ResponseStruct: "PublicMaintenanceConfigurationsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armmaintenance())
}
