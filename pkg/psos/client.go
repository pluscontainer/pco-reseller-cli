package psos

import "github.com/pluscloudopen/reseller-cli/v2/pkg/openapi"

type PsOpenstackClient struct {
	openapiClient *openapi.ClientWithResponses
	loginToken    *openapi.Token
}
