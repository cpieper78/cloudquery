// Code generated by codegen; DO NOT EDIT.

package run

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/run/v2"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:      "gcp_run_services",
		Resolver:  fetchServices,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uid"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
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
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
			{
				Name:     "delete_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("DeleteTime"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("ExpireTime"),
			},
			{
				Name:     "creator",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Creator"),
			},
			{
				Name:     "last_modifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastModifier"),
			},
			{
				Name:     "client",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Client"),
			},
			{
				Name:     "client_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientVersion"),
			},
			{
				Name:     "ingress",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Ingress"),
			},
			{
				Name:     "launch_stage",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("LaunchStage"),
			},
			{
				Name:     "binary_authorization",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BinaryAuthorization"),
			},
			{
				Name:     "template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Template"),
			},
			{
				Name:     "traffic",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Traffic"),
			},
			{
				Name:     "observed_generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ObservedGeneration"),
			},
			{
				Name:     "terminal_condition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TerminalCondition"),
			},
			{
				Name:     "conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Conditions"),
			},
			{
				Name:     "latest_ready_revision",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestReadyRevision"),
			},
			{
				Name:     "latest_created_revision",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LatestCreatedRevision"),
			},
			{
				Name:     "traffic_statuses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TrafficStatuses"),
			},
			{
				Name:     "uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uri"),
			},
			{
				Name:     "reconciling",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Reconciling"),
			},
			{
				Name:     "etag",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Etag"),
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListServicesRequest{
		Parent: "projects/" + c.ProjectId + "locations/-",
	}
	it := c.Services.RunServicesClient.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
