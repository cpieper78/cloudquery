package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var Version = "Development"

func Snyk() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"snyk",
		Version,
		tables(),
		client.Configure,
	)
}
