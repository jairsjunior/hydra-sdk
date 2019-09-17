// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jairsjunior/hydra-sdk/models"
)

// UserinfoReader is a Reader for the Userinfo structure.
type UserinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserinfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserinfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserinfoOK creates a UserinfoOK with default headers values
func NewUserinfoOK() *UserinfoOK {
	return &UserinfoOK{}
}

/*UserinfoOK handles this case with default header values.

userinfoResponse
*/
type UserinfoOK struct {
	Payload *models.SwaggeruserinfoResponsePayload
}

func (o *UserinfoOK) Error() string {
	return fmt.Sprintf("[GET /userinfo][%d] userinfoOK  %+v", 200, o.Payload)
}

func (o *UserinfoOK) GetPayload() *models.SwaggeruserinfoResponsePayload {
	return o.Payload
}

func (o *UserinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggeruserinfoResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserinfoUnauthorized creates a UserinfoUnauthorized with default headers values
func NewUserinfoUnauthorized() *UserinfoUnauthorized {
	return &UserinfoUnauthorized{}
}

/*UserinfoUnauthorized handles this case with default header values.

genericError
*/
type UserinfoUnauthorized struct {
	Payload *models.GenericError
}

func (o *UserinfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /userinfo][%d] userinfoUnauthorized  %+v", 401, o.Payload)
}

func (o *UserinfoUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UserinfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserinfoInternalServerError creates a UserinfoInternalServerError with default headers values
func NewUserinfoInternalServerError() *UserinfoInternalServerError {
	return &UserinfoInternalServerError{}
}

/*UserinfoInternalServerError handles this case with default header values.

genericError
*/
type UserinfoInternalServerError struct {
	Payload *models.GenericError
}

func (o *UserinfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /userinfo][%d] userinfoInternalServerError  %+v", 500, o.Payload)
}

func (o *UserinfoInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UserinfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
