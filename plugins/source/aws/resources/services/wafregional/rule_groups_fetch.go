package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchWafregionalRuleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	var params wafregional.ListRuleGroupsInput
	for {
		result, err := svc.ListRuleGroups(ctx, &params, func(o *wafregional.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, g := range result.RuleGroups {
			detail, err := svc.GetRuleGroup(
				ctx,
				&wafregional.GetRuleGroupInput{RuleGroupId: g.RuleGroupId},
				func(o *wafregional.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if detail.RuleGroup == nil {
				continue
			}
			res <- *detail.RuleGroup
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return nil
}

func resolveWafregionalRuleGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.RuleGroup)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.WAFRegional),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "rulegroup/" + aws.ToString(item.RuleGroupId),
	}
	return resource.Set(c.Name, arn.String())
}

func resolveWafregionalRuleGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Wafregional
	item := resource.Item.(types.RuleGroup)
	arn := arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.WAFRegional),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "rulegroup/" + aws.ToString(item.RuleGroupId),
	}
	params := wafregional.ListTagsForResourceInput{ResourceARN: aws.String(arn.String())}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			return err
		}
		if result.TagInfoForResource != nil {
			client.TagsIntoMap(result.TagInfoForResource.TagList, tags)
		}
		if aws.ToString(result.NextMarker) == "" {
			break
		}
		params.NextMarker = result.NextMarker
	}
	return resource.Set(c.Name, tags)
}
