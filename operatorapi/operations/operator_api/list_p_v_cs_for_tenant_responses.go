// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package operator_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// ListPVCsForTenantOKCode is the HTTP code returned for type ListPVCsForTenantOK
const ListPVCsForTenantOKCode int = 200

/*
ListPVCsForTenantOK A successful response.

swagger:response listPVCsForTenantOK
*/
type ListPVCsForTenantOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListPVCsResponse `json:"body,omitempty"`
}

// NewListPVCsForTenantOK creates ListPVCsForTenantOK with default headers values
func NewListPVCsForTenantOK() *ListPVCsForTenantOK {

	return &ListPVCsForTenantOK{}
}

// WithPayload adds the payload to the list p v cs for tenant o k response
func (o *ListPVCsForTenantOK) WithPayload(payload *models.ListPVCsResponse) *ListPVCsForTenantOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list p v cs for tenant o k response
func (o *ListPVCsForTenantOK) SetPayload(payload *models.ListPVCsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPVCsForTenantOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
ListPVCsForTenantDefault Generic error response.

swagger:response listPVCsForTenantDefault
*/
type ListPVCsForTenantDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListPVCsForTenantDefault creates ListPVCsForTenantDefault with default headers values
func NewListPVCsForTenantDefault(code int) *ListPVCsForTenantDefault {
	if code <= 0 {
		code = 500
	}

	return &ListPVCsForTenantDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the list p v cs for tenant default response
func (o *ListPVCsForTenantDefault) WithStatusCode(code int) *ListPVCsForTenantDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the list p v cs for tenant default response
func (o *ListPVCsForTenantDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the list p v cs for tenant default response
func (o *ListPVCsForTenantDefault) WithPayload(payload *models.Error) *ListPVCsForTenantDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list p v cs for tenant default response
func (o *ListPVCsForTenantDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPVCsForTenantDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
