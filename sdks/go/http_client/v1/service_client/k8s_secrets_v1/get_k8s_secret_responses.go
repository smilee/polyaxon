// Copyright 2019 Polyaxon, Inc.
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

package k8s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetK8sSecretReader is a Reader for the GetK8sSecret structure.
type GetK8sSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetK8sSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetK8sSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetK8sSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetK8sSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetK8sSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetK8sSecretOK creates a GetK8sSecretOK with default headers values
func NewGetK8sSecretOK() *GetK8sSecretOK {
	return &GetK8sSecretOK{}
}

/*GetK8sSecretOK handles this case with default header values.

A successful response.
*/
type GetK8sSecretOK struct {
	Payload *service_model.V1K8sResource
}

func (o *GetK8sSecretOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/{uuid}][%d] getK8sSecretOK  %+v", 200, o.Payload)
}

func (o *GetK8sSecretOK) GetPayload() *service_model.V1K8sResource {
	return o.Payload
}

func (o *GetK8sSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8sResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8sSecretNoContent creates a GetK8sSecretNoContent with default headers values
func NewGetK8sSecretNoContent() *GetK8sSecretNoContent {
	return &GetK8sSecretNoContent{}
}

/*GetK8sSecretNoContent handles this case with default header values.

No content.
*/
type GetK8sSecretNoContent struct {
	Payload interface{}
}

func (o *GetK8sSecretNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/{uuid}][%d] getK8sSecretNoContent  %+v", 204, o.Payload)
}

func (o *GetK8sSecretNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetK8sSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8sSecretForbidden creates a GetK8sSecretForbidden with default headers values
func NewGetK8sSecretForbidden() *GetK8sSecretForbidden {
	return &GetK8sSecretForbidden{}
}

/*GetK8sSecretForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetK8sSecretForbidden struct {
	Payload interface{}
}

func (o *GetK8sSecretForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/{uuid}][%d] getK8sSecretForbidden  %+v", 403, o.Payload)
}

func (o *GetK8sSecretForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetK8sSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8sSecretNotFound creates a GetK8sSecretNotFound with default headers values
func NewGetK8sSecretNotFound() *GetK8sSecretNotFound {
	return &GetK8sSecretNotFound{}
}

/*GetK8sSecretNotFound handles this case with default header values.

Resource does not exist.
*/
type GetK8sSecretNotFound struct {
	Payload interface{}
}

func (o *GetK8sSecretNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets/{uuid}][%d] getK8sSecretNotFound  %+v", 404, o.Payload)
}

func (o *GetK8sSecretNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetK8sSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}