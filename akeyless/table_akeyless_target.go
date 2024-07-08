package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	targetsTableName        = "akeyless_target"
	targetsTableDescription = "Akeyless endpoint targets"
)

func tableTargets() *plugin.Table {

	return &plugin.Table{
		Name:        targetsTableName,
		Description: targetsTableDescription,

		List: &plugin.ListConfig{
			Hydrate: hydrateListFetchFunc(sdkApiClientProvider{}, targetsFetcher{}),
		},
		Columns: targetsColumns(),
	}
}

func targetsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the target.",
		},
		{
			Name:        "creation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the target was created.",
		},
		{
			Name:        "modification_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the target was last modified.",
		},
		{
			Name:        "access_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the target was last accessed.",
			Transform:   transform.FromMethod("GetAccessDateDisplay").Transform(convertDisplayedDate),
		},
		{
			Name:        "target_type",
			Type:        proto.ColumnType_STRING,
			Description: "The type of the target (e.g., database, cloud platform, server).",
		},
		{
			Name:        "with_customer_fragment",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the target includes a customer-specific fragment.",
		},
		{
			Name:        "protection_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the protection key used for securing the target.",
		},
		{
			Name:        "client_permissions",
			Type:        proto.ColumnType_JSON,
			Description: "The permissions assigned to clients for this target.",
		},
		{
			Name:        "last_version",
			Type:        proto.ColumnType_INT,
			Description: "The last version number of the target configuration.",
		},
		{
			Name:        "attributes",
			Type:        proto.ColumnType_JSON,
			Description: "Additional attributes associated with the target.",
		},
		{
			Name:        "is_access_request_enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if access requests are enabled for this target.",
		},
		{
			Name:        "access_request_status",
			Type:        proto.ColumnType_STRING,
			Description: "The current status of access requests for this target.",
		},
	}
}
