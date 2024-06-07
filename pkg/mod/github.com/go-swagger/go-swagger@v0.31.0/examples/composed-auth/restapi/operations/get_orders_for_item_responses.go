// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/composed-auth/models"
)

// GetOrdersForItemOKCode is the HTTP code returned for type GetOrdersForItemOK
const GetOrdersForItemOKCode int = 200

/*
GetOrdersForItemOK multiple orders

swagger:response getOrdersForItemOK
*/
type GetOrdersForItemOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Order `json:"body,omitempty"`
}

// NewGetOrdersForItemOK creates GetOrdersForItemOK with default headers values
func NewGetOrdersForItemOK() *GetOrdersForItemOK {

	return &GetOrdersForItemOK{}
}

// WithPayload adds the payload to the get orders for item o k response
func (o *GetOrdersForItemOK) WithPayload(payload []*models.Order) *GetOrdersForItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get orders for item o k response
func (o *GetOrdersForItemOK) SetPayload(payload []*models.Order) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrdersForItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Order, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetOrdersForItemUnauthorizedCode is the HTTP code returned for type GetOrdersForItemUnauthorized
const GetOrdersForItemUnauthorizedCode int = 401

/*
GetOrdersForItemUnauthorized unauthorized access for a lack of authentication

swagger:response getOrdersForItemUnauthorized
*/
type GetOrdersForItemUnauthorized struct {
}

// NewGetOrdersForItemUnauthorized creates GetOrdersForItemUnauthorized with default headers values
func NewGetOrdersForItemUnauthorized() *GetOrdersForItemUnauthorized {

	return &GetOrdersForItemUnauthorized{}
}

// WriteResponse to the client
func (o *GetOrdersForItemUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetOrdersForItemForbiddenCode is the HTTP code returned for type GetOrdersForItemForbidden
const GetOrdersForItemForbiddenCode int = 403

/*
GetOrdersForItemForbidden forbidden access for a lack of sufficient privileges

swagger:response getOrdersForItemForbidden
*/
type GetOrdersForItemForbidden struct {
}

// NewGetOrdersForItemForbidden creates GetOrdersForItemForbidden with default headers values
func NewGetOrdersForItemForbidden() *GetOrdersForItemForbidden {

	return &GetOrdersForItemForbidden{}
}

// WriteResponse to the client
func (o *GetOrdersForItemForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*
GetOrdersForItemDefault other error response

swagger:response getOrdersForItemDefault
*/
type GetOrdersForItemDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetOrdersForItemDefault creates GetOrdersForItemDefault with default headers values
func NewGetOrdersForItemDefault(code int) *GetOrdersForItemDefault {
	if code <= 0 {
		code = 500
	}

	return &GetOrdersForItemDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get orders for item default response
func (o *GetOrdersForItemDefault) WithStatusCode(code int) *GetOrdersForItemDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get orders for item default response
func (o *GetOrdersForItemDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get orders for item default response
func (o *GetOrdersForItemDefault) WithPayload(payload *models.Error) *GetOrdersForItemDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get orders for item default response
func (o *GetOrdersForItemDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetOrdersForItemDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
