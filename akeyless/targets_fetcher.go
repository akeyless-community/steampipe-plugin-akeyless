package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
)

type targetsFetcher struct {
}

type targetPaginator akeyless_sdk.ListTargetsOutput

func (p targetPaginator) GetPageItems() []akeyless_sdk.Target {
	tmp := akeyless_sdk.ListTargetsOutput(p)
	return tmp.GetTargets()
}

func (p targetPaginator) NextPageToken() *string {
	if p.Targets == nil || len(*p.Targets) == 0 {
		return nil
	}
	return p.NextPage
}

func (f targetsFetcher) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[akeyless_sdk.Target], *http.Response, error) {
	out, httpResp, err := apiService.ListTargets(ctx).Body(akeyless_sdk.ListTargets{Token: token, PaginationToken: paginationToken}).Execute()
	return targetPaginator(out), httpResp, err
}
