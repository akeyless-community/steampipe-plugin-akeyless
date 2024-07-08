package akeyless

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	gatewaysTableName        = "akeyless_gateway"
	gatewaysTableDescription = "Akeyless gateways"
)

func tableGateways() *plugin.Table {

	return &plugin.Table{
		Name:        gatewaysTableName,
		Description: gatewaysTableDescription,

		List: &plugin.ListConfig{
			Hydrate: hydrateListFetchFunc(sdkApiClientProvider{}, gatewaysFetcher{}),
		},
		Columns: gatewayColumns(),
	}
}

func gatewayColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "gateway_id",
			Type:        proto.ColumnType_INT,
			Description: "The unique identifier for the gateway.",
			Transform:   transform.FromMethod("GetId"),
		},
		{
			Name:        "display_name",
			Type:        proto.ColumnType_STRING,
			Description: "The display name of the gateway.",
		},
		{
			Name:        "cluster_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the cluster associated with the gateway.",
		},
		{
			Name:        "cluster_url",
			Type:        proto.ColumnType_STRING,
			Description: "The URL of the cluster associated with the gateway.",
			Transform:   transform.FromMethod("GetClusterUrl"),
		},
		{
			Name:        "customer_fragments",
			Type:        proto.ColumnType_JSON,
			Description: "Customer-specific fragments associated with the gateway.",
		},
		{
			Name:        "status",
			Type:        proto.ColumnType_STRING,
			Description: "The current status of the gateway.",
		},
		{
			Name:        "allowed",
			Type:        proto.ColumnType_BOOL,
			Description: "Indicates if the gateway is allowed.",
		},
		{
			Name:        "allowed_access_ids",
			Type:        proto.ColumnType_JSON,
			Description: "A list of allowed access IDs for the gateway.",
		},
		{
			Name:        "default_protection_key_id",
			Type:        proto.ColumnType_INT,
			Description: "The default protection key ID used for the gateway.",
			Transform:   transform.FromMethod("GetDefaultProtectionKeyId"),
		},
		{
			Name:        "default_secret_location",
			Type:        proto.ColumnType_STRING,
			Description: "The default location where secrets are stored for the gateway.",
		},
	}
}
