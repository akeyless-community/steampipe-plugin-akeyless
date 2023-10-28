package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type akeylessConfig struct {
	CLIPath      *string `cty:"cli_path"`      // Path to the Akeyless CLI executable
	Profile      *string `cty:"profile"`       // Name of the Akeyless CLI profile to use
	AkeylessPath *string `cty:"akeyless_path"` // Path to the .akeyless directory
	ExpiryBuffer *string `cty:"expiry_buffer"` // Buffer time before token expiry to trigger re-authentication "2h" or "10m" (default)
	Debug        *string `cty:"debug"`         // Debug flag to enable or disable debug logging
}

var ConfigSchema = map[string]*schema.Attribute{
	"cli_path": {
		Type: schema.TypeString,
	},
	"profile": {
		Type: schema.TypeString,
	},
	"akeyless_path": {
		Type: schema.TypeString,
	},
	"expiry_buffer": {
		Type: schema.TypeString,
	},
	"debug": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &akeylessConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) akeylessConfig {
	if connection == nil || connection.Config == nil {
		return akeylessConfig{}
	}
	config, _ := connection.Config.(akeylessConfig)
	return config
}
