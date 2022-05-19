// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewServiceBrokerHealthHeadParams creates a new ServiceBrokerHealthHeadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceBrokerHealthHeadParams() *ServiceBrokerHealthHeadParams {
	return &ServiceBrokerHealthHeadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerHealthHeadParamsWithTimeout creates a new ServiceBrokerHealthHeadParams object
// with the ability to set a timeout on a request.
func NewServiceBrokerHealthHeadParamsWithTimeout(timeout time.Duration) *ServiceBrokerHealthHeadParams {
	return &ServiceBrokerHealthHeadParams{
		timeout: timeout,
	}
}

// NewServiceBrokerHealthHeadParamsWithContext creates a new ServiceBrokerHealthHeadParams object
// with the ability to set a context for a request.
func NewServiceBrokerHealthHeadParamsWithContext(ctx context.Context) *ServiceBrokerHealthHeadParams {
	return &ServiceBrokerHealthHeadParams{
		Context: ctx,
	}
}

// NewServiceBrokerHealthHeadParamsWithHTTPClient creates a new ServiceBrokerHealthHeadParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceBrokerHealthHeadParamsWithHTTPClient(client *http.Client) *ServiceBrokerHealthHeadParams {
	return &ServiceBrokerHealthHeadParams{
		HTTPClient: client,
	}
}

/* ServiceBrokerHealthHeadParams contains all the parameters to send to the API endpoint
   for the service broker health head operation.

   Typically these are written to a http.Request.
*/
type ServiceBrokerHealthHeadParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service broker health head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerHealthHeadParams) WithDefaults() *ServiceBrokerHealthHeadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service broker health head params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerHealthHeadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) WithTimeout(timeout time.Duration) *ServiceBrokerHealthHeadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) WithContext(ctx context.Context) *ServiceBrokerHealthHeadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) WithHTTPClient(client *http.Client) *ServiceBrokerHealthHeadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker health head params
func (o *ServiceBrokerHealthHeadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerHealthHeadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}