// Code generated by codegen; DO NOT EDIT.

package storage

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CsiDrivers() *schema.Table {
	return &schema.Table{
		Name:      "k8s_storage_csi_drivers",
		Resolver:  fetchCsiDrivers,
		Multiplex: client.ContextMultiplex,
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OwnerReferences"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Finalizers"),
			},
			{
				Name:     "spec_attach_required",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.AttachRequired"),
			},
			{
				Name:     "spec_pod_info_on_mount",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.PodInfoOnMount"),
			},
			{
				Name:     "spec_volume_lifecycle_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.VolumeLifecycleModes"),
			},
			{
				Name:     "spec_storage_capacity",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.StorageCapacity"),
			},
			{
				Name:     "spec_fs_group_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.FSGroupPolicy"),
			},
			{
				Name:     "spec_token_requests",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.TokenRequests"),
			},
			{
				Name:     "spec_requires_republish",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.RequiresRepublish"),
			},
			{
				Name:     "spec_se_linux_mount",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.SELinuxMount"),
			},
		},
	}
}

func fetchCsiDrivers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client).Client().StorageV1().CSIDrivers()

	opts := metav1.ListOptions{}
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return err
		}
		res <- result.Items
		if result.GetContinue() == "" {
			return nil
		}
		opts.Continue = result.GetContinue()
	}
}