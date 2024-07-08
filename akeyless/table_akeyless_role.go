package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	rolesTableName        = "akeyless_role"
	rolesTableDescription = "Akeyless access roles"
)

func tableRoles() *plugin.Table {

	return &plugin.Table{
		Name:        rolesTableName,
		Description: rolesTableDescription,
		List: &plugin.ListConfig{
			Hydrate: hydrateListFetchFunc(sdkApiClientProvider{}, rolesFetcher{}),
		},
		Columns: rolesColumns(),
	}
}

func rolesColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "role_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the role.",
		},
		{
			Name:        "creation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the role was created.",
		},
		{
			Name:        "modification_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the role was last modified.",
		},
		{
			Name:        "access_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the role was last accessed.",
			Transform:   transform.FromMethod("GetAccessDateDisplay").Transform(convertDisplayedDate),
		},
		{
			Name:        "rules",
			Type:        proto.ColumnType_JSON,
			Description: "The rules associated with the role.",
		},
		{
			Name:        "role_auth_methods_assoc",
			Type:        proto.ColumnType_JSON,
			Description: "The authentication methods associated with the role.",
		},
	}
}
