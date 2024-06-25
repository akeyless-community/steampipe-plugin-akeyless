package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
)

type gatewaysFetcher struct {
}

type gatewayPaginator akeyless_sdk.GatewaysListResponse

func (p gatewayPaginator) GetPageItems() []akeyless_sdk.GwClusterIdentity {
	tmp := akeyless_sdk.GatewaysListResponse(p)
	return tmp.GetClusters()
}

func (p gatewayPaginator) NextPageToken() *string {
	return nil
}

func (f gatewaysFetcher) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[akeyless_sdk.GwClusterIdentity], *http.Response, error) {
	out, httpResp, err := apiService.ListGateways(ctx).Body(akeyless_sdk.ListGateways{Token: token}).Execute()
	return gatewayPaginator(out), httpResp, err
}
