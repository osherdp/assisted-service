// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewGetHostIgnitionParams creates a new GetHostIgnitionParams object
// with the default values initialized.
func NewGetHostIgnitionParams() *GetHostIgnitionParams {
	var ()
	return &GetHostIgnitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostIgnitionParamsWithTimeout creates a new GetHostIgnitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHostIgnitionParamsWithTimeout(timeout time.Duration) *GetHostIgnitionParams {
	var ()
	return &GetHostIgnitionParams{

		timeout: timeout,
	}
}

// NewGetHostIgnitionParamsWithContext creates a new GetHostIgnitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHostIgnitionParamsWithContext(ctx context.Context) *GetHostIgnitionParams {
	var ()
	return &GetHostIgnitionParams{

		Context: ctx,
	}
}

// NewGetHostIgnitionParamsWithHTTPClient creates a new GetHostIgnitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHostIgnitionParamsWithHTTPClient(client *http.Client) *GetHostIgnitionParams {
	var ()
	return &GetHostIgnitionParams{
		HTTPClient: client,
	}
}

/*GetHostIgnitionParams contains all the parameters to send to the API endpoint
for the get host ignition operation typically these are written to a http.Request
*/
type GetHostIgnitionParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID
	/*HostID*/
	HostID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get host ignition params
func (o *GetHostIgnitionParams) WithTimeout(timeout time.Duration) *GetHostIgnitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host ignition params
func (o *GetHostIgnitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host ignition params
func (o *GetHostIgnitionParams) WithContext(ctx context.Context) *GetHostIgnitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host ignition params
func (o *GetHostIgnitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host ignition params
func (o *GetHostIgnitionParams) WithHTTPClient(client *http.Client) *GetHostIgnitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host ignition params
func (o *GetHostIgnitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get host ignition params
func (o *GetHostIgnitionParams) WithClusterID(clusterID strfmt.UUID) *GetHostIgnitionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get host ignition params
func (o *GetHostIgnitionParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the get host ignition params
func (o *GetHostIgnitionParams) WithHostID(hostID strfmt.UUID) *GetHostIgnitionParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the get host ignition params
func (o *GetHostIgnitionParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostIgnitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
