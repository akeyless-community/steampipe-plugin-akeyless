package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin() *plugin.Plugin {
	return &plugin.Plugin{
		PluginInfo: &plugin.PluginInfo{
			Name:        "akeyless",
			Description: "Steampipe plugin for Akeyless",
		},
		ConnectionConfig: map[string]*plugin.PluginConfigSchema{
			// TODO: Add connection configurations
		},
		TableMap: map[string]*plugin.Table{
			// TODO: Add table definitions
		},
	}
}
