// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetDegreesOfSeparationParams creates a new GetDegreesOfSeparationParams object
//
// There are no default values defined in the spec.
func NewGetDegreesOfSeparationParams() GetDegreesOfSeparationParams {

	return GetDegreesOfSeparationParams{}
}

// GetDegreesOfSeparationParams contains all the bound params for the get degrees of separation operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetDegreesOfSeparation
type GetDegreesOfSeparationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ActorName string
	/*
	  Required: true
	  In: path
	*/
	CastRestriction int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetDegreesOfSeparationParams() beforehand.
func (o *GetDegreesOfSeparationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rActorName, rhkActorName, _ := route.Params.GetOK("actorName")
	if err := o.bindActorName(rActorName, rhkActorName, route.Formats); err != nil {
		res = append(res, err)
	}

	rCastRestriction, rhkCastRestriction, _ := route.Params.GetOK("castRestriction")
	if err := o.bindCastRestriction(rCastRestriction, rhkCastRestriction, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindActorName binds and validates parameter ActorName from path.
func (o *GetDegreesOfSeparationParams) bindActorName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ActorName = raw

	return nil
}

// bindCastRestriction binds and validates parameter CastRestriction from path.
func (o *GetDegreesOfSeparationParams) bindCastRestriction(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("castRestriction", "path", "int64", raw)
	}
	o.CastRestriction = value

	return nil
}
