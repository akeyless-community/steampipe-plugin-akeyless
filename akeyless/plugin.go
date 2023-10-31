package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-akeyless",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"akeyless_auth_method": tableAkeylessAuthMethod(),
		},
	}

	return p
}
