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

package runs_v1

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
	"github.com/go-openapi/swag"
)

// NewRunsV1GetRunResourcesParams creates a new RunsV1GetRunResourcesParams object
// with the default values initialized.
func NewRunsV1GetRunResourcesParams() *RunsV1GetRunResourcesParams {
	var ()
	return &RunsV1GetRunResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRunsV1GetRunResourcesParamsWithTimeout creates a new RunsV1GetRunResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRunsV1GetRunResourcesParamsWithTimeout(timeout time.Duration) *RunsV1GetRunResourcesParams {
	var ()
	return &RunsV1GetRunResourcesParams{

		timeout: timeout,
	}
}

// NewRunsV1GetRunResourcesParamsWithContext creates a new RunsV1GetRunResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewRunsV1GetRunResourcesParamsWithContext(ctx context.Context) *RunsV1GetRunResourcesParams {
	var ()
	return &RunsV1GetRunResourcesParams{

		Context: ctx,
	}
}

// NewRunsV1GetRunResourcesParamsWithHTTPClient creates a new RunsV1GetRunResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRunsV1GetRunResourcesParamsWithHTTPClient(client *http.Client) *RunsV1GetRunResourcesParams {
	var ()
	return &RunsV1GetRunResourcesParams{
		HTTPClient: client,
	}
}

/*RunsV1GetRunResourcesParams contains all the parameters to send to the API endpoint
for the runs v1 get run resources operation typically these are written to a http.Request
*/
type RunsV1GetRunResourcesParams struct {

	/*Names
	  Names query param.

	*/
	Names *string
	/*Namespace
	  namespace

	*/
	Namespace string
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the run will be assigned

	*/
	Project string
	/*Tail
	  Query param flag to tail the values.

	*/
	Tail *bool
	/*UUID
	  Uuid identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithTimeout(timeout time.Duration) *RunsV1GetRunResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithContext(ctx context.Context) *RunsV1GetRunResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithHTTPClient(client *http.Client) *RunsV1GetRunResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithNames(names *string) *RunsV1GetRunResourcesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetNames(names *string) {
	o.Names = names
}

// WithNamespace adds the namespace to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithNamespace(namespace string) *RunsV1GetRunResourcesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithOwner adds the owner to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithOwner(owner string) *RunsV1GetRunResourcesParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithProject(project string) *RunsV1GetRunResourcesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetProject(project string) {
	o.Project = project
}

// WithTail adds the tail to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithTail(tail *bool) *RunsV1GetRunResourcesParams {
	o.SetTail(tail)
	return o
}

// SetTail adds the tail to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetTail(tail *bool) {
	o.Tail = tail
}

// WithUUID adds the uuid to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) WithUUID(uuid string) *RunsV1GetRunResourcesParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the runs v1 get run resources params
func (o *RunsV1GetRunResourcesParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *RunsV1GetRunResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// query param names
		var qrNames string
		if o.Names != nil {
			qrNames = *o.Names
		}
		qNames := qrNames
		if qNames != "" {
			if err := r.SetQueryParam("names", qNames); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if o.Tail != nil {

		// query param tail
		var qrTail bool
		if o.Tail != nil {
			qrTail = *o.Tail
		}
		qTail := swag.FormatBool(qrTail)
		if qTail != "" {
			if err := r.SetQueryParam("tail", qTail); err != nil {
				return err
			}
		}

	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}