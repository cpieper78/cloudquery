// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"

func Armportal() []*Table {
	tables := []*Table{
		{
			NewFunc:        armportal.NewListTenantConfigurationViolationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal",
			URL:            "/providers/Microsoft.Portal/listTenantConfigurationViolations",
			Namespace:      "Microsoft.Portal",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Portal)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ListTenantConfigurationViolationsClientListResponse",
		},
		{
			NewFunc:        armportal.NewTenantConfigurationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal",
			URL:            "/providers/Microsoft.Portal/tenantConfigurations",
			Namespace:      "Microsoft.Portal",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Portal)`,
			Pager:          `NewListPager`,
			ResponseStruct: "TenantConfigurationsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armportal())
}
