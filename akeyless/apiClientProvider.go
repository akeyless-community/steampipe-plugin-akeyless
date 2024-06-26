package akeyless

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type apiClientProvider interface {
	GetApiService(ctx context.Context, d *plugin.QueryData) (*akeyless_sdk.V2ApiService, *string, error)
}

type sdkApiClientProvider struct {
}

func (p sdkApiClientProvider) GetApiService(ctx context.Context, d *plugin.QueryData) (*akeyless_sdk.V2ApiService, *string, error) {

	pluginLogger := plugin.Logger(ctx)
	const logPrefix = "utils.Connect"
	config, err := getPluginConfig(ctx, d.Connection)

	if err != nil {
		pluginLogger.Error(logPrefix, "failed to get config", err)
		return nil, nil, fmt.Errorf("failed to get config %w", err)
	}

	httpClient, err := p.setupHttpClient(*config)
	if err != nil {
		pluginLogger.Error(logPrefix, "failed to config http client", err)
		return nil, nil, fmt.Errorf("failed to config http client %w", err)
	}

	client := akeyless_sdk.NewAPIClient(&akeyless_sdk.Configuration{
		Servers: []akeyless_sdk.ServerConfiguration{
			{URL: *config.ApiUrl},
		},
		HTTPClient: httpClient,
		DefaultHeader: map[string]string{
			"akeylessclienttype":    PluginName,
			"akeylessclientversion": PluginVersion},
	}).V2Api

	pluginLogger.Trace(logPrefix+" client successfully created", "akeylessclienttype", PluginName, "akeylessclientversion", PluginVersion)

	token, err := authenticate(ctx, *config, client)

	if err != nil {
		pluginLogger.Error(logPrefix, "authentication failed", err)
		return nil, nil, err
	}

	pluginLogger.Trace(logPrefix, "client successfully created and authenticated")

	return client, token, nil
}

func (p sdkApiClientProvider) setupHttpClient(config akeylessConfig) (*http.Client, error) {

	client := http.Client{}
	//Setup self signed certificates support
	if len(*config.GatewayCaCert) > 0 && config.GatewayCaCert != nil {

		cert, err := os.ReadFile(*config.GatewayCaCert)
		if err != nil {
			return nil, fmt.Errorf("failed to read gateway ca certificate: %w", err)
		}

		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("failed to get system cert pool: %w", err)
		}

		if ok := rootCAs.AppendCertsFromPEM(cert); !ok {
			return nil, fmt.Errorf("failed to append gateway ca certificate")
		}

		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAs,
			},
		}

		client.Transport = transport
	}

	return &client, nil
}
