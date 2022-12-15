package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "TODO"

func main() {
	serve.Source(plugin.Snyk(), serve.WithSourceSentryDSN(sentryDSN))
}
