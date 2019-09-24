// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/petuhovskiy/hydra/sdk/go/hydra/models"
)

// GetConsentRequestReader is a Reader for the GetConsentRequest structure.
type GetConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConsentRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetConsentRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetConsentRequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConsentRequestOK creates a GetConsentRequestOK with default headers values
func NewGetConsentRequestOK() *GetConsentRequestOK {
	return &GetConsentRequestOK{}
}

/*GetConsentRequestOK handles this case with default header values.

consentRequest
*/
type GetConsentRequestOK struct {
	Payload *models.ConsentRequest
}

func (o *GetConsentRequestOK) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/consent][%d] getConsentRequestOK  %+v", 200, o.Payload)
}

func (o *GetConsentRequestOK) GetPayload() *models.ConsentRequest {
	return o.Payload
}

func (o *GetConsentRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentRequestNotFound creates a GetConsentRequestNotFound with default headers values
func NewGetConsentRequestNotFound() *GetConsentRequestNotFound {
	return &GetConsentRequestNotFound{}
}

/*GetConsentRequestNotFound handles this case with default header values.

genericError
*/
type GetConsentRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetConsentRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/consent][%d] getConsentRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetConsentRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetConsentRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentRequestConflict creates a GetConsentRequestConflict with default headers values
func NewGetConsentRequestConflict() *GetConsentRequestConflict {
	return &GetConsentRequestConflict{}
}

/*GetConsentRequestConflict handles this case with default header values.

genericError
*/
type GetConsentRequestConflict struct {
	Payload *models.GenericError
}

func (o *GetConsentRequestConflict) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/consent][%d] getConsentRequestConflict  %+v", 409, o.Payload)
}

func (o *GetConsentRequestConflict) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetConsentRequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConsentRequestInternalServerError creates a GetConsentRequestInternalServerError with default headers values
func NewGetConsentRequestInternalServerError() *GetConsentRequestInternalServerError {
	return &GetConsentRequestInternalServerError{}
}

/*GetConsentRequestInternalServerError handles this case with default header values.

genericError
*/
type GetConsentRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/consent][%d] getConsentRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConsentRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
