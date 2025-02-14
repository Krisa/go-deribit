// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPublicGetLastTradesByCurrencyParams creates a new GetPublicGetLastTradesByCurrencyParams object
// with the default values initialized.
func NewGetPublicGetLastTradesByCurrencyParams() *GetPublicGetLastTradesByCurrencyParams {
	var ()
	return &GetPublicGetLastTradesByCurrencyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicGetLastTradesByCurrencyParamsWithTimeout creates a new GetPublicGetLastTradesByCurrencyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicGetLastTradesByCurrencyParamsWithTimeout(timeout time.Duration) *GetPublicGetLastTradesByCurrencyParams {
	var ()
	return &GetPublicGetLastTradesByCurrencyParams{

		timeout: timeout,
	}
}

// NewGetPublicGetLastTradesByCurrencyParamsWithContext creates a new GetPublicGetLastTradesByCurrencyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicGetLastTradesByCurrencyParamsWithContext(ctx context.Context) *GetPublicGetLastTradesByCurrencyParams {
	var ()
	return &GetPublicGetLastTradesByCurrencyParams{

		Context: ctx,
	}
}

// NewGetPublicGetLastTradesByCurrencyParamsWithHTTPClient creates a new GetPublicGetLastTradesByCurrencyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicGetLastTradesByCurrencyParamsWithHTTPClient(client *http.Client) *GetPublicGetLastTradesByCurrencyParams {
	var ()
	return &GetPublicGetLastTradesByCurrencyParams{
		HTTPClient: client,
	}
}

/*GetPublicGetLastTradesByCurrencyParams contains all the parameters to send to the API endpoint
for the get public get last trades by currency operation typically these are written to a http.Request
*/
type GetPublicGetLastTradesByCurrencyParams struct {

	/*Count
	  Number of requested items, default - `10`

	*/
	Count *int64
	/*Currency
	  The currency symbol

	*/
	Currency string
	/*EndID
	  The ID number of the last trade to be returned

	*/
	EndID *string
	/*IncludeOld
	  Include trades older than 7 days, default - `false`

	*/
	IncludeOld *bool
	/*Kind
	  Instrument kind, if not provided instruments of all kinds are considered

	*/
	Kind *string
	/*Sorting
	  Direction of results sorting (`default` value means no sorting, results will be returned in order in which they left the database)

	*/
	Sorting *string
	/*StartID
	  The ID number of the first trade to be returned

	*/
	StartID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithTimeout(timeout time.Duration) *GetPublicGetLastTradesByCurrencyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithContext(ctx context.Context) *GetPublicGetLastTradesByCurrencyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithHTTPClient(client *http.Client) *GetPublicGetLastTradesByCurrencyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithCount(count *int64) *GetPublicGetLastTradesByCurrencyParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetCount(count *int64) {
	o.Count = count
}

// WithCurrency adds the currency to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithCurrency(currency string) *GetPublicGetLastTradesByCurrencyParams {
	o.SetCurrency(currency)
	return o
}

// SetCurrency adds the currency to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetCurrency(currency string) {
	o.Currency = currency
}

// WithEndID adds the endID to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithEndID(endID *string) *GetPublicGetLastTradesByCurrencyParams {
	o.SetEndID(endID)
	return o
}

// SetEndID adds the endId to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetEndID(endID *string) {
	o.EndID = endID
}

// WithIncludeOld adds the includeOld to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithIncludeOld(includeOld *bool) *GetPublicGetLastTradesByCurrencyParams {
	o.SetIncludeOld(includeOld)
	return o
}

// SetIncludeOld adds the includeOld to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetIncludeOld(includeOld *bool) {
	o.IncludeOld = includeOld
}

// WithKind adds the kind to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithKind(kind *string) *GetPublicGetLastTradesByCurrencyParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithSorting adds the sorting to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithSorting(sorting *string) *GetPublicGetLastTradesByCurrencyParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithStartID adds the startID to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) WithStartID(startID *string) *GetPublicGetLastTradesByCurrencyParams {
	o.SetStartID(startID)
	return o
}

// SetStartID adds the startId to the get public get last trades by currency params
func (o *GetPublicGetLastTradesByCurrencyParams) SetStartID(startID *string) {
	o.StartID = startID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicGetLastTradesByCurrencyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	// query param currency
	qrCurrency := o.Currency
	qCurrency := qrCurrency
	if qCurrency != "" {
		if err := r.SetQueryParam("currency", qCurrency); err != nil {
			return err
		}
	}

	if o.EndID != nil {

		// query param end_id
		var qrEndID string
		if o.EndID != nil {
			qrEndID = *o.EndID
		}
		qEndID := qrEndID
		if qEndID != "" {
			if err := r.SetQueryParam("end_id", qEndID); err != nil {
				return err
			}
		}

	}

	if o.IncludeOld != nil {

		// query param include_old
		var qrIncludeOld bool
		if o.IncludeOld != nil {
			qrIncludeOld = *o.IncludeOld
		}
		qIncludeOld := swag.FormatBool(qrIncludeOld)
		if qIncludeOld != "" {
			if err := r.SetQueryParam("include_old", qIncludeOld); err != nil {
				return err
			}
		}

	}

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if o.StartID != nil {

		// query param start_id
		var qrStartID string
		if o.StartID != nil {
			qrStartID = *o.StartID
		}
		qStartID := qrStartID
		if qStartID != "" {
			if err := r.SetQueryParam("start_id", qStartID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
