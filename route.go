// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

// A Route holds all necessary route information.
// This includes the path, before and after middleware as well as the controller.
type Route struct {
	// Path stores the full route path including group prefixes.
	Path string

	// Before is the middleware stack that will be executed before the controller.
	Before MiddlewareChain

	// Controller is the function that defines and returns a result.
	Controller ControllerFunc

	// After is the middleware stack that will be executed after the controller.
	After MiddlewareChain
}

// SetBefore registers the given middlewares as middlewares that will be
// called before the controller method is executed.
func (r *Route) SetBefore(middlewares ...MiddlewareFunc) *Route {
	// There could be already some before middleware registered by the router.
	r.Before = append(r.Before, middlewares...)
	return r
}

// SetAfter registers the given middlewares as middlewares that will be
// called after the controller method is executed.
func (r *Route) SetAfter(middlewares ...MiddlewareFunc) *Route {
	r.After = middlewares
	return r
}

// NewRoute initializes and returns a pointer to a Route struct.
func NewRoute() *Route {
	return &Route{
		Before: make(MiddlewareChain, 0),
		After:  make(MiddlewareChain, 0),
	}
}
