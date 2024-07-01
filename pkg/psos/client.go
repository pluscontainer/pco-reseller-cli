package psos

import "github.com/pluscontainer/reseller-cli/pkg/openapi"

type PsOpenstackClient struct {
	openapiClient *openapi.ClientWithResponses
	loginToken    *openapi.Token
}
