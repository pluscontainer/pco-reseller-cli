package psos

import (
	"context"
	"fmt"

	"github.com/pluscontainer/pco-reseller-cli/pkg/openapi"
)

func (client PsOpenstackClient) GetImage(ctx context.Context, imageID string) (*openapi.ImageResponse, error) {
	resp, err := client.openapiClient.GetImageApiV1ImageImageIdGetWithResponse(ctx, imageID, client.authorizedRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 404 {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("invalid response while retrieving image: %s", string(resp.Body))
	}

	return resp.JSON200.Data, nil
}

func (client PsOpenstackClient) UpdateImageVisibility(ctx context.Context, imageID string, visibility openapi.ImageVisibility) error {
	if _, err := client.GetImage(ctx, imageID); err != nil {
		return err
	}

	params := openapi.UpdateImageVisibility{
		Visibility: visibility,
	}

	resp, err := client.openapiClient.UpdateImageVisibilityApiV1ImageImageIdVisibilityPutWithResponse(ctx, imageID, params, client.authorizedRequest)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("invalid response while updating image visibility: %s", string(resp.Body))
	}

	return nil
}
