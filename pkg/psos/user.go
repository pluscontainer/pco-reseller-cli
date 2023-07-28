package psos

import (
	"context"
	"fmt"

	"github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"
)

func (client PsOpenstackClient) GetUser(ctx context.Context, userID string) (*openapi.CreatedOpenStackUser, error) {
	resp, err := client.openapiClient.GetUserApiV1UserUserIdGetWithResponse(ctx, userID, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 404 {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("invalid response while retrieving users: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) GetUsers(ctx context.Context) (*[]openapi.CreatedOpenStackUser, error) {
	resp, err := client.openapiClient.ListUsersApiV1UserGetWithResponse(ctx, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving users: %s", string(resp.Body))
	}

	if resp.JSON200.Data == nil {
		resp.JSON200.Data = &[]openapi.CreatedOpenStackUser{}
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) CreateUser(ctx context.Context, params openapi.CreateOpenStackUser) (*openapi.CreatedOpenStackUser, error) {
	resp, err := client.openapiClient.CreateUserApiV1UserPostWithResponse(ctx, params, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("invalid response while retrieving users: %s", string(resp.Body))
	}

	return resp.JSON201.Data, nil
}

func (client PsOpenstackClient) UpdateUser(ctx context.Context, userID string, params openapi.UpdateOpenStackUser) (*openapi.CreatedOpenStackUser, error) {
	//Safety check (does the specified user exist and can be retrieved)
	if _, err := client.GetUser(ctx, userID); err != nil {
		return nil, err
	}

	resp, err := client.openapiClient.UpdateUserApiV1UserUserIdPutWithResponse(ctx, userID, params, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving users: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) DeleteUser(ctx context.Context, userID string) error {
	//Safety check (does the specified user exist and can be retrieved)
	if _, err := client.GetUser(ctx, userID); err != nil {
		return err
	}

	resp, err := client.openapiClient.DeleteUserApiV1UserUserIdDeleteWithResponse(ctx, userID, client.authorizedRequest)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 {
		return fmt.Errorf("invalid response while retrieving users: %s", string(resp.Body))
	}

	return nil
}
