package akeyless

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	itemsTableName        = "akeyless_item"
	itemsTableDescription = "Akeyless items"
)

func tableItems() *plugin.Table {
	return &plugin.Table{
		Name:        itemsTableName,
		Description: itemsTableDescription,
		List: &plugin.ListConfig{
			Hydrate: hydrateListFetchFunc(sdkApiClientProvider{}, itemsFetcher{}),
		},
		Columns: itemColumns(),
	}
}

func itemColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "item_id",
			Type:        proto.ColumnType_INT,
			Description: "The unique identifier for the item.",
			Transform:   transform.FromMethod("GetItemId"),
		},
		{
			Name:        "account_id",
			Type:        proto.ColumnType_STRING,
			Description: "The account identifier associated with the item.",
			Transform:   transform.FromMethod("GetDisplayId").Transform(transformItem),
		},
		{
			Name:        "creation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date and time when the item was created.",
		},
		{
			Name:        "modification_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date and time when the item was last modified.",
		},
		{
			Name:        "item_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the item.",
		},
		{
			Name:        "item_sub_type",
			Type:        proto.ColumnType_STRING,
			Description: "The sub-type of the item.",
		},
		{
			Name:        "last_version",
			Type:        proto.ColumnType_INT,
			Description: "The most recent version number of the item.",
		},
		{
			Name:        "with_customer_fragment",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the item includes a customer-specific fragment.",
		},
		{
			Name:        "is_enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the item is currently enabled.",
		},
		{
			Name:        "protection_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the protection key used for securing the item.",
		},
		{
			Name:        "client_permissions",
			Type:        proto.ColumnType_JSON,
			Description: "Permissions assigned to clients for this item.",
		},
		{
			Name:        "item_state",
			Type:        proto.ColumnType_STRING,
			Description: "The current state of the item.",
		},
		{
			Name:        "rotation_interval",
			Type:        proto.ColumnType_INT,
			Description: "The interval at which the item is rotated.",
		},
		{
			Name:        "item_general_info",
			Type:        proto.ColumnType_JSON,
			Description: "General information about the item.",
		},
		{
			Name:        "item_targets_assoc",
			Type:        proto.ColumnType_JSON,
			Description: "Associations between the item and its targets.",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if delete protection is enabled for the item.",
		},
		{
			Name:        "is_access_request_enabled",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if access requests are enabled for this item.",
		},
		{
			Name:        "access_request_status",
			Type:        proto.ColumnType_STRING,
			Description: "The current status of access requests for the item.",
		},
		{
			Name:        "next_rotation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the next rotation for the item is scheduled.",
		},
		{
			Name:        "last_rotation_date",
			Type:        proto.ColumnType_TIMESTAMP,
			Description: "The date when the last rotation for the item occurred.",
		},
		{
			Name:        "auto_rotate",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates whether the item is set to rotate automatically.",
		},
	}
}

func transformItem(ctx context.Context, input *transform.TransformData) (interface{}, error) {
	const logPrefix = itemsTableName + ".transform"

	if displayId, ok := input.Value.(string); ok {
		dashIndex := strings.Index(displayId, "-")
		if dashIndex == -1 {
			plugin.Logger(ctx).Warn(logPrefix, "failed to parse display id ", displayId)
			return nil, nil
		}
		return displayId[:dashIndex], nil
	}

	plugin.Logger(ctx).Warn(logPrefix, "failed to cast input value to string", input.Value)

	return nil, nil
}
