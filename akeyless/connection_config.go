package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func ConnectionConfig() *plugin.PluginConfigSchema {
	return &plugin.PluginConfigSchema{
		Fields: map[string]*plugin.Schema{
			"access_id": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Access ID",
			},
			"access_type": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Access Type (access_key/password/saml/ldap/k8s/azure_ad/oidc/aws_iam/universal_identity/jwt/gcp/cert)",
			},
			"access_key": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Access key (relevant only for access-type=access_key)",
			},
			"cloud_id": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "The cloud identity (relevant only for access-type=azure_ad,aws_iam,gcp)",
			},
			"uid_token": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "The universal_identity token (relevant only for access-type=universal_identity)",
			},
			"jwt": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "The Json Web Token (relevant only for access-type=jwt/oidc)",
			},
			"admin_password": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Password (relevant only for access-type=password)",
			},
			"admin_email": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Email (relevant only for access-type=password)",
			},
			"account_id": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Account id (relevant only for access-type=password where the email address is associated with more than one account)",
			},
			"oidc_sp": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "OIDC Service Provider (relevant only for access-type=oidc, inferred if empty), supported SPs: google, github",
			},
			"ldap_proxy_url": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Address URL for LDAP proxy (relevant only for access-type=ldap)",
			},
			"username": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "LDAP username (relevant only for access-type=ldap)",
			},
			"password": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "LDAP password (relevant only for access-type=ldap)",
			},
			"gcp_audience": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "GCP audience to use in signed JWT (relevant only for access-type=gcp)",
			},
			"gateway_url": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Gateway URL for the K8S authenticated (relevant only for access-type=k8s/oauth2)",
			},
			"k8s_auth_config_name": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "The K8S Auth config name (relevant only for access-type=k8s)",
			},
			"k8s_service_account_token": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "The K8S service account token",
			},
			"cert_file_name": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Name of the cert file to use (relevant only for access-type=cert)",
			},
			"cert_data": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type=cert)",
			},
			"key_file_name": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Name of the private key file to use (relevant only for access-type=cert)",
			},
			"key_data": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "Private key data encoded in base64. Used if file was not provided.(relevant only for access-type=cert)",
			},
			"debug": {
				Type:        plugin.TypeBool,
				Optional:    true,
				Description: "Set to 'true' for a printout of the authorization jwts'",
			},
			"json": {
				Type:        plugin.TypeBool,
				Optional:    true,
				Description: "Set output format to JSON",
			},
			"jq_expression": {
				Type:        plugin.TypeString,
				Optional:    true,
				Description: "JQ expression to filter result output",
			},
			"no_creds_cleanup": {
				Type:        plugin.TypeBool,
				Optional:    true,
				Description: "Do not clean local temporary expired creds",
			},
		},
	}
}
