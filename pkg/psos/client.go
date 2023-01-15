package psos

import "github.com/plusserver/pluscloudopen-reseller-cli/v2/pkg/openapi"

type PsOpenstackClient struct {
	openapiClient *openapi.ClientWithResponses
	loginToken    *openapi.Token
}
