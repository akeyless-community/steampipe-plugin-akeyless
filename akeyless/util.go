package akeyless

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/akeyless-community/akeyless-sheller/sheller"
	"github.com/akeylesslabs/akeyless-go/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type loginType string

const (
	ApiKeyLogin   loginType = "api_key"
	PasswordLogin loginType = "password"
	SamlLogin     loginType = "saml"
	LdapLogin     loginType = "ldap"
	K8sLogin      loginType = "k8s"
	AzureAdLogin  loginType = "azure_ad"
	OidcLogin     loginType = "oidc"
	AwsIamLogin   loginType = "aws_iam"
	UidLogin      loginType = "universal_identity"
	CertLogin     loginType = "cert"
)

type AkeylessService struct {
	client *akeyless.V2ApiService
	token  *string
}

func newAkeylessService(client *akeyless.V2ApiService) *AkeylessService {
	return &AkeylessService{
		client: client,
	}
}

func BuildAkeylessService(url ...string) *AkeylessService {

	urlString := "https://api.akeyless.io"
	if url != nil {
		urlString = url[0] + "/api/v2"
	}
	client := akeyless.NewAPIClient(&akeyless.Configuration{
		Servers: []akeyless.ServerConfiguration{
			{
				URL: urlString,
			},
		},
	}).V2Api

	return newAkeylessService(client)
}

func connect(ctx context.Context, d *plugin.QueryData) (*AkeylessService, error) {
	akeylessConfig := GetConfig(d.Connection)

	// Define the configuration
	config := sheller.NewConfigWithDefaults()

	// set any of the set properties from the akeylessConfig struct over any defaults
	if akeylessConfig.CLIPath != nil && *akeylessConfig.CLIPath != "" {
		config.CLIPath = *akeylessConfig.CLIPath
	}
	if akeylessConfig.Profile != nil && *akeylessConfig.Profile != "" {
		config.Profile = *akeylessConfig.Profile
	}
	if akeylessConfig.AkeylessPath != nil && *akeylessConfig.AkeylessPath != "" {
		config.AkeylessPath = *akeylessConfig.AkeylessPath
	}
	if akeylessConfig.ExpiryBuffer != nil && *akeylessConfig.ExpiryBuffer != "" {
		expiryBufferString := string(*akeylessConfig.ExpiryBuffer)
		expiryBuffer, err := time.ParseDuration(expiryBufferString)
		if err == nil {
			config.ExpiryBuffer = expiryBuffer
		}
	}
	if akeylessConfig.Debug != nil && *akeylessConfig.Debug != "" {
		debugString := string(*akeylessConfig.Debug)
		debug, err := strconv.ParseBool(debugString)
		if err == nil {
			config.Debug = debug
		}
	}

	token, err := sheller.InitializeAndGetToken(config)
	if err != nil {
		fmt.Printf("Failed to initialize and get token: %v\n", err)
	}

	plugin.Logger(ctx).Trace("connect", "token", &token.Token)

	akeylessService := BuildAkeylessService()

	akeylessService.token = &token.Token

	return akeylessService, nil
}
