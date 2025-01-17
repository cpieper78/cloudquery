// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"

func init() {
	tables := []Table{
		{
			Service:        "armcontainerservice",
			Name:           "managed_clusters",
			Struct:         &armcontainerservice.ManagedCluster{},
			ResponseStruct: &armcontainerservice.ManagedClustersClientListResponse{},
			Client:         &armcontainerservice.ManagedClustersClient{},
			ListFunc:       (&armcontainerservice.ManagedClustersClient{}).NewListPager,
			NewFunc:        armcontainerservice.NewManagedClustersClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/managedClusters",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ContainerService)`,
			ExtraColumns:   DefaultExtraColumns,
		},
		{
			Service:        "armcontainerservice",
			Name:           "snapshots",
			Struct:         &armcontainerservice.Snapshot{},
			ResponseStruct: &armcontainerservice.SnapshotsClientListResponse{},
			Client:         &armcontainerservice.SnapshotsClient{},
			ListFunc:       (&armcontainerservice.SnapshotsClient{}).NewListPager,
			NewFunc:        armcontainerservice.NewSnapshotsClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/snapshots",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.NamespaceMicrosoft_ContainerService)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
