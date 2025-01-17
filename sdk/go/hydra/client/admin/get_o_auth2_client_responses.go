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

// GetOAuth2ClientReader is a Reader for the GetOAuth2Client structure.
type GetOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOAuth2ClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetOAuth2ClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOAuth2ClientOK creates a GetOAuth2ClientOK with default headers values
func NewGetOAuth2ClientOK() *GetOAuth2ClientOK {
	return &GetOAuth2ClientOK{}
}

/*GetOAuth2ClientOK handles this case with default header values.

oAuth2Client
*/
type GetOAuth2ClientOK struct {
	Payload *models.Client
}

func (o *GetOAuth2ClientOK) Error() string {
	return fmt.Sprintf("[GET /clients/{id}][%d] getOAuth2ClientOK  %+v", 200, o.Payload)
}

func (o *GetOAuth2ClientOK) GetPayload() *models.Client {
	return o.Payload
}

func (o *GetOAuth2ClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOAuth2ClientNotFound creates a GetOAuth2ClientNotFound with default headers values
func NewGetOAuth2ClientNotFound() *GetOAuth2ClientNotFound {
	return &GetOAuth2ClientNotFound{}
}

/*GetOAuth2ClientNotFound handles this case with default header values.

genericError
*/
type GetOAuth2ClientNotFound struct {
	Payload *models.GenericError
}

func (o *GetOAuth2ClientNotFound) Error() string {
	return fmt.Sprintf("[GET /clients/{id}][%d] getOAuth2ClientNotFound  %+v", 404, o.Payload)
}

func (o *GetOAuth2ClientNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetOAuth2ClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOAuth2ClientInternalServerError creates a GetOAuth2ClientInternalServerError with default headers values
func NewGetOAuth2ClientInternalServerError() *GetOAuth2ClientInternalServerError {
	return &GetOAuth2ClientInternalServerError{}
}

/*GetOAuth2ClientInternalServerError handles this case with default header values.

genericError
*/
type GetOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[GET /clients/{id}][%d] getOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOAuth2ClientInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
