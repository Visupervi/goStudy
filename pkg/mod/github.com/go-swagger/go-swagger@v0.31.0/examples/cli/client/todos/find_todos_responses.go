// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/cli/models"
)

// FindTodosReader is a Reader for the FindTodos structure.
type FindTodosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindTodosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindTodosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindTodosDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindTodosOK creates a FindTodosOK with default headers values
func NewFindTodosOK() *FindTodosOK {
	return &FindTodosOK{}
}

/*
FindTodosOK describes a response with status code 200, with default header values.

list the todo operations
*/
type FindTodosOK struct {
	Payload []*models.Item
}

// IsSuccess returns true when this find todos o k response has a 2xx status code
func (o *FindTodosOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find todos o k response has a 3xx status code
func (o *FindTodosOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find todos o k response has a 4xx status code
func (o *FindTodosOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find todos o k response has a 5xx status code
func (o *FindTodosOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find todos o k response a status code equal to that given
func (o *FindTodosOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find todos o k response
func (o *FindTodosOK) Code() int {
	return 200
}

func (o *FindTodosOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /][%d] findTodosOK %s", 200, payload)
}

func (o *FindTodosOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /][%d] findTodosOK %s", 200, payload)
}

func (o *FindTodosOK) GetPayload() []*models.Item {
	return o.Payload
}

func (o *FindTodosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindTodosDefault creates a FindTodosDefault with default headers values
func NewFindTodosDefault(code int) *FindTodosDefault {
	return &FindTodosDefault{
		_statusCode: code,
	}
}

/*
FindTodosDefault describes a response with status code -1, with default header values.

generic error response
*/
type FindTodosDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this find todos default response has a 2xx status code
func (o *FindTodosDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find todos default response has a 3xx status code
func (o *FindTodosDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find todos default response has a 4xx status code
func (o *FindTodosDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find todos default response has a 5xx status code
func (o *FindTodosDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find todos default response a status code equal to that given
func (o *FindTodosDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the find todos default response
func (o *FindTodosDefault) Code() int {
	return o._statusCode
}

func (o *FindTodosDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /][%d] findTodos default %s", o._statusCode, payload)
}

func (o *FindTodosDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /][%d] findTodos default %s", o._statusCode, payload)
}

func (o *FindTodosDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindTodosDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}