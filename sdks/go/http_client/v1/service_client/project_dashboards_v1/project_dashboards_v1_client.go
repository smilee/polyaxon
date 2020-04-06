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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new project dashboards v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project dashboards v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ProjectDashboardsV1CreateProjectDashboard(params *ProjectDashboardsV1CreateProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1CreateProjectDashboardOK, *ProjectDashboardsV1CreateProjectDashboardNoContent, error)

	ProjectDashboardsV1DeleteProjectDashboard(params *ProjectDashboardsV1DeleteProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1DeleteProjectDashboardOK, *ProjectDashboardsV1DeleteProjectDashboardNoContent, error)

	ProjectDashboardsV1GetProjectDashboard(params *ProjectDashboardsV1GetProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1GetProjectDashboardOK, *ProjectDashboardsV1GetProjectDashboardNoContent, error)

	ProjectDashboardsV1ListProjectDashboardNames(params *ProjectDashboardsV1ListProjectDashboardNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1ListProjectDashboardNamesOK, *ProjectDashboardsV1ListProjectDashboardNamesNoContent, error)

	ProjectDashboardsV1ListProjectDashboards(params *ProjectDashboardsV1ListProjectDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1ListProjectDashboardsOK, *ProjectDashboardsV1ListProjectDashboardsNoContent, error)

	ProjectDashboardsV1PatchProjectDashboard(params *ProjectDashboardsV1PatchProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1PatchProjectDashboardOK, *ProjectDashboardsV1PatchProjectDashboardNoContent, error)

	ProjectDashboardsV1PromoteProjectDashboard(params *ProjectDashboardsV1PromoteProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1PromoteProjectDashboardOK, *ProjectDashboardsV1PromoteProjectDashboardNoContent, error)

	ProjectDashboardsV1UpdateProjectDashboard(params *ProjectDashboardsV1UpdateProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1UpdateProjectDashboardOK, *ProjectDashboardsV1UpdateProjectDashboardNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ProjectDashboardsV1CreateProjectDashboard creates project dashboard
*/
func (a *Client) ProjectDashboardsV1CreateProjectDashboard(params *ProjectDashboardsV1CreateProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1CreateProjectDashboardOK, *ProjectDashboardsV1CreateProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1CreateProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_CreateProjectDashboard",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1CreateProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1CreateProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1CreateProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1CreateProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1DeleteProjectDashboard deletes project dashboard
*/
func (a *Client) ProjectDashboardsV1DeleteProjectDashboard(params *ProjectDashboardsV1DeleteProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1DeleteProjectDashboardOK, *ProjectDashboardsV1DeleteProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1DeleteProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_DeleteProjectDashboard",
		Method:             "DELETE",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1DeleteProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1DeleteProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1DeleteProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1DeleteProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1GetProjectDashboard gets project dashboard
*/
func (a *Client) ProjectDashboardsV1GetProjectDashboard(params *ProjectDashboardsV1GetProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1GetProjectDashboardOK, *ProjectDashboardsV1GetProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1GetProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_GetProjectDashboard",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1GetProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1GetProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1GetProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1GetProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1ListProjectDashboardNames lists project dashboard
*/
func (a *Client) ProjectDashboardsV1ListProjectDashboardNames(params *ProjectDashboardsV1ListProjectDashboardNamesParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1ListProjectDashboardNamesOK, *ProjectDashboardsV1ListProjectDashboardNamesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1ListProjectDashboardNamesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_ListProjectDashboardNames",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/names",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1ListProjectDashboardNamesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1ListProjectDashboardNamesOK:
		return value, nil, nil
	case *ProjectDashboardsV1ListProjectDashboardNamesNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1ListProjectDashboardNamesDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1ListProjectDashboards lists project dashboards
*/
func (a *Client) ProjectDashboardsV1ListProjectDashboards(params *ProjectDashboardsV1ListProjectDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1ListProjectDashboardsOK, *ProjectDashboardsV1ListProjectDashboardsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1ListProjectDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_ListProjectDashboards",
		Method:             "GET",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1ListProjectDashboardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1ListProjectDashboardsOK:
		return value, nil, nil
	case *ProjectDashboardsV1ListProjectDashboardsNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1ListProjectDashboardsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1PatchProjectDashboard patches project dashboard
*/
func (a *Client) ProjectDashboardsV1PatchProjectDashboard(params *ProjectDashboardsV1PatchProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1PatchProjectDashboardOK, *ProjectDashboardsV1PatchProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1PatchProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_PatchProjectDashboard",
		Method:             "PATCH",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1PatchProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1PatchProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1PatchProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1PatchProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1PromoteProjectDashboard promotes project dashboard
*/
func (a *Client) ProjectDashboardsV1PromoteProjectDashboard(params *ProjectDashboardsV1PromoteProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1PromoteProjectDashboardOK, *ProjectDashboardsV1PromoteProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1PromoteProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_PromoteProjectDashboard",
		Method:             "POST",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}/promote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1PromoteProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1PromoteProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1PromoteProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1PromoteProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ProjectDashboardsV1UpdateProjectDashboard updates project dashboard
*/
func (a *Client) ProjectDashboardsV1UpdateProjectDashboard(params *ProjectDashboardsV1UpdateProjectDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectDashboardsV1UpdateProjectDashboardOK, *ProjectDashboardsV1UpdateProjectDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDashboardsV1UpdateProjectDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDashboardsV1_UpdateProjectDashboard",
		Method:             "PUT",
		PathPattern:        "/api/v1/{owner}/{project}/dashboards/{dashboard.uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ProjectDashboardsV1UpdateProjectDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ProjectDashboardsV1UpdateProjectDashboardOK:
		return value, nil, nil
	case *ProjectDashboardsV1UpdateProjectDashboardNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDashboardsV1UpdateProjectDashboardDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}