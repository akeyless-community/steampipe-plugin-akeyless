package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
)

type authMethodsFetcher struct {
}

type authMethodsPaginator akeyless_sdk.ListAuthMethodsOutput

func (p authMethodsPaginator) GetPageItems() []akeyless_sdk.AuthMethod {
	tmp := akeyless_sdk.ListAuthMethodsOutput(p)
	return tmp.GetAuthMethods()
}

func (p authMethodsPaginator) NextPageToken() *string {
	if p.AuthMethods == nil || len(*p.AuthMethods) == 0 {
		return nil
	}
	return p.NextPage
}

func (f authMethodsFetcher) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[akeyless_sdk.AuthMethod], *http.Response, error) {
	out, httpResp, err := apiService.ListAuthMethods(ctx).Body(akeyless_sdk.ListAuthMethods{Token: token, PaginationToken: paginationToken}).Execute()
	return authMethodsPaginator(out), httpResp, err
}
