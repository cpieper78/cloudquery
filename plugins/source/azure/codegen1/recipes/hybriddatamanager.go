// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager"

func Armhybriddatamanager() []*Table {
	tables := []*Table{
		{
			NewFunc:        armhybriddatamanager.NewDataManagersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybriddatamanager/armhybriddatamanager",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HybridData/dataManagers",
			Namespace:      "Microsoft.HybridData",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_HybridData)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DataManagersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhybriddatamanager())
}
