// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx"

func Armnginx() []*Table {
	tables := []*Table{
		{
			NewFunc:        armnginx.NewDeploymentsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx",
			URL:            "/subscriptions/{subscriptionId}/providers/Nginx.NginxPlus/nginxDeployments",
			Namespace:      "Nginx.NginxPlus",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceNginx_NginxPlus)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DeploymentsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armnginx())
}
