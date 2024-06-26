package main

import (
	"github.com/akeylesslabs/steampipe-plugin-akeyless/akeyless"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// The version injected by goreleaser
// https://goreleaser.com/cookbooks/using-main.version/

var (
	version = "0.0.0-dev"
)

func main() {
	akeyless.PluginVersion = version
	plugin.Serve(&plugin.ServeOpts{PluginFunc: akeyless.Plugin, PluginName: akeyless.PluginName})
}
