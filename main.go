package main

import (
	"github.com/akeyless-community/steampipe-plugin-akeyless/akeyless"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: akeyless.Plugin})
}
