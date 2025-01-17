// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice"

func Armappservice() []*Table {
	tables := []*Table{
		{
			NewFunc:        armappservice.NewCertificateOrdersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.CertificateRegistration/certificateOrders",
			Namespace:      "Microsoft.CertificateRegistration",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_CertificateRegistration)`,
			Pager:          `NewListPager`,
			ResponseStruct: "CertificateOrdersClientListResponse",
		},
		{
			NewFunc:        armappservice.NewCertificatesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/certificates",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "CertificatesClientListResponse",
		},
		{
			NewFunc:        armappservice.NewDeletedWebAppsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DeletedWebAppsClientListResponse",
		},
		{
			NewFunc:        armappservice.NewDomainsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/domains",
			Namespace:      "Microsoft.DomainRegistration",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DomainRegistration)`,
			Pager:          `NewListPager`,
			ResponseStruct: "DomainsClientListResponse",
		},
		{
			NewFunc:        armappservice.NewEnvironmentsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/hostingEnvironments",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "EnvironmentsClientListResponse",
		},
		{
			NewFunc:        armappservice.NewPlansClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/serverfarms",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "PlansClientListResponse",
		},
		{
			NewFunc:        armappservice.NewRecommendationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/recommendations",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "RecommendationsClientListResponse",
		},
		{
			NewFunc:        armappservice.NewResourceHealthMetadataClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/resourceHealthMetadata",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ResourceHealthMetadataClientListResponse",
		},
		{
			NewFunc:        armappservice.NewStaticSitesClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/staticSites",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "StaticSitesClientListResponse",
		},
		{
			NewFunc:        armappservice.NewTopLevelDomainsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DomainRegistration/topLevelDomains",
			Namespace:      "Microsoft.DomainRegistration",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_DomainRegistration)`,
			Pager:          `NewListPager`,
			ResponseStruct: "TopLevelDomainsClientListResponse",
		},
		{
			NewFunc:        armappservice.NewWebAppsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Web/sites",
			Namespace:      "Microsoft.Web",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_Web)`,
			Pager:          `NewListPager`,
			ResponseStruct: "WebAppsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armappservice())
}
