package akeyless

import (
	"context"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	authMethodsTableName   = "akeyless_auth_method"
	authMethodsDescription = "Akeyless authentication methods"
)

func tableAuthMethods() *plugin.Table {

	return &plugin.Table{
		Name:        authMethodsTableName,
		Description: authMethodsDescription,
		List: &plugin.ListConfig{
			Hydrate: hydrateListFetchFunc(sdkApiClientProvider{}, authMethodsFetcher{}),
		},
		Columns: authMethodsColumns(),
	}
}

func authMethodsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "auth_method_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the authentication method.",
		},
		{
			Name:        "creation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the authentication method was created.",
		},
		{
			Name:        "modification_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the authentication method was last modified.",
		},
		{
			Name:        "access_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The last accessed date of the authentication method, formatted for display.",
			Transform:   transform.FromMethod("GetAccessDateDisplay").Transform(convertDisplayedDate),
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The account identifier associated with the authentication method.",
			Transform:   transform.FromMethod("GetAccountId"),
		},
		{
			Name:        "ttl",
			Type:        proto.ColumnType_INT,
			Description: "The time-to-live (TTL) for the authentication method.",
			Transform:   transform.FromMethod("GetAccessInfo").TransformP(transformAuthMethodAccessInfo, "ttl"),
		},
		{
			Name:        "rules_type",
			Type:        proto.ColumnType_STRING,
			Description: "The type of rules associated with the authentication method.",
			Transform:   transform.FromMethod("GetAccessInfo").TransformP(transformAuthMethodAccessInfo, "rules_type"),
		},
		{
			Name:        "force_sub_claims",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if sub-claims are forced for the authentication method.",
			Transform:   transform.FromMethod("GetAccessInfo").TransformP(transformAuthMethodAccessInfo, "force_sub_claims"),
		},
		{
			Name:        "access_info",
			Type:        proto.ColumnType_JSON,
			Description: "Additional access information for the authentication method.",
			Transform:   transform.FromMethod("GetAccessInfo"),
		},
	}
}

func transformAuthMethodAccessInfo(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	const logPrefix = authMethodsTableName + ".transformAuthMethodAccessInfo"

	if amInfo, ok := input.Value.(akeyless_sdk.AuthMethodAccessInfo); ok {
		field := input.Param.(string)
		switch field {
		case "ttl":
			return amInfo.GetJwtTtl(), nil
		case "rules_type":
			return amInfo.GetRulesType(), nil
		case "force_sub_claims":
			return amInfo.GetForceSubClaims(), nil
		}
	}

	plugin.Logger(ctx).Warn(logPrefix, "failed to cast input value to AuthMethodAccessInfo", input.Value)

	return nil, nil
}
