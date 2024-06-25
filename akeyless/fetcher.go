package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type dataPaginator[T any] interface {
	GetPageItems() []T
	NextPageToken() *string
}

type apiDataFetcher[T any] interface {
	FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[T], *http.Response, error)
}

func hydrateListFetchFunc[T any](apiProvider apiClientProvider, fetcher apiDataFetcher[T]) plugin.HydrateFunc {
	return func(ctx context.Context, qd *plugin.QueryData, hd *plugin.HydrateData) (interface{}, error) {
		return listTableTemplate(ctx, qd, hd, apiProvider, fetcher)
	}
}
