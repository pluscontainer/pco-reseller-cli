package psos

import "github.com/Wieneo/pco-reseller-cli/v2/pkg/openapi"

type PsOpenstackClient struct {
	openapiClient *openapi.ClientWithResponses
	loginToken    *openapi.Token
}
