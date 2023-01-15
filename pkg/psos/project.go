package psos

import (
	"context"
	"fmt"

	"github.com/plusserver/pluscloudopen-reseller-cli/v2/pkg/openapi"
)

func (client PsOpenstackClient) GetProjects(ctx context.Context) (*[]openapi.ProjectCreatedResponse, error) {
	resp, err := client.openapiClient.ListProjectsApiV1ProjectGetWithResponse(ctx, client.authorizedRequest)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) GetProject(ctx context.Context, id string) (*openapi.ProjectCreatedResponse, error) {
	resp, err := client.openapiClient.GetProjectApiV1ProjectProjectIdGetWithResponse(ctx, id, client.authorizedRequest)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 404 {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) UpdateProject(ctx context.Context, id string, params openapi.UpdateProjectApiV1ProjectProjectIdPutJSONRequestBody) (*openapi.ProjectCreatedResponse, error) {
	//Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, id); err != nil {
		return nil, err
	}

	resp, err := client.openapiClient.UpdateProjectApiV1ProjectProjectIdPutWithResponse(ctx, id, params, client.authorizedRequest)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) DeleteProject(ctx context.Context, id string) error {
	//Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, id); err != nil {
		return err
	}

	resp, err := client.openapiClient.DeleteProjectApiV1ProjectProjectIdDeleteWithResponse(ctx, id, client.authorizedRequest)

	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 {
		return fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return nil
}

func (client PsOpenstackClient) CreateProject(ctx context.Context, params openapi.CreateProjectApiV1ProjectPostJSONRequestBody) (*openapi.ProjectCreatedResponse, error) {
	resp, err := client.openapiClient.CreateProjectApiV1ProjectPostWithResponse(ctx, params, client.authorizedRequest)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 201 {
		return nil, fmt.Errorf("invalid response while creating project: %s", string(resp.Body))
	}

	return resp.JSON201.Data, nil
}

func (client PsOpenstackClient) GetUsersInProject(ctx context.Context, projectId string) (*[]openapi.ProjectUser, error) {
	//Don't need safety check here -> API returns 404 if project is missing

	resp, err := client.openapiClient.GetUserForProjectApiV1ProjectProjectIdUserGetWithResponse(ctx, projectId, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 404 {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) AddUserToProject(ctx context.Context, projectId, userID string) error {
	//Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, projectId); err != nil {
		return err
	}

	//Safety check (does the specified user exist and can be retrieved)
	if _, err := client.GetUser(ctx, userID); err != nil {
		return err
	}

	resp, err := client.openapiClient.AddUserToProjectApiV1ProjectProjectIdUserUserIdPostWithResponse(ctx, projectId, userID, client.authorizedRequest)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return nil
}

func (client PsOpenstackClient) RemoveUserFromProject(ctx context.Context, projectId, userID string) error {
	//Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, projectId); err != nil {
		return err
	}

	//Safety check (does the specified user exist and can be retrieved)
	if _, err := client.GetUser(ctx, userID); err != nil {
		return err
	}

	resp, err := client.openapiClient.RemoveUserFromProjectApiV1ProjectProjectIdUserUserIdDeleteWithResponse(ctx, projectId, userID, client.authorizedRequest)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 {
		return fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return nil
}

func (client PsOpenstackClient) GetProjectQuota(ctx context.Context, projectId string) (*openapi.UpdateQuota, error) {
	///Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, projectId); err != nil {
		return nil, err
	}

	resp, err := client.openapiClient.GetQuotaApiV1ProjectProjectIdQuotaGetWithResponse(ctx, projectId, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) UpdateProjectQuota(ctx context.Context, projectId string, params openapi.UpdateQuota) (*openapi.UpdateQuota, error) {
	///Safety check (does the specified project exist and can be retrieved)
	if _, err := client.GetProject(ctx, projectId); err != nil {
		return nil, err
	}

	resp, err := client.openapiClient.UpdateQuotaApiV1ProjectProjectIdQuotaPutWithResponse(ctx, projectId, params, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("invalid response while retrieving projects: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}
