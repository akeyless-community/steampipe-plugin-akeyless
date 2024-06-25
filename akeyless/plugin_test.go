package akeyless

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/turbot/steampipe-plugin-sdk/v5/logging"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/context_key"
)

// MockApiDataFetcher is a mock implementation of the apiDataFetcher interface.
type MockApiDataFetcher[T any] struct {
	mock.Mock
}

// FetchData is the mock method that implements the apiDataFetcher interface.
func (m *MockApiDataFetcher[T]) FetchData(ctx context.Context, apiService *akeyless_sdk.V2ApiService, token, paginationToken *string) (dataPaginator[T], *http.Response, error) {
	args := m.Called(ctx, apiService, token, paginationToken)
	return args.Get(0).(dataPaginator[T]), args.Get(1).(*http.Response), args.Error(2)
}

// MockApiClientProvider is a mock implementation of the apiClientProvider interface.
type MockApiClientProvider struct {
	mock.Mock
}

// GetApiService is the mock method that implements the apiClientProvider interface.
func (m *MockApiClientProvider) GetApiService(ctx context.Context, d *plugin.QueryData) (*akeyless_sdk.V2ApiService, *string, error) {
	args := m.Called(ctx, d)
	return args.Get(0).(*akeyless_sdk.V2ApiService), args.Get(1).(*string), args.Error(2)
}

func TestPluginTables(t *testing.T) {

	testCases := []struct {
		name       string
		tblFunc    func() *plugin.Table
		dataSize   int
		respStatus int
		fetchErr   error
		apiErr     error
		isSuccess  bool
		act        func(int, int, error, error, *plugin.QueryData) (interface{}, error)
	}{
		{
			name:       "table_role_success",
			tblFunc:    tableRoles,
			respStatus: 200,
			dataSize:   100,
			isSuccess:  true,
			act:        setupRolesMock,
		},
		{
			name:       "table_role_fail_api_status",
			tblFunc:    tableRoles,
			respStatus: 400,
			isSuccess:  false,
			act:        setupRolesMock,
		},
		{
			name:       "table_role_fail_fetch",
			tblFunc:    tableRoles,
			respStatus: 200,
			isSuccess:  false,
			fetchErr:   fmt.Errorf("fetch err"),
			act:        setupRolesMock,
		},
		{
			name:    "table_role_fail_api",
			tblFunc: tableRoles,
			apiErr:  fmt.Errorf("api err"),
			act:     setupRolesMock,
		},
		{
			name:       "table_item_success",
			tblFunc:    tableItems,
			respStatus: 200,
			dataSize:   100,
			isSuccess:  true,
			act:        setupItemsMock,
		},
		{
			name:       "table_item_fail_api_status",
			tblFunc:    tableItems,
			dataSize:   15,
			respStatus: 400,
			isSuccess:  false,
			act:        setupItemsMock,
		},
		{
			name:       "table_item_fail_fetch",
			tblFunc:    tableItems,
			dataSize:   15,
			respStatus: 200,
			isSuccess:  false,
			fetchErr:   fmt.Errorf("fetch err"),
			act:        setupItemsMock,
		},
		{
			name:    "table_item_fail_api",
			tblFunc: tableItems,
			apiErr:  fmt.Errorf("api err"),
			act:     setupItemsMock,
		},
		{
			name:       "table_authmethod_success",
			tblFunc:    tableAuthMethods,
			respStatus: 200,
			dataSize:   100,
			isSuccess:  true,
			act:        setupAuthMethodsMock,
		},
		{
			name:       "table_authmethod_fail_api_status",
			tblFunc:    tableAuthMethods,
			dataSize:   15,
			respStatus: 400,
			isSuccess:  false,
			act:        setupAuthMethodsMock,
		},
		{
			name:       "table_authmethod_fail_fetch",
			tblFunc:    tableAuthMethods,
			dataSize:   15,
			respStatus: 200,
			isSuccess:  false,
			fetchErr:   fmt.Errorf("fetch err"),
			act:        setupAuthMethodsMock,
		},
		{
			name:    "table_authmethod_fail_api",
			tblFunc: tableAuthMethods,
			apiErr:  fmt.Errorf("api err"),
			act:     setupAuthMethodsMock,
		},
		{
			name:       "table_gateway_success",
			tblFunc:    tableGateways,
			respStatus: 200,
			dataSize:   100,
			isSuccess:  true,
			act:        setupGatewaysMock,
		},
		{
			name:       "table_gateway_fail_api_status",
			tblFunc:    tableGateways,
			dataSize:   15,
			respStatus: 400,
			isSuccess:  false,
			act:        setupGatewaysMock,
		},
		{
			name:       "table_gateway_fail_fetch",
			tblFunc:    tableGateways,
			dataSize:   15,
			respStatus: 200,
			isSuccess:  false,
			fetchErr:   fmt.Errorf("fetch err"),
			act:        setupGatewaysMock,
		},
		{
			name:    "table_gateway_fail_api",
			tblFunc: tableGateways,
			apiErr:  fmt.Errorf("api err"),
			act:     setupGatewaysMock,
		},
		{
			name:       "table_target_success",
			tblFunc:    tableTargets,
			respStatus: 200,
			dataSize:   100,
			isSuccess:  true,
			act:        setupTargetsMock,
		},
		{
			name:       "table_target_fail_api_status",
			tblFunc:    tableTargets,
			respStatus: 400,
			isSuccess:  false,
			act:        setupTargetsMock,
		},
		{
			name:       "table_target_fail_fetch",
			tblFunc:    tableTargets,
			respStatus: 200,
			isSuccess:  false,
			fetchErr:   fmt.Errorf("fetch err"),
			act:        setupTargetsMock,
		},
		{
			name:    "table_target_fail_api",
			tblFunc: tableTargets,
			apiErr:  fmt.Errorf("api err"),
			act:     setupTargetsMock,
		}}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			tC := tC
			t.Parallel()

			//Arrange
			streamCallCount := 0

			qd := &plugin.QueryData{
				Connection: &plugin.Connection{},
				Table:      tC.tblFunc(),
				StreamListItem: func(ctx context.Context, i ...interface{}) {
					streamCallCount++
				}}

			// Act
			dt, err := tC.act(tC.dataSize, tC.respStatus, tC.fetchErr, tC.apiErr, qd)

			//Assert
			if tC.isSuccess {
				require.NoError(t, err, "hydrate_error")
				require.Empty(t, dt, "empty_data")
				require.Equal(t, streamCallCount, tC.dataSize, "data count mismatch")
			} else {
				require.Error(t, err, "no error while expect")
				require.Empty(t, dt, "non empty data on error")
				require.Equal(t, streamCallCount, 0, "non zero data on error")
			}

		})
	}

}

