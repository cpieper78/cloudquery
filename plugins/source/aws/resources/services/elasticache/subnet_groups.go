// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_subnet_groups",
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CacheSubnetGroup.html`,
		Resolver:    fetchElasticacheSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "cache_subnet_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheSubnetGroupDescription"),
			},
			{
				Name:     "cache_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheSubnetGroupName"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "supported_network_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedNetworkTypes"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
