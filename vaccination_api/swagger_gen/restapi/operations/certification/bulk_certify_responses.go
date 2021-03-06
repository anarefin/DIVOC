// Code generated by go-swagger; DO NOT EDIT.

package certification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// BulkCertifyOKCode is the HTTP code returned for type BulkCertifyOK
const BulkCertifyOKCode int = 200

/*BulkCertifyOK OK

swagger:response bulkCertifyOK
*/
type BulkCertifyOK struct {
}

// NewBulkCertifyOK creates BulkCertifyOK with default headers values
func NewBulkCertifyOK() *BulkCertifyOK {

	return &BulkCertifyOK{}
}

// WriteResponse to the client
func (o *BulkCertifyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// BulkCertifyBadRequestCode is the HTTP code returned for type BulkCertifyBadRequest
const BulkCertifyBadRequestCode int = 400

/*BulkCertifyBadRequest Invalid input

swagger:response bulkCertifyBadRequest
*/
type BulkCertifyBadRequest struct {
}

// NewBulkCertifyBadRequest creates BulkCertifyBadRequest with default headers values
func NewBulkCertifyBadRequest() *BulkCertifyBadRequest {

	return &BulkCertifyBadRequest{}
}

// WriteResponse to the client
func (o *BulkCertifyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// BulkCertifyUnauthorizedCode is the HTTP code returned for type BulkCertifyUnauthorized
const BulkCertifyUnauthorizedCode int = 401

/*BulkCertifyUnauthorized Unauthorized

swagger:response bulkCertifyUnauthorized
*/
type BulkCertifyUnauthorized struct {
}

// NewBulkCertifyUnauthorized creates BulkCertifyUnauthorized with default headers values
func NewBulkCertifyUnauthorized() *BulkCertifyUnauthorized {

	return &BulkCertifyUnauthorized{}
}

// WriteResponse to the client
func (o *BulkCertifyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
