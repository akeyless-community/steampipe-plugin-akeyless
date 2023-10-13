package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func ConnectionConfig() *plugin.PluginConfigSchema {
	return &plugin.PluginConfigSchema{
		Fields: map[string]*plugin.Schema{
			"access_type": {
				Type:        plugin.TypeString,
				Required:    true,
				Description: "The authentication method to use. Options are: api_key_auth, aws_iam_auth, azure_ad_auth, jwt_auth, email_auth, uid_auth, cert_auth.",
			},
			"api_key_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "API key authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"aws_iam_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "AWS IAM authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"azure_ad_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "Azure AD authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"jwt_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "JWT authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"email_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "Email authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"uid_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "UID authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
			"cert_auth": {
				Type:        plugin.TypeMap,
				Optional:    true,
				Description: "Certificate authentication parameters.",
				Elem: &plugin.Schema{
					Type: plugin.TypeString,
				},
			},
		},
	}
}