func setupRolesMock(dataSize int, respStatus int, fetchErr, apiErr error, qd *plugin.QueryData) (interface{}, error) {
	data := make([]akeyless_sdk.Role, dataSize)
	return runHydrateListFetch(rolesPaginator{Roles: &data}, qd, respStatus, fetchErr, apiErr)
}

func setupItemsMock(dataSize int, respStatus int, fetchErr, apiErr error, qd *plugin.QueryData) (interface{}, error) {
	data := make([]akeyless_sdk.Item, dataSize)
	return runHydrateListFetch(itemsPaginator{Items: &data}, qd, respStatus, fetchErr, apiErr)

}

func setupGatewaysMock(dataSize int, respStatus int, fetchErr, apiErr error, qd *plugin.QueryData) (interface{}, error) {
	data := make([]akeyless_sdk.GwClusterIdentity, dataSize)
	return runHydrateListFetch(gatewayPaginator{Clusters: &data}, qd, respStatus, fetchErr, apiErr)
}

func setupAuthMethodsMock(dataSize int, respStatus int, fetchErr, apiErr error, qd *plugin.QueryData) (interface{}, error) {
	data := make([]akeyless_sdk.AuthMethod, dataSize)
	return runHydrateListFetch(authMethodsPaginator{AuthMethods: &data}, qd, respStatus, fetchErr, apiErr)
}

func setupTargetsMock(dataSize int, respStatus int, fetchErr, apiErr error, qd *plugin.QueryData) (interface{}, error) {
	data := make([]akeyless_sdk.Target, dataSize)
	return runHydrateListFetch(targetPaginator{Targets: &data}, qd, respStatus, fetchErr, apiErr)
}

func runHydrateListFetch[T any](paginator dataPaginator[T], qd *plugin.QueryData, respStatus int, fetchErr, apiErr error) (interface{}, error) {

	p := setupMockApiClientProvider(apiErr)
	ctx := setupContext()
	hd := &plugin.HydrateData{}

	mockFetcher := &MockApiDataFetcher[T]{}
	mockFetcher.On("FetchData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(paginator, &http.Response{StatusCode: respStatus}, fetchErr)

	// Act
	return hydrateListFetchFunc(p, mockFetcher)(ctx, qd, hd)
}

func setupContext() context.Context {
	loggOpts := &hclog.LoggerOptions{Name: "plugin", Output: io.Discard}
	logger := logging.NewLogger(loggOpts)

	bg := context.Background()
	ctx := context.WithValue(bg, context_key.Logger, logger)
	return ctx
}

func setupMockApiClientProvider(err error) apiClientProvider {
	fakeToken := "fakeToken"
	mockApiClientProvider := &MockApiClientProvider{}
	mockApiClientProvider.On("GetApiService", mock.Anything, mock.Anything).Return(&akeyless_sdk.V2ApiService{}, &fakeToken, err)
	return mockApiClientProvider
}
