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

package queues_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchQueueReader is a Reader for the PatchQueue structure.
type PatchQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchQueueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchQueueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchQueueOK creates a PatchQueueOK with default headers values
func NewPatchQueueOK() *PatchQueueOK {
	return &PatchQueueOK{}
}

/*PatchQueueOK handles this case with default header values.

A successful response.
*/
type PatchQueueOK struct {
	Payload *service_model.V1Queue
}

func (o *PatchQueueOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] patchQueueOK  %+v", 200, o.Payload)
}

func (o *PatchQueueOK) GetPayload() *service_model.V1Queue {
	return o.Payload
}

func (o *PatchQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Queue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueNoContent creates a PatchQueueNoContent with default headers values
func NewPatchQueueNoContent() *PatchQueueNoContent {
	return &PatchQueueNoContent{}
}

/*PatchQueueNoContent handles this case with default header values.

No content.
*/
type PatchQueueNoContent struct {
	Payload interface{}
}

func (o *PatchQueueNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] patchQueueNoContent  %+v", 204, o.Payload)
}

func (o *PatchQueueNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchQueueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueForbidden creates a PatchQueueForbidden with default headers values
func NewPatchQueueForbidden() *PatchQueueForbidden {
	return &PatchQueueForbidden{}
}

/*PatchQueueForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchQueueForbidden struct {
	Payload interface{}
}

func (o *PatchQueueForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] patchQueueForbidden  %+v", 403, o.Payload)
}

func (o *PatchQueueForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueNotFound creates a PatchQueueNotFound with default headers values
func NewPatchQueueNotFound() *PatchQueueNotFound {
	return &PatchQueueNotFound{}
}

/*PatchQueueNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchQueueNotFound struct {
	Payload interface{}
}

func (o *PatchQueueNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] patchQueueNotFound  %+v", 404, o.Payload)
}

func (o *PatchQueueNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchQueueDefault creates a PatchQueueDefault with default headers values
func NewPatchQueueDefault(code int) *PatchQueueDefault {
	return &PatchQueueDefault{
		_statusCode: code,
	}
}

/*PatchQueueDefault handles this case with default header values.

An unexpected error response
*/
type PatchQueueDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the patch queue default response
func (o *PatchQueueDefault) Code() int {
	return o._statusCode
}

func (o *PatchQueueDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/orgs/{owner}/agents/{agent}/queues/{queue.uuid}][%d] PatchQueue default  %+v", o._statusCode, o.Payload)
}

func (o *PatchQueueDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PatchQueueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
