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

// GetLogoutRequestReader is a Reader for the GetLogoutRequest structure.
type GetLogoutRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogoutRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogoutRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogoutRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogoutRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogoutRequestOK creates a GetLogoutRequestOK with default headers values
func NewGetLogoutRequestOK() *GetLogoutRequestOK {
	return &GetLogoutRequestOK{}
}

/*GetLogoutRequestOK handles this case with default header values.

logoutRequest
*/
type GetLogoutRequestOK struct {
	Payload *models.LogoutRequest
}

func (o *GetLogoutRequestOK) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/logout][%d] getLogoutRequestOK  %+v", 200, o.Payload)
}

func (o *GetLogoutRequestOK) GetPayload() *models.LogoutRequest {
	return o.Payload
}

func (o *GetLogoutRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogoutRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogoutRequestNotFound creates a GetLogoutRequestNotFound with default headers values
func NewGetLogoutRequestNotFound() *GetLogoutRequestNotFound {
	return &GetLogoutRequestNotFound{}
}

/*GetLogoutRequestNotFound handles this case with default header values.

genericError
*/
type GetLogoutRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetLogoutRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/logout][%d] getLogoutRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetLogoutRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetLogoutRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogoutRequestInternalServerError creates a GetLogoutRequestInternalServerError with default headers values
func NewGetLogoutRequestInternalServerError() *GetLogoutRequestInternalServerError {
	return &GetLogoutRequestInternalServerError{}
}

/*GetLogoutRequestInternalServerError handles this case with default header values.

genericError
*/
type GetLogoutRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetLogoutRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oauth2/auth/requests/logout][%d] getLogoutRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogoutRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetLogoutRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
