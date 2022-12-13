// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"

func Armhanaonazure() []*Table {
	tables := []*Table{
		{
			NewFunc:        armhanaonazure.NewSapMonitorsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.HanaOnAzure/sapMonitors",
			Namespace:      "Microsoft.HanaOnAzure",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_HanaOnAzure)`,
			Pager:          `NewListPager`,
			ResponseStruct: "SapMonitorsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armhanaonazure())
}