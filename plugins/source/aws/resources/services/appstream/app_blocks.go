// Code generated by codegen; DO NOT EDIT.

package appstream

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AppBlocks() *schema.Table {
	return &schema.Table{
		Name:        "aws_appstream_app_blocks",
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html`,
		Resolver:    fetchAppstreamAppBlocks,
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "setup_script_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SetupScriptDetails"),
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
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "source_s3_location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceS3Location"),
			},
		},
	}
}