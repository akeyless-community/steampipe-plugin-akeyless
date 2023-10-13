package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type akeylessConfig struct {
	AccessID               *string `cty:"access_id"`
	AccessType             *string `cty:"access_type"`
	AccessKey              *string `cty:"access_key"`
	CloudID                *string `cty:"cloud_id"`
	UIDToken               *string `cty:"uid_token"`
	JWT                    *string `cty:"jwt"`
	AdminPassword          *string `cty:"admin_password"`
	AdminEmail             *string `cty:"admin_email"`
	AccountID              *string `cty:"account_id"`
	OIDCSP                 *string `cty:"oidc_sp"`
	LdapProxyURL           *string `cty:"ldap_proxy_url"`
	Username               *string `cty:"username"`
	Password               *string `cty:"password"`
	GcpAudience            *string `cty:"gcp_audience"`
	GatewayURL             *string `cty:"gateway_url"`
	K8SAuthConfigName      *string `cty:"k8s_auth_config_name"`
	K8SServiceAccountToken *string `cty:"k8s_service_account_token"`
	CertFileName           *string `cty:"cert_file_name"`
	CertData               *string `cty:"cert_data"`
	KeyFileName            *string `cty:"key_file_name"`
	KeyData                *string `cty:"key_data"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"access_id": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access_type": {
		Type:     schema.TypeString,
		Required: true,
	},
	"access_key": {
		Type: schema.TypeString,
	},
	"cloud_id": {
		Type: schema.TypeString,
	},
	"uid_token": {
		Type: schema.TypeString,
	},
	"jwt": {
		Type: schema.TypeString,
	},
	"admin_password": {
		Type: schema.TypeString,
	},
	"admin_email": {
		Type: schema.TypeString,
	},
	"account_id": {
		Type: schema.TypeString,
	},
	"oidc_sp": {
		Type: schema.TypeString,
	},
	"ldap_proxy_url": {
		Type: schema.TypeString,
	},
	"username": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"gcp_audience": {
		Type: schema.TypeString,
	},
	"gateway_url": {
		Type: schema.TypeString,
	},
	"k8s_auth_config_name": {
		Type: schema.TypeString,
	},
	"k8s_service_account_token": {
		Type: schema.TypeString,
	},
	"cert_file_name": {
		Type: schema.TypeString,
	},
	"cert_data": {
		Type: schema.TypeString,
	},
	"key_file_name": {
		Type: schema.TypeString,
	},
	"key_data": {
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
