package akeyless

import (
	"context"
	"fmt"

	"github.com/akeylesslabs/akeyless-go/v2"
	"github.com/go-errors/errors"
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

	if akeylessConfig.AccessType == nil {
		return nil, errors.New("A valid access_type property is required. (api_key/password/saml/ldap/k8s/azure_ad/oidc/aws_iam/universal_identity/jwt/gcp/cert)")
	}

	akeylessService := BuildAkeylessService()

	authType := loginType(*akeylessConfig.AccessType)

	authBody, err := getAuthInfo(authType, akeylessConfig)
	if err != nil {
		return nil, err
	}

	var apiErr akeyless.GenericOpenAPIError

	authOut, _, err := akeylessService.client.Auth(ctx).Body(*authBody).Execute()
	if err != nil {
		if errors.As(err, &apiErr) {
			return nil, errors.New(fmt.Sprintf("authentication failed: %v", string(apiErr.Body())))
		}
		return nil, errors.New(fmt.Sprintf("authentication failed: %v", err))
	}
	token := authOut.GetToken()

	akeylessService.token = &token

	return akeylessService, nil
}

func getAuthInfo(authType loginType, akeylessConfig akeylessConfig) (*akeyless.Auth, error) {
	authBody := akeyless.NewAuthWithDefaults()

	err := setAuthBody(authBody, authType, akeylessConfig)
	if err != nil {
		return nil, err
	}

	return authBody, nil
}

func setAuthBody(authBody *akeyless.Auth, authType loginType, akeylessConfig akeylessConfig) error {
	switch authType {
	case ApiKeyLogin:
		authBody.AccessId = akeylessConfig.AccessId
		authBody.AccessKey = akeylessConfig.AccessKey
		authBody.AccessType = akeylessConfig.AccessType
		return nil
	// case PasswordLogin:
	default:
		return fmt.Errorf("please choose supported auth type for login method: api_key/password/saml/ldap/k8s/azure_ad/oidc/aws_iam/universal_identity/jwt/gcp/cert")
	}
}
