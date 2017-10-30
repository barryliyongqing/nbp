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

	"github.com/opensds/opensds/api/models"
)

// NewOperateVolumeParams creates a new OperateVolumeParams object
// with the default values initialized.
func NewOperateVolumeParams() *OperateVolumeParams {
	var ()
	return &OperateVolumeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperateVolumeParamsWithTimeout creates a new OperateVolumeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperateVolumeParamsWithTimeout(timeout time.Duration) *OperateVolumeParams {
	var ()
	return &OperateVolumeParams{

		timeout: timeout,
	}
}

// NewOperateVolumeParamsWithContext creates a new OperateVolumeParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperateVolumeParamsWithContext(ctx context.Context) *OperateVolumeParams {
	var ()
	return &OperateVolumeParams{

		Context: ctx,
	}
}

// NewOperateVolumeParamsWithHTTPClient creates a new OperateVolumeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperateVolumeParamsWithHTTPClient(client *http.Client) *OperateVolumeParams {
	var ()
	return &OperateVolumeParams{
		HTTPClient: client,
	}
}

/*OperateVolumeParams contains all the parameters to send to the API endpoint
for the operate volume operation typically these are written to a http.Request
*/
type OperateVolumeParams struct {

	/*ID
	  ID of specified volume

	*/
	ID string
	/*ResourceType
	  Type of specified volume backend resource

	*/
	ResourceType string
	/*VolumeRequest
	  Volume request object

	*/
	VolumeRequest *models.VolumeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operate volume params
func (o *OperateVolumeParams) WithTimeout(timeout time.Duration) *OperateVolumeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operate volume params
func (o *OperateVolumeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operate volume params
func (o *OperateVolumeParams) WithContext(ctx context.Context) *OperateVolumeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operate volume params
func (o *OperateVolumeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operate volume params
func (o *OperateVolumeParams) WithHTTPClient(client *http.Client) *OperateVolumeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operate volume params
func (o *OperateVolumeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the operate volume params
func (o *OperateVolumeParams) WithID(id string) *OperateVolumeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the operate volume params
func (o *OperateVolumeParams) SetID(id string) {
	o.ID = id
}

// WithResourceType adds the resourceType to the operate volume params
func (o *OperateVolumeParams) WithResourceType(resourceType string) *OperateVolumeParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the operate volume params
func (o *OperateVolumeParams) SetResourceType(resourceType string) {
	o.ResourceType = resourceType
}

// WithVolumeRequest adds the volumeRequest to the operate volume params
func (o *OperateVolumeParams) WithVolumeRequest(volumeRequest *models.VolumeRequest) *OperateVolumeParams {
	o.SetVolumeRequest(volumeRequest)
	return o
}

// SetVolumeRequest adds the volumeRequest to the operate volume params
func (o *OperateVolumeParams) SetVolumeRequest(volumeRequest *models.VolumeRequest) {
	o.VolumeRequest = volumeRequest
}

// WriteToRequest writes these params to a swagger request
func (o *OperateVolumeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param resourceType
	if err := r.SetPathParam("resourceType", o.ResourceType); err != nil {
		return err
	}

	if o.VolumeRequest == nil {
		o.VolumeRequest = new(models.VolumeRequest)
	}

	if err := r.SetBodyParam(o.VolumeRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
