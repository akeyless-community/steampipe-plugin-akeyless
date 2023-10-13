package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func ConnectionConfig() *plugin.PluginConfigSchema {
	return &plugin.PluginConfigSchema{
		Fields: map[string]*plugin.Schema{
			"access_id": {
				Type:        plugin.TypeString,
				Required:    true,
				Description: "The access ID for Akeyless.",
			},
			"access_key": {
				Type:        plugin.TypeString,
				Required:    true,
				Description: "The access key for Akeyless.",
			},
		},
	}
}
