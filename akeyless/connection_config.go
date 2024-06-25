package akeyless

import (
	"context"
	"fmt"
	"strings"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type akeylessConfig struct {
	AccessType             *string `cty:"access_type"`
	AccessId               *string `cty:"access_id"`
	AccessKey              *string `cty:"access_key"`
	ApiUrl                 *string `cty:"api_url"`
	Jwt                    *string `cty:"jwt"`
	UidToken               *string `cty:"uid_token"`
	GcpAudience            *string `cty:"gcp_audience"`
	AzureObjectId          *string `cty:"azure_object_id"`
	K8SServiceAccountToken *string `cty:"k8s_service_account_token"`
	K8SAuthConfigName      *string `cty:"k8s_auth_config_name"`
	GatewayCaCert          *string `cty:"gateway_ca_cert"`
}

var configSchema = map[string]*schema.Attribute{
	"access_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access_id": {
		Type:     schema.TypeString,
		Required: false,
	},
	"access_key": {
		Type:     schema.TypeString,
		Required: false,
	},
	"api_url": {
		Type:     schema.TypeString,
		Required: false,
	},
	"jwt": {
		Type:     schema.TypeString,
		Required: false,
	},
	"uid_token": {
		Type:     schema.TypeString,
		Required: false,
	},
	"gcp_audience": {
		Type:     schema.TypeString,
		Required: false,
	},
	"azure_object_id": {
		Type:     schema.TypeString,
		Required: false,
	},
	"k8s_service_account_token": {
		Type:     schema.TypeString,
		Required: false,
	},
	"k8s_auth_config_name": {
		Type:     schema.TypeString,
		Required: false,
	},
	"gateway_ca_cert": {
		Type:     schema.TypeString,
		Required: false,
	},
}

func configInstance() interface{} {
	return &akeylessConfig{}
}

func getPluginConfig(_ context.Context, connection *plugin.Connection) (*akeylessConfig, error) {

	if connection == nil || connection.Config == nil {
		return nil, fmt.Errorf("plugin connection or config is nil")
	}

	config, ok := connection.Config.(akeylessConfig)

	if !ok {
		return nil, fmt.Errorf("plugin config casting to akeylessConfig failed")
	}

	if config.AccessType == nil || strings.TrimSpace(*config.AccessType) == "" {
		return nil, fmt.Errorf("access_type is not set")
	}

	if config.AccessType == nil || strings.TrimSpace(*config.AccessType) == "" {
		return nil, fmt.Errorf("access_type is not set")
	}

	if config.AccessId == nil {
		config.AccessId = new(string)
	}
	if config.AccessKey == nil {
		config.Jwt = new(string)
	}
	if config.ApiUrl == nil || strings.TrimSpace(*config.ApiUrl) == "" {
		config.ApiUrl = akeyless_sdk.PtrString("https://api.akeyless.io")
	}
	if config.Jwt == nil {
		config.Jwt = new(string)
	}
	if config.UidToken == nil {
		config.UidToken = new(string)
	}
	if config.GcpAudience == nil {
		config.GcpAudience = new(string)
	}
	if config.AzureObjectId == nil {
		config.AzureObjectId = new(string)
	}
	if config.K8SServiceAccountToken == nil {
		config.K8SServiceAccountToken = new(string)
	}
	if config.K8SAuthConfigName == nil {
		config.K8SAuthConfigName = new(string)
	}
	if config.GatewayCaCert == nil {
		config.GatewayCaCert = new(string)
	}

	return &config, nil
}
