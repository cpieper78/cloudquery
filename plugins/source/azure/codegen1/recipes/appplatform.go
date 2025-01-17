// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform"

func Armappplatform() []*Table {
	tables := []*Table{
		{
			NewFunc:        armappplatform.NewSKUsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/skus",
			Namespace:      "Microsoft.AppPlatform",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_AppPlatform)`,
			Pager:          `NewListPager`,
			ResponseStruct: "SKUsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armappplatform())
}
