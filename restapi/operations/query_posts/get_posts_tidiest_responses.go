// Code generated by go-swagger; DO NOT EDIT.

package query_posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPostsTidiestOKCode is the HTTP code returned for type GetPostsTidiestOK
const GetPostsTidiestOKCode int = 200

/*GetPostsTidiestOK OK

swagger:response getPostsTidiestOK
*/
type GetPostsTidiestOK struct {
}

// NewGetPostsTidiestOK creates GetPostsTidiestOK with default headers values
func NewGetPostsTidiestOK() *GetPostsTidiestOK {

	return &GetPostsTidiestOK{}
}

// WriteResponse to the client
func (o *GetPostsTidiestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}