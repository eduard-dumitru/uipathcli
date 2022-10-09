package executor

import (
	"net/url"

	"github.com/UiPath/uipathcli/config"
)

type ExecutionContext struct {
	Method           string
	BaseUri          url.URL
	Route            string
	PathParameters   []ExecutionParameter
	QueryParameters  []ExecutionParameter
	HeaderParameters []ExecutionParameter
	BodyParameters   []ExecutionParameter
	FormParameters   []ExecutionParameter
	AuthConfig       config.AuthConfig
	Insecure         bool
	Debug            bool
}

func NewExecutionContext(
	method string,
	uri url.URL,
	route string,
	pathParameters []ExecutionParameter,
	queryParameters []ExecutionParameter,
	headerParameters []ExecutionParameter,
	bodyParameters []ExecutionParameter,
	formParameters []ExecutionParameter,
	authConfig config.AuthConfig,
	insecure bool,
	debug bool) *ExecutionContext {
	return &ExecutionContext{method, uri, route, pathParameters, queryParameters, headerParameters, bodyParameters, formParameters, authConfig, insecure, debug}
}
