// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/users"
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/user_identities"
)

func tables() []*schema.Table {
	return []*schema.Table{
        users.(),
        user_identities.(),
	}
}
