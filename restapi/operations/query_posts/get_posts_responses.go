// Code generated by go-swagger; DO NOT EDIT.

package query_posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetPostsOKCode is the HTTP code returned for type GetPostsOK
const GetPostsOKCode int = 200

/*GetPostsOK OK

swagger:response getPostsOK
*/
type GetPostsOK struct {
}

// NewGetPostsOK creates GetPostsOK with default headers values
func NewGetPostsOK() *GetPostsOK {

	return &GetPostsOK{}
}

// WriteResponse to the client
func (o *GetPostsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}