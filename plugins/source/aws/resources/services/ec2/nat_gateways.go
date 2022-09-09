package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2NatGateways() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_nat_gateways",
		Description: "Describes a NAT gateway.",
		Resolver:    fetchEc2NatGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"natgateway", *resource.Item.(types.NatGateway).NatGatewayId}, nil
				}),
			},
			{
				Name:            "id",
				Description:     "The ID of the NAT gateway.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("NatGatewayId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "create_time",
				Description: "The date and time the NAT gateway was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "delete_time",
				Description:   "The date and time the NAT gateway was deleted, if applicable.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "failure_code",
				Description:   "If the NAT gateway could not be created, specifies the error code for the failure.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "failure_message",
				Description:   "If the NAT gateway could not be created, specifies the error message for the failure, that corresponds to the error code.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "provisioned_bandwidth",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("ProvisionedBandwidth"),
				IgnoreInTests: true,
			},
			{
				Name:        "state",
				Description: "The state of the NAT gateway.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "The ID of the subnet in which the NAT gateway is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags for the NAT gateway.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the NAT gateway is located.",
				Type:        schema.TypeString,
			},
			{
				Name:        "nat_gateway_addresses",
				Description: "Describes the IP addresses and network interface associated with a NAT gateway.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("NatGatewayAddresses"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2NatGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeNatGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeNatGateways(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.NatGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
