package akeyless

import (
	"context"
	"net/http"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
)

type itemsFetcher struct {
}

type itemsPaginator akeyless_sdk.ListItemsInPathOutput

func (f itemsPaginator) GetPageItems() []akeyless_sdk.Item {
	tmp := akeyless_sdk.ListItemsInPathOutput(f)
	return tmp.GetItems()
}

func (f itemsPaginator) NextPageToken() *string {
	if f.Items == nil || len(*f.Items) == 0 {
		return nil
	}
	return f.NextPage
}

func (f itemsFetcher) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[akeyless_sdk.Item], *http.Response, error) {
	out, resp, err := apiService.ListItems(ctx).Body(akeyless_sdk.ListItems{Token: token, MinimalView: akeyless_sdk.PtrBool(true), PaginationToken: paginationToken}).Execute()
	return itemsPaginator(out), resp, err
}
