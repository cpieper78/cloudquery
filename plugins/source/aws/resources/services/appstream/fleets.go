// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Fleets() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_fleets",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Fleet.html`,
		Resolver:    fetchAppstreamFleets,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appstream2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "compute_capacity_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ComputeCapacityStatus"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disconnect_timeout_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DisconnectTimeoutInSeconds"),
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "domain_join_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainJoinInfo"),
			},
			{
				Name:     "enable_default_internet_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDefaultInternetAccess"),
			},
			{
				Name:     "fleet_errors",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FleetErrors"),
			},
			{
				Name:     "fleet_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FleetType"),
			},
			{
				Name:     "iam_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamRoleArn"),
			},
			{
				Name:     "idle_disconnect_timeout_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("IdleDisconnectTimeoutInSeconds"),
			},
			{
				Name:     "image_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageArn"),
			},
			{
				Name:     "image_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageName"),
			},
			{
				Name:     "max_concurrent_sessions",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxConcurrentSessions"),
			},
			{
				Name:     "max_user_duration_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxUserDurationInSeconds"),
			},
			{
				Name:     "platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "session_script_s3_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SessionScriptS3Location"),
			},
			{
				Name:     "stream_view",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamView"),
			},
			{
				Name:     "usb_device_filter_strings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("UsbDeviceFilterStrings"),
			},
			{
				Name:     "vpc_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VpcConfig"),
			},
		},
	}
}