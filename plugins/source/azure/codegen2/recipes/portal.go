// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"

func init() {
	tables := []Table{
		{
			Service:        "armportal",
			Name:           "list_tenant_configuration_violations",
			Struct:         &armportal.Violation{},
			ResponseStruct: &armportal.ListTenantConfigurationViolationsClientListResponse{},
			Client:         &armportal.ListTenantConfigurationViolationsClient{},
			ListFunc:       (&armportal.ListTenantConfigurationViolationsClient{}).NewListPager,
			NewFunc:        armportal.NewListTenantConfigurationViolationsClient,
			URL:            "/providers/Microsoft.Portal/listTenantConfigurationViolations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Portal)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armportal",
			Name:           "tenant_configurations",
			Struct:         &armportal.Configuration{},
			ResponseStruct: &armportal.TenantConfigurationsClientListResponse{},
			Client:         &armportal.TenantConfigurationsClient{},
			ListFunc:       (&armportal.TenantConfigurationsClient{}).NewListPager,
			NewFunc:        armportal.NewTenantConfigurationsClient,
			URL:            "/providers/Microsoft.Portal/tenantConfigurations",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Portal)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
