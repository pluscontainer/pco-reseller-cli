package psos

import "github.com/pluscontainer/pco-reseller-cli/pkg/openapi"

type PsOpenstackClient struct {
	openapiClient *openapi.ClientWithResponses
	loginToken    *openapi.Token
}
