// Code generated by go-swagger; DO NOT EDIT.

package query_posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Tech4GoodPH/go-sample-api.git/models"
)

// GetPostsOKCode is the HTTP code returned for type GetPostsOK
const GetPostsOKCode int = 200

/*GetPostsOK OK

swagger:response getPostsOK
*/
type GetPostsOK struct {

	/*
	  In: Body
	*/
	Payload models.Posts `json:"body,omitempty"`
}

// NewGetPostsOK creates GetPostsOK with default headers values
func NewGetPostsOK() *GetPostsOK {

	return &GetPostsOK{}
}

// WithPayload adds the payload to the get posts o k response
func (o *GetPostsOK) WithPayload(payload models.Posts) *GetPostsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get posts o k response
func (o *GetPostsOK) SetPayload(payload models.Posts) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPostsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Posts{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}