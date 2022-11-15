package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/ses/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSesIdentities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2

	config := sesv2.ListEmailIdentitiesInput{}
	p := sesv2.NewListEmailIdentitiesPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.EmailIdentities
	}
	return nil
}

func getEmailIdentity(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sesv2
	ei := resource.Item.(types.IdentityInfo)

	getOutput, err := svc.GetEmailIdentity(ctx, &sesv2.GetEmailIdentityInput{EmailIdentity: ei.IdentityName})
	if err != nil {
		return err
	}

	resource.Item = &models.EmailIdentityWrapper{
		IdentityName:   ei.IdentityName,
		SendingEnabled: ei.SendingEnabled,

		GetEmailIdentityOutput: getOutput,
	}
	return nil
}

func resolveEmailIdentityArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(models.EmailIdentityWrapper)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   "ses",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "identity/" + aws.ToString(item.IdentityName),
	}
	return resource.Set(c.Name, arn.String())
}
