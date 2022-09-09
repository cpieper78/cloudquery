package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

func Ec2Images() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_images",
		Description:   "Describes an image.",
		Resolver:      fetchEc2Images,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
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
				Resolver: client.ResolveARNWithRegion(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"image", *resource.Item.(types.Image).ImageId}, nil
				}),
			},
			{
				Name:            "id",
				Description:     "The ID of the AMI.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ImageId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_date",
				Description: "The date and time the image was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("CreationDate"),
			},
			{
				Name:        "description",
				Description: "The description of the AMI that was provided during image creation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_location",
				Description: "The location of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_owner_alias",
				Description: "The AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "image_type",
				Description: "The type of image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "kernel_id",
				Description: "The kernel associated with the image, if any.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the AMI that was provided during image creation.",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner_id",
				Description: "The AWS account ID of the image owner.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform",
				Description: "This value is set to windows for Windows AMIs; otherwise, it is blank.",
				Type:        schema.TypeString,
			},
			{
				Name:        "platform_details",
				Description: "The platform details associated with the billing code of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "product_codes",
				Description: "Any product codes associated with the AMI.",
				Type:        schema.TypeJSON,
				Resolver:    resolveEc2imageProductCodes,
			},
			{
				Name:        "public",
				Description: "Indicates whether the image has public launch permissions.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ramdisk_id",
				Description: "The RAM disk associated with the image, if any.",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_type",
				Description: "The type of root device used by the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sriov_net_support",
				Description: "Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The current state of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateReason"),
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the image.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "usage_operation",
				Description: "The operation of the Amazon EC2 instance and the billing code that is associated with the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "virtualization_type",
				Description: "The type of virtualization of the AMI.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_launched_time",
				Description: "The timestamp of the last time the AMI was used for an EC2 instance launch.",
				Type:        schema.TypeTimestamp,
				Resolver:    resolveEc2ImageLastLaunchedTime,
			},
			{
				Name:        "deprecation_time",
				Description: "The date and time to deprecate the AMI.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DeprecationTime"),
			},
			{
				Name:        "block_device_mappings",
				Description: "Describes a block device mapping.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("BlockDeviceMappings"),
			},
		},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func fetchEc2Images(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)

	svc := c.Services().EC2
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// fetch ec2.Images owned by this account
		response, err := svc.DescribeImages(ctx, &ec2.DescribeImagesInput{Owners: []string{"self"}})
		if err != nil {
			return err
		}
		res <- response.Images
		return nil
	})

	g.Go(func() error {
		// fetch ec2.Images that are shared with this account
		response, err := svc.DescribeImages(ctx, &ec2.DescribeImagesInput{ExecutableUsers: []string{"self"}})
		if err != nil {
			return err
		}
		res <- response.Images
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func resolveEc2imageProductCodes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Image)
	productCodes := map[string]string{}
	for _, t := range r.ProductCodes {
		productCodes[*t.ProductCodeId] = string(t.ProductCodeType)
	}
	return resource.Set(c.Name, productCodes)
}

func resolveEc2ImageLastLaunchedTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	image := resource.Item.(types.Image)
	svc := cl.Services().EC2
	if *image.OwnerId != cl.AccountID {
		return nil
	}
	opts := ec2.DescribeImageAttributeInput{
		Attribute: types.ImageAttributeNameLastLaunchedTime,
		ImageId:   image.ImageId,
	}
	result, err := svc.DescribeImageAttribute(ctx, &opts, func(options *ec2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if result.LastLaunchedTime == nil || result.LastLaunchedTime.Value == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339, *result.LastLaunchedTime.Value)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, t)
}
