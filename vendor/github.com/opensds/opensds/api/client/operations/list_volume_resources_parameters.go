package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListVolumeResourcesParams creates a new ListVolumeResourcesParams object
// with the default values initialized.
func NewListVolumeResourcesParams() *ListVolumeResourcesParams {

	return &ListVolumeResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVolumeResourcesParamsWithTimeout creates a new ListVolumeResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVolumeResourcesParamsWithTimeout(timeout time.Duration) *ListVolumeResourcesParams {

	return &ListVolumeResourcesParams{

		timeout: timeout,
	}
}

// NewListVolumeResourcesParamsWithContext creates a new ListVolumeResourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListVolumeResourcesParamsWithContext(ctx context.Context) *ListVolumeResourcesParams {

	return &ListVolumeResourcesParams{

		Context: ctx,
	}
}

// NewListVolumeResourcesParamsWithHTTPClient creates a new ListVolumeResourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListVolumeResourcesParamsWithHTTPClient(client *http.Client) *ListVolumeResourcesParams {

	return &ListVolumeResourcesParams{
		HTTPClient: client,
	}
}

/*ListVolumeResourcesParams contains all the parameters to send to the API endpoint
for the list volume resources operation typically these are written to a http.Request
*/
type ListVolumeResourcesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list volume resources params
func (o *ListVolumeResourcesParams) WithTimeout(timeout time.Duration) *ListVolumeResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list volume resources params
func (o *ListVolumeResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list volume resources params
func (o *ListVolumeResourcesParams) WithContext(ctx context.Context) *ListVolumeResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list volume resources params
func (o *ListVolumeResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list volume resources params
func (o *ListVolumeResourcesParams) WithHTTPClient(client *http.Client) *ListVolumeResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list volume resources params
func (o *ListVolumeResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListVolumeResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
