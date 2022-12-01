package settings

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	if c.BaseURL == "" {
		c.Logger().Warn().Msg("not supported for GitLab SaaS, skipping...")
		return nil
	}
	setting, _, err := c.Gitlab.Settings.GetSettings(nil)
	if err != nil {
		return err
	}

	res <- setting

	return nil
}