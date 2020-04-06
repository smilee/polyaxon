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

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewTeamsV1UpdateTeamMemberParams creates a new TeamsV1UpdateTeamMemberParams object
// with the default values initialized.
func NewTeamsV1UpdateTeamMemberParams() *TeamsV1UpdateTeamMemberParams {
	var ()
	return &TeamsV1UpdateTeamMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsV1UpdateTeamMemberParamsWithTimeout creates a new TeamsV1UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsV1UpdateTeamMemberParamsWithTimeout(timeout time.Duration) *TeamsV1UpdateTeamMemberParams {
	var ()
	return &TeamsV1UpdateTeamMemberParams{

		timeout: timeout,
	}
}

// NewTeamsV1UpdateTeamMemberParamsWithContext creates a new TeamsV1UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsV1UpdateTeamMemberParamsWithContext(ctx context.Context) *TeamsV1UpdateTeamMemberParams {
	var ()
	return &TeamsV1UpdateTeamMemberParams{

		Context: ctx,
	}
}

// NewTeamsV1UpdateTeamMemberParamsWithHTTPClient creates a new TeamsV1UpdateTeamMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsV1UpdateTeamMemberParamsWithHTTPClient(client *http.Client) *TeamsV1UpdateTeamMemberParams {
	var ()
	return &TeamsV1UpdateTeamMemberParams{
		HTTPClient: client,
	}
}

/*TeamsV1UpdateTeamMemberParams contains all the parameters to send to the API endpoint
for the teams v1 update team member operation typically these are written to a http.Request
*/
type TeamsV1UpdateTeamMemberParams struct {

	/*Body
	  Team body

	*/
	Body *service_model.V1TeamMember
	/*MemberUser
	  User

	*/
	MemberUser string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Team
	  Team

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithTimeout(timeout time.Duration) *TeamsV1UpdateTeamMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithContext(ctx context.Context) *TeamsV1UpdateTeamMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithHTTPClient(client *http.Client) *TeamsV1UpdateTeamMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithBody(body *service_model.V1TeamMember) *TeamsV1UpdateTeamMemberParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetBody(body *service_model.V1TeamMember) {
	o.Body = body
}

// WithMemberUser adds the memberUser to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithMemberUser(memberUser string) *TeamsV1UpdateTeamMemberParams {
	o.SetMemberUser(memberUser)
	return o
}

// SetMemberUser adds the memberUser to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetMemberUser(memberUser string) {
	o.MemberUser = memberUser
}

// WithOwner adds the owner to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithOwner(owner string) *TeamsV1UpdateTeamMemberParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithTeam adds the team to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) WithTeam(team string) *TeamsV1UpdateTeamMemberParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the teams v1 update team member params
func (o *TeamsV1UpdateTeamMemberParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsV1UpdateTeamMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param member.user
	if err := r.SetPathParam("member.user", o.MemberUser); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}