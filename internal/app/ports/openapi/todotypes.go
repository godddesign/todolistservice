// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package openapi

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreateList defines model for CreateList.
type CreateList struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	UserUUID    string `json:"userUUID"`
}

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// DispatchJSONBody defines parameters for Dispatch.
type DispatchJSONBody interface{}

// DispatchJSONRequestBody defines body for Dispatch for application/json ContentType.
type DispatchJSONRequestBody DispatchJSONBody
