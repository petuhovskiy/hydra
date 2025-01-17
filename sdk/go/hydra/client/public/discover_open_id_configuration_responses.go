// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/petuhovskiy/hydra/sdk/go/hydra/models"
)

// DiscoverOpenIDConfigurationReader is a Reader for the DiscoverOpenIDConfiguration structure.
type DiscoverOpenIDConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverOpenIDConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverOpenIDConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDiscoverOpenIDConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDiscoverOpenIDConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDiscoverOpenIDConfigurationOK creates a DiscoverOpenIDConfigurationOK with default headers values
func NewDiscoverOpenIDConfigurationOK() *DiscoverOpenIDConfigurationOK {
	return &DiscoverOpenIDConfigurationOK{}
}

/*DiscoverOpenIDConfigurationOK handles this case with default header values.

wellKnown
*/
type DiscoverOpenIDConfigurationOK struct {
	Payload *models.WellKnown
}

func (o *DiscoverOpenIDConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationOK  %+v", 200, o.Payload)
}

func (o *DiscoverOpenIDConfigurationOK) GetPayload() *models.WellKnown {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WellKnown)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationUnauthorized creates a DiscoverOpenIDConfigurationUnauthorized with default headers values
func NewDiscoverOpenIDConfigurationUnauthorized() *DiscoverOpenIDConfigurationUnauthorized {
	return &DiscoverOpenIDConfigurationUnauthorized{}
}

/*DiscoverOpenIDConfigurationUnauthorized handles this case with default header values.

genericError
*/
type DiscoverOpenIDConfigurationUnauthorized struct {
	Payload *models.GenericError
}

func (o *DiscoverOpenIDConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *DiscoverOpenIDConfigurationUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationInternalServerError creates a DiscoverOpenIDConfigurationInternalServerError with default headers values
func NewDiscoverOpenIDConfigurationInternalServerError() *DiscoverOpenIDConfigurationInternalServerError {
	return &DiscoverOpenIDConfigurationInternalServerError{}
}

/*DiscoverOpenIDConfigurationInternalServerError handles this case with default header values.

genericError
*/
type DiscoverOpenIDConfigurationInternalServerError struct {
	Payload *models.GenericError
}

func (o *DiscoverOpenIDConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *DiscoverOpenIDConfigurationInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
