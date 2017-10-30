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

// NewListSharesParams creates a new ListSharesParams object
// with the default values initialized.
func NewListSharesParams() *ListSharesParams {
	var ()
	return &ListSharesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSharesParamsWithTimeout creates a new ListSharesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSharesParamsWithTimeout(timeout time.Duration) *ListSharesParams {
	var ()
	return &ListSharesParams{

		timeout: timeout,
	}
}

// NewListSharesParamsWithContext creates a new ListSharesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSharesParamsWithContext(ctx context.Context) *ListSharesParams {
	var ()
	return &ListSharesParams{

		Context: ctx,
	}
}

// NewListSharesParamsWithHTTPClient creates a new ListSharesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSharesParamsWithHTTPClient(client *http.Client) *ListSharesParams {
	var ()
	return &ListSharesParams{
		HTTPClient: client,
	}
}

/*ListSharesParams contains all the parameters to send to the API endpoint
for the list shares operation typically these are written to a http.Request
*/
type ListSharesParams struct {

	/*ResourceType
	  Type of specified share backend resource

	*/
	ResourceType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list shares params
func (o *ListSharesParams) WithTimeout(timeout time.Duration) *ListSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list shares params
func (o *ListSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list shares params
func (o *ListSharesParams) WithContext(ctx context.Context) *ListSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list shares params
func (o *ListSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list shares params
func (o *ListSharesParams) WithHTTPClient(client *http.Client) *ListSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list shares params
func (o *ListSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceType adds the resourceType to the list shares params
func (o *ListSharesParams) WithResourceType(resourceType string) *ListSharesParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the list shares params
func (o *ListSharesParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WriteToRequest writes these params to a swagger request
func (o *ListSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
