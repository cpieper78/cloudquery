// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

func Armcontainerregistry() []*Table {
	tables := []*Table{
		{
			NewFunc:        armcontainerregistry.NewRegistriesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries",
			Namespace:      "Microsoft.ContainerRegistry",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ContainerRegistry)`,
			Pager:          `NewListPager`,
			ResponseStruct: "RegistriesClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armcontainerregistry())
}
