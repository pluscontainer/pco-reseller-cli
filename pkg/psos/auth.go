package psos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"
)

func Login(endpoint, username, password string) (*PsOpenstackClient, error) {
	client, err := openapi.NewClientWithResponses(endpoint)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	data := fmt.Sprintf("grant_type=&username=%s&password=%s", url.QueryEscape(username), url.QueryEscape(password))
	loginResp, err := client.LoginAccessTokenApiV1AuthLoginAccessTokenPostWithBodyWithResponse(ctx, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return nil, err
	}

	if loginResp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving oauth token: %s", string(loginResp.Body))
	}

	return &PsOpenstackClient{
		openapiClient: client,
		loginToken:    loginResp.JSON200,
	}, nil
}

func (client PsOpenstackClient) authorizedRequest(ctx context.Context, req *http.Request) error {
	if client.loginToken == nil {
		return errors.New("api request requested to be authorized, but no token was previously obtained")
	}

	req.Header.Add("Authorization", fmt.Sprintf("%s %s", client.loginToken.TokenType, client.loginToken.AccessToken))
	return nil
}
