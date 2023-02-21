package main

import (
	"github.com/jreyesr/steampipe-plugin-samplerest/samplerest"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: samplerest.Plugin})
}
