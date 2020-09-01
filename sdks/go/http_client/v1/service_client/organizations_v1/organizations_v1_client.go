// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new organizations v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateOrganization(params *CreateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationOK, *CreateOrganizationNoContent, error)

	CreateOrganizationMember(params *CreateOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationMemberOK, *CreateOrganizationMemberNoContent, error)

	DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationOK, *DeleteOrganizationNoContent, error)

	DeleteOrganizationInvitation(params *DeleteOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationInvitationOK, *DeleteOrganizationInvitationNoContent, error)

	DeleteOrganizationMember(params *DeleteOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationMemberOK, *DeleteOrganizationMemberNoContent, error)

	GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationOK, *GetOrganizationNoContent, error)

	GetOrganizationInvitation(params *GetOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationInvitationOK, *GetOrganizationInvitationNoContent, error)

	GetOrganizationMember(params *GetOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMemberOK, *GetOrganizationMemberNoContent, error)

	GetOrganizationSettings(params *GetOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSettingsOK, *GetOrganizationSettingsNoContent, error)

	ListOrganizationMembers(params *ListOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationMembersOK, *ListOrganizationMembersNoContent, error)

	ListOrganizationNames(params *ListOrganizationNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationNamesOK, *ListOrganizationNamesNoContent, error)

	ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationsOK, *ListOrganizationsNoContent, error)

	PatchOrganization(params *PatchOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationOK, *PatchOrganizationNoContent, error)

	PatchOrganizationInvitation(params *PatchOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationInvitationOK, *PatchOrganizationInvitationNoContent, error)

	PatchOrganizationMember(params *PatchOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationMemberOK, *PatchOrganizationMemberNoContent, error)

	PatchOrganizationSettings(params *PatchOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationSettingsOK, *PatchOrganizationSettingsNoContent, error)

	UpdateOrganization(params *UpdateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationOK, *UpdateOrganizationNoContent, error)

	UpdateOrganizationInvitation(params *UpdateOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationInvitationOK, *UpdateOrganizationInvitationNoContent, error)

	UpdateOrganizationMember(params *UpdateOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationMemberOK, *UpdateOrganizationMemberNoContent, error)

	UpdateOrganizationSettings(params *UpdateOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationSettingsOK, *UpdateOrganizationSettingsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateOrganization creates organization
*/
func (a *Client) CreateOrganization(params *CreateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationOK, *CreateOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateOrganization",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateOrganizationOK:
		return value, nil, nil
	case *CreateOrganizationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOrganizationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateOrganizationMember creates organization member
*/
func (a *Client) CreateOrganizationMember(params *CreateOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationMemberOK, *CreateOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateOrganizationMember",
		Method:             "POST",
		PathPattern:        "/api/v1/orgs/{owner}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateOrganizationMemberOK:
		return value, nil, nil
	case *CreateOrganizationMemberNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateOrganizationMemberDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteOrganization deletes organization
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationOK, *DeleteOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteOrganizationOK:
		return value, nil, nil
	case *DeleteOrganizationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOrganizationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteOrganizationInvitation deletes organization invitation details
*/
func (a *Client) DeleteOrganizationInvitation(params *DeleteOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationInvitationOK, *DeleteOrganizationInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOrganizationInvitation",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteOrganizationInvitationOK:
		return value, nil, nil
	case *DeleteOrganizationInvitationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOrganizationInvitationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteOrganizationMember deletes organization member details
*/
func (a *Client) DeleteOrganizationMember(params *DeleteOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationMemberOK, *DeleteOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOrganizationMember",
		Method:             "DELETE",
		PathPattern:        "/api/v1/orgs/{owner}/members/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteOrganizationMemberOK:
		return value, nil, nil
	case *DeleteOrganizationMemberNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteOrganizationMemberDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOrganization gets organization
*/
func (a *Client) GetOrganization(params *GetOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationOK, *GetOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganization",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationOK:
		return value, nil, nil
	case *GetOrganizationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrganizationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOrganizationInvitation gets organization invitation details
*/
func (a *Client) GetOrganizationInvitation(params *GetOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationInvitationOK, *GetOrganizationInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizationInvitation",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationInvitationOK:
		return value, nil, nil
	case *GetOrganizationInvitationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrganizationInvitationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOrganizationMember gets organization member details
*/
func (a *Client) GetOrganizationMember(params *GetOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationMemberOK, *GetOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizationMember",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/members/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationMemberOK:
		return value, nil, nil
	case *GetOrganizationMemberNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrganizationMemberDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetOrganizationSettings gets organization settings
*/
func (a *Client) GetOrganizationSettings(params *GetOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetOrganizationSettingsOK, *GetOrganizationSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOrganizationSettings",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetOrganizationSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetOrganizationSettingsOK:
		return value, nil, nil
	case *GetOrganizationSettingsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetOrganizationSettingsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListOrganizationMembers gets organization members
*/
func (a *Client) ListOrganizationMembers(params *ListOrganizationMembersParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationMembersOK, *ListOrganizationMembersNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOrganizationMembers",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/{owner}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOrganizationMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListOrganizationMembersOK:
		return value, nil, nil
	case *ListOrganizationMembersNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOrganizationMembersDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListOrganizationNames lists organizations names
*/
func (a *Client) ListOrganizationNames(params *ListOrganizationNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationNamesOK, *ListOrganizationNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOrganizationNames",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOrganizationNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListOrganizationNamesOK:
		return value, nil, nil
	case *ListOrganizationNamesNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOrganizationNamesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListOrganizations lists organizations
*/
func (a *Client) ListOrganizations(params *ListOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrganizationsOK, *ListOrganizationsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListOrganizations",
		Method:             "GET",
		PathPattern:        "/api/v1/orgs/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListOrganizationsOK:
		return value, nil, nil
	case *ListOrganizationsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListOrganizationsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchOrganization patches organization
*/
func (a *Client) PatchOrganization(params *PatchOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationOK, *PatchOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchOrganization",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchOrganizationOK:
		return value, nil, nil
	case *PatchOrganizationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchOrganizationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchOrganizationInvitation patches organization invitation
*/
func (a *Client) PatchOrganizationInvitation(params *PatchOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationInvitationOK, *PatchOrganizationInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrganizationInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchOrganizationInvitation",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchOrganizationInvitationOK:
		return value, nil, nil
	case *PatchOrganizationInvitationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchOrganizationInvitationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchOrganizationMember patches organization member
*/
func (a *Client) PatchOrganizationMember(params *PatchOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationMemberOK, *PatchOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchOrganizationMember",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/members/{member.user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchOrganizationMemberOK:
		return value, nil, nil
	case *PatchOrganizationMemberNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchOrganizationMemberDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchOrganizationSettings patches oranization settings
*/
func (a *Client) PatchOrganizationSettings(params *PatchOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*PatchOrganizationSettingsOK, *PatchOrganizationSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrganizationSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchOrganizationSettings",
		Method:             "PATCH",
		PathPattern:        "/api/v1/orgs/{owner}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchOrganizationSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PatchOrganizationSettingsOK:
		return value, nil, nil
	case *PatchOrganizationSettingsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchOrganizationSettingsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOrganization updates organization
*/
func (a *Client) UpdateOrganization(params *UpdateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationOK, *UpdateOrganizationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateOrganization",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationOK:
		return value, nil, nil
	case *UpdateOrganizationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrganizationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOrganizationInvitation updates organization invitation
*/
func (a *Client) UpdateOrganizationInvitation(params *UpdateOrganizationInvitationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationInvitationOK, *UpdateOrganizationInvitationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationInvitationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateOrganizationInvitation",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/invitations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrganizationInvitationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationInvitationOK:
		return value, nil, nil
	case *UpdateOrganizationInvitationNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrganizationInvitationDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOrganizationMember updates organization member
*/
func (a *Client) UpdateOrganizationMember(params *UpdateOrganizationMemberParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationMemberOK, *UpdateOrganizationMemberNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationMemberParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateOrganizationMember",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/members/{member.user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrganizationMemberReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationMemberOK:
		return value, nil, nil
	case *UpdateOrganizationMemberNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrganizationMemberDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateOrganizationSettings updates organization settings
*/
func (a *Client) UpdateOrganizationSettings(params *UpdateOrganizationSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationSettingsOK, *UpdateOrganizationSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateOrganizationSettings",
		Method:             "PUT",
		PathPattern:        "/api/v1/orgs/{owner}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateOrganizationSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationSettingsOK:
		return value, nil, nil
	case *UpdateOrganizationSettingsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateOrganizationSettingsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
