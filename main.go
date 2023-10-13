package main

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/user/steampipe-plugin-akeyless/akeyless"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: akeyless.Plugin})
}
