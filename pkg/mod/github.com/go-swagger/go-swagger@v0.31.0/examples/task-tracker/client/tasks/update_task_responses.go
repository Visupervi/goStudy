// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/task-tracker/models"
)

// UpdateTaskReader is a Reader for the UpdateTask structure.
type UpdateTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewUpdateTaskUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTaskOK creates a UpdateTaskOK with default headers values
func NewUpdateTaskOK() *UpdateTaskOK {
	return &UpdateTaskOK{}
}

/*
UpdateTaskOK describes a response with status code 200, with default header values.

Task details
*/
type UpdateTaskOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this update task o k response has a 2xx status code
func (o *UpdateTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update task o k response has a 3xx status code
func (o *UpdateTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task o k response has a 4xx status code
func (o *UpdateTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update task o k response has a 5xx status code
func (o *UpdateTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update task o k response a status code equal to that given
func (o *UpdateTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update task o k response
func (o *UpdateTaskOK) Code() int {
	return 200
}

func (o *UpdateTaskOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTaskOK %s", 200, payload)
}

func (o *UpdateTaskOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTaskOK %s", 200, payload)
}

func (o *UpdateTaskOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *UpdateTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTaskUnprocessableEntity creates a UpdateTaskUnprocessableEntity with default headers values
func NewUpdateTaskUnprocessableEntity() *UpdateTaskUnprocessableEntity {
	return &UpdateTaskUnprocessableEntity{}
}

/*
UpdateTaskUnprocessableEntity describes a response with status code 422, with default header values.

Validation error
*/
type UpdateTaskUnprocessableEntity struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this update task unprocessable entity response has a 2xx status code
func (o *UpdateTaskUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update task unprocessable entity response has a 3xx status code
func (o *UpdateTaskUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update task unprocessable entity response has a 4xx status code
func (o *UpdateTaskUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this update task unprocessable entity response has a 5xx status code
func (o *UpdateTaskUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this update task unprocessable entity response a status code equal to that given
func (o *UpdateTaskUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the update task unprocessable entity response
func (o *UpdateTaskUnprocessableEntity) Code() int {
	return 422
}

func (o *UpdateTaskUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTaskUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTaskUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTaskUnprocessableEntity %s", 422, payload)
}

func (o *UpdateTaskUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateTaskUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTaskDefault creates a UpdateTaskDefault with default headers values
func NewUpdateTaskDefault(code int) *UpdateTaskDefault {
	return &UpdateTaskDefault{
		_statusCode: code,
	}
}

/*
UpdateTaskDefault describes a response with status code -1, with default header values.

Error response
*/
type UpdateTaskDefault struct {
	_statusCode int
	XErrorCode  string

	Payload *models.Error
}

// IsSuccess returns true when this update task default response has a 2xx status code
func (o *UpdateTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update task default response has a 3xx status code
func (o *UpdateTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update task default response has a 4xx status code
func (o *UpdateTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update task default response has a 5xx status code
func (o *UpdateTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update task default response a status code equal to that given
func (o *UpdateTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update task default response
func (o *UpdateTaskDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTask default %s", o._statusCode, payload)
}

func (o *UpdateTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /tasks/{id}][%d] updateTask default %s", o._statusCode, payload)
}

func (o *UpdateTaskDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Error-Code
	hdrXErrorCode := response.GetHeader("X-Error-Code")

	if hdrXErrorCode != "" {
		o.XErrorCode = hdrXErrorCode
	}

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}