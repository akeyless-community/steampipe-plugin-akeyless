package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const PluginName string = "steampipe-plugin-akeyless"

var PluginVersion string = ""

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: PluginName,

		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: configInstance,
			Schema:      configSchema,
		},
		DefaultTransform: transform.FromJSONTag(),

		TableMap: map[string]*plugin.Table{
			rolesTableName:       tableRoles(),
			targetsTableName:     tableTargets(),
			authMethodsTableName: tableAuthMethods(),
			itemsTableName:       tableItems(),
			gatewaysTableName:    tableGateways(),
		},
	}

	return p
}
