package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
)

type rolesFetcher struct {
}

type rolesPaginator akeyless_sdk.ListRolesOutput

func (f rolesPaginator) GetPageItems() []akeyless_sdk.Role {
	tmp := akeyless_sdk.ListRolesOutput(f)
	return tmp.GetRoles()
}
func (f rolesPaginator) NextPageToken() *string {
	if f.Roles == nil || len(*f.Roles) == 0 {
		return nil
	}
	return f.NextPage
}

func (f rolesFetcher) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[akeyless_sdk.Role], *http.Response, error) {
	out, httpResp, err := apiService.ListRoles(ctx).Body(akeyless_sdk.ListRoles{Token: token, PaginationToken: paginationToken}).Execute()
	return rolesPaginator(out), httpResp, err
}
