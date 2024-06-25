package akeyless

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/akeylesslabs/akeyless-go-cloud-id/cloudprovider/aws"
	"github.com/akeylesslabs/akeyless-go-cloud-id/cloudprovider/azure"
	"github.com/akeylesslabs/akeyless-go-cloud-id/cloudprovider/gcp"
	akeyless_sdk "github.com/akeylesslabs/akeyless-go/v4"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type AccessType string

const (
	AccessTypeApiKey  AccessType = "api_key"
	AccessTypeAwsIAM  AccessType = "aws_iam"
	AccessTypeAzureAd AccessType = "azure_ad"
	AccessTypeGCP     AccessType = "gcp"
	AccessTypeUid     AccessType = "universal_identity"
	K8S               AccessType = "k8s"
	JWT               AccessType = "jwt"
)

func authenticate(ctx context.Context, config akeylessConfig, client *akeyless_sdk.V2ApiService) (*string, error) {
	authParams, err := setupAuthParams(config)
	if err != nil {
		plugin.Logger(ctx).Error("authentication error", err)
		return nil, err
	}
	out, _, err := client.Auth(ctx).Body(*authParams).Execute()

	if err != nil {
		plugin.Logger(ctx).Error("authentication error", err)
		return nil, getAklApiErrMsg(err)
	}
	return akeyless_sdk.PtrString(out.GetToken()), nil
}

func setupAuthParams(authInput akeylessConfig) (*akeyless_sdk.Auth, error) {
	authParams := akeyless_sdk.NewAuth()
	authParams.SetAccessType(*authInput.AccessType)
	authParams.SetAccessId(*authInput.AccessId)

	switch AccessType(*authInput.AccessType) {
	case AccessTypeApiKey:
		authParams.SetAccessKey(*authInput.AccessKey)
	case AccessTypeAwsIAM:
		id, err := aws.GetCloudId()
		if err != nil {
			return nil, fmt.Errorf("failed to get AWS cloud id: %w", err)
		}
		authParams.SetCloudId(id)

	case AccessTypeAzureAd:
		id, err := azure.GetCloudId(*authInput.AzureObjectId)
		if err != nil {
			return nil, fmt.Errorf("failed to get azure cloud id: %w", err)
		}
		if _, err := base64.StdEncoding.DecodeString(id); err != nil {
			id = base64.StdEncoding.EncodeToString([]byte(id))
		}
		authParams.SetCloudId(id)

	case AccessTypeGCP:
		id, err := gcp.GetCloudID(*authInput.GcpAudience)
		if err != nil {
			return nil, fmt.Errorf("failed to get GCP cloud id: %w", err)
		}
		authParams.SetCloudId(id)

	case AccessTypeUid:
		if *authInput.UidToken == "" {
			return nil, fmt.Errorf("UidToken is required for access type %q", AccessTypeUid)
		}
		authParams.SetUidToken(*authInput.UidToken)

	case K8S:
		authParams.SetGatewayUrl(*authInput.ApiUrl)
		authParams.SetK8sServiceAccountToken(*authInput.K8SServiceAccountToken)
		authParams.SetK8sAuthConfigName(*authInput.K8SAuthConfigName)

	case JWT:
		authParams.SetJwt(*authInput.Jwt)

	default:
		return nil, fmt.Errorf("unknown access type: %s", *authInput.AccessType)
	}

	return authParams, nil
}

func getAklApiErrMsg(err error) error {
	msg := "no response body"

	var apiErr akeyless_sdk.GenericOpenAPIError
	if errors.As(err, &apiErr) {
		msg = string(apiErr.Body())
	}

	return fmt.Errorf("can't authenticate with static creds: %s: %w", msg, err)
}
