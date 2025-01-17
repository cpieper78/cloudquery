// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"

func Armdnsresolver() []*Table {
	tables := []*Table{
		{
			NewFunc:        armdnsresolver.NewDNSForwardingRulesetsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsForwardingRulesets",
			Namespace:      "Microsoft.Network",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DNSForwardingRulesetsClientListResponse",
		},
		{
			NewFunc:        armdnsresolver.NewDNSResolversClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnsResolvers",
			Namespace:      "Microsoft.Network",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Network)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DNSResolversClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armdnsresolver())
}
