package akeyless

import (
	"context"

	"errors"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/akeylesslabs/akeyless-go/v2"
)

func tableAuthMethod() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_auth_method",
		Description: "Akeyless Auth Methods",
		List: &plugin.ListConfig{
			Hydrate: listAuthMethods,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("path"),
			Hydrate:    getAuthMethod,
		},
		Columns: authMethodColumns(),
	}
}

func listAuthMethods(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	listAuthMethodBody := akeyless.NewListAuthMethods()
	listAuthMethodBody.Token = conn.token

	listAuthMethodResponse, _, err := conn.client.ListAuthMethods(ctx).Body(*listAuthMethodBody).Execute()
	if err != nil {
		return nil, err
	}

	authMethods := listAuthMethodResponse.GetAuthMethods()

	for _, authMethod := range authMethods {

		plugin.Logger(ctx).Trace("listAuthMethods", "authMethod", authMethod)

		authMethodListItem := &AuthMethod{}

		authMethodListItemTranslate(authMethod, authMethodListItem)

		d.StreamListItem(ctx, authMethodListItem)
	}

	return nil, nil
}

// newFunction creates a new AuthMethod object from the given akeyless.AuthMethod and AuthMethodListItem.
// It sets the values of the AuthMethodListItem fields based on the values of the akeyless.AuthMethod fields.
func authMethodListItemTranslate(authMethod akeyless.AuthMethod, authMethodListItem *AuthMethod) {
	authMethodName := authMethod.GetAuthMethodName()
	authMethodListItem.Path = authMethodName

	var accessIdToUse = authMethod.GetAuthMethodAccessId()
	if authMethod.AccessInfo != nil && authMethod.AccessInfo.RulesType != nil && *authMethod.AccessInfo.RulesType == "email_pass" {
		accessIdToUse = authMethod.AccessInfo.GetAccessIdAlias()
	}

	authMethodListItem.AuthMethodAccessId = accessIdToUse

	authMethodAccountId := authMethod.GetAccountId()
	authMethodListItem.AccountId = authMethodAccountId

	accessInfo := authMethod.GetAccessInfo()

	accessInfoRulesType := accessInfo.GetRulesType()
	authMethodListItem.AccessInfoRulesType = accessInfoRulesType

	accessInfoJwtTtl := accessInfo.GetJwtTtl()
	authMethodListItem.AccessInfoJwtTtl = accessInfoJwtTtl

	accessInfoAccessExpires := accessInfo.GetAccessExpires()
	authMethodListItem.AccessInfoAccessExpires = accessInfoAccessExpires

	accessInfoCidrWhiteList := accessInfo.GetCidrWhitelist()
	authMethodListItem.AccessInfoCidrWhiteList = accessInfoCidrWhiteList

	accessInfoGwCidrWhiteList := accessInfo.GetGwCidrWhitelist()
	authMethodListItem.AccessInfoGwCidrWhiteList = accessInfoGwCidrWhiteList

	accessInfoForceSubClaims := accessInfo.GetForceSubClaims()
	authMethodListItem.AccessInfoForceSubClaims = accessInfoForceSubClaims

	creationDate := authMethod.GetCreationDate()
	authMethodListItem.CreationDate = creationDate.String()

	modificationDate := authMethod.GetModificationDate()
	authMethodListItem.ModificationDate = modificationDate.String()
}

func getAuthMethod(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	q := d.EqualsQuals
	path := q["path"].GetStringValue()

	getAuthMethodBody := akeyless.NewGetAuthMethodWithDefaults()
	getAuthMethodBody.Token = conn.token
	getAuthMethodBody.Name = path

	getAuthMethodResponse, httpResponse, err := conn.client.GetAuthMethod(ctx).Body(*getAuthMethodBody).Execute()
	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode == 404 {
		return nil, errors.New("auth method [" + path + "] not found")
	} else {
		authMethod := getAuthMethodResponse
		authMethodListItem := &AuthMethod{}
		authMethodListItemTranslate(authMethod, authMethodListItem)
		return authMethodListItem, nil
	}
}

type AuthMethod struct {
	Path                      string
	AuthMethodName            string
	AuthMethodAccessId        string
	AccountId                 string
	AccessInfoRulesType       string
	AccessInfoJwtTtl          int64
	AccessInfoAccessExpires   int64
	AccessInfoCidrWhiteList   string
	AccessInfoGwCidrWhiteList string
	AccessInfoForceSubClaims  bool
	CreationDate              string
	ModificationDate          string
}

func authMethodColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "path",
			Type:        proto.ColumnType_STRING,
			Description: "The full path of the auth method which includes the name.",
		},
		{
			Name:        "auth_method_name",
			Type:        proto.ColumnType_STRING,
			Description: "The full path of the auth method which includes the name.",
			Transform:   transform.FromField("Path"),
		},
		{
			Name:        "auth_method_access_id",
			Type:        proto.ColumnType_STRING,
			Description: "The full virtual file folder path of the auth method which includes the name.",
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The account ID of the auth method.",
		},
		{
			Name:        "access_info_rules_type",
			Type:        proto.ColumnType_STRING,
			Description: "The rules type to use for association of auth method to access role for the auth method.",
		},
		{
			Name:        "access_info_jwt_ttl",
			Type:        proto.ColumnType_INT,
			Description: "The JWT TTL for the auth method.",
		},
		{
			Name:        "access_info_access_expires",
			Type:        proto.ColumnType_INT,
			Description: "The access expiration date. This parameter is optional. Leave it empty for access to continue without an expiration date.",
		},
		{
			Name:        "access_info_cidr_white_list",
			Type:        proto.ColumnType_STRING,
			Description: "Enter a comma-separated list of CIDR blocks from which the client can issue calls to the proxy. By 'client,' we mean CURL, SDK, etc. This parameter is optional. Leave it empty for unrestricted access.",
		},
		{
			Name:        "access_info_gw_cidr_white_list",
			Type:        proto.ColumnType_STRING,
			Description: "Comma separated CIDR blocks. If specified, the Gateway using this IP range will be trusted to forward the original client IP. If empty, the Gateway's IP address will be used.",
		},
		{
			Name:        "access_info_force_sub_claims",
			Type:        proto.ColumnType_BOOL,
			Description: "If set to true, access roles will enforce role-association must include sub claims.",
		},
		{
			Name:        "creation_date",
			Type:        proto.ColumnType_STRING,
			Description: "The creation date of the auth method.",
		},
		{
			Name:        "modification_date",
			Type:        proto.ColumnType_STRING,
			Description: "The modification date of the auth method.",
		},
	}
}
