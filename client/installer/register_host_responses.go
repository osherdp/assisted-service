// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// RegisterHostReader is a Reader for the RegisterHost structure.
type RegisterHostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterHostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterHostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterHostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRegisterHostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRegisterHostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRegisterHostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewRegisterHostMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRegisterHostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegisterHostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewRegisterHostServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterHostCreated creates a RegisterHostCreated with default headers values
func NewRegisterHostCreated() *RegisterHostCreated {
	return &RegisterHostCreated{}
}

/*RegisterHostCreated handles this case with default header values.

Success.
*/
type RegisterHostCreated struct {
	Payload *models.HostRegistrationResponse
}

func (o *RegisterHostCreated) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostCreated  %+v", 201, o.Payload)
}

func (o *RegisterHostCreated) GetPayload() *models.HostRegistrationResponse {
	return o.Payload
}

func (o *RegisterHostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostRegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostBadRequest creates a RegisterHostBadRequest with default headers values
func NewRegisterHostBadRequest() *RegisterHostBadRequest {
	return &RegisterHostBadRequest{}
}

/*RegisterHostBadRequest handles this case with default header values.

Error.
*/
type RegisterHostBadRequest struct {
	Payload *models.Error
}

func (o *RegisterHostBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostBadRequest  %+v", 400, o.Payload)
}

func (o *RegisterHostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterHostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostUnauthorized creates a RegisterHostUnauthorized with default headers values
func NewRegisterHostUnauthorized() *RegisterHostUnauthorized {
	return &RegisterHostUnauthorized{}
}

/*RegisterHostUnauthorized handles this case with default header values.

Unauthorized.
*/
type RegisterHostUnauthorized struct {
	Payload *models.InfraError
}

func (o *RegisterHostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostUnauthorized  %+v", 401, o.Payload)
}

func (o *RegisterHostUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegisterHostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostForbidden creates a RegisterHostForbidden with default headers values
func NewRegisterHostForbidden() *RegisterHostForbidden {
	return &RegisterHostForbidden{}
}

/*RegisterHostForbidden handles this case with default header values.

Forbidden.
*/
type RegisterHostForbidden struct {
	Payload *models.InfraError
}

func (o *RegisterHostForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostForbidden  %+v", 403, o.Payload)
}

func (o *RegisterHostForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegisterHostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostNotFound creates a RegisterHostNotFound with default headers values
func NewRegisterHostNotFound() *RegisterHostNotFound {
	return &RegisterHostNotFound{}
}

/*RegisterHostNotFound handles this case with default header values.

Error.
*/
type RegisterHostNotFound struct {
	Payload *models.Error
}

func (o *RegisterHostNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostNotFound  %+v", 404, o.Payload)
}

func (o *RegisterHostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterHostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostMethodNotAllowed creates a RegisterHostMethodNotAllowed with default headers values
func NewRegisterHostMethodNotAllowed() *RegisterHostMethodNotAllowed {
	return &RegisterHostMethodNotAllowed{}
}

/*RegisterHostMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type RegisterHostMethodNotAllowed struct {
	Payload *models.Error
}

func (o *RegisterHostMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *RegisterHostMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterHostMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostConflict creates a RegisterHostConflict with default headers values
func NewRegisterHostConflict() *RegisterHostConflict {
	return &RegisterHostConflict{}
}

/*RegisterHostConflict handles this case with default header values.

Cluster cannot accept new hosts due to its current state.
*/
type RegisterHostConflict struct {
	Payload *models.InfraError
}

func (o *RegisterHostConflict) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostConflict  %+v", 409, o.Payload)
}

func (o *RegisterHostConflict) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *RegisterHostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostInternalServerError creates a RegisterHostInternalServerError with default headers values
func NewRegisterHostInternalServerError() *RegisterHostInternalServerError {
	return &RegisterHostInternalServerError{}
}

/*RegisterHostInternalServerError handles this case with default header values.

Error.
*/
type RegisterHostInternalServerError struct {
	Payload *models.Error
}

func (o *RegisterHostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostInternalServerError  %+v", 500, o.Payload)
}

func (o *RegisterHostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterHostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterHostServiceUnavailable creates a RegisterHostServiceUnavailable with default headers values
func NewRegisterHostServiceUnavailable() *RegisterHostServiceUnavailable {
	return &RegisterHostServiceUnavailable{}
}

/*RegisterHostServiceUnavailable handles this case with default header values.

Unavailable.
*/
type RegisterHostServiceUnavailable struct {
	Payload *models.Error
}

func (o *RegisterHostServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/hosts][%d] registerHostServiceUnavailable  %+v", 503, o.Payload)
}

func (o *RegisterHostServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterHostServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
