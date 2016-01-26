// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

// Router is a http.Handler and can be used with any net/http packages which
// dispatches requests to matching routes.
//
// It is a wrapper struct for Julien Schmidts httprouter.
type Router struct {
	// Holds all registered middlewares and applies them to every handler.
	middlewares MiddlewareChain

	// Holds a reference to the main application.
	kallisto *Kallisto

	// httprouter holds a reference to Julien Schmidts router.
	httprouter *httprouter.Router

	// Can be set to prepend a common path prefix to all registered routes.
	// This is used in route groups.
	pathPrefix string

	// Can be set to prepend a common name prefix to all registered routes.
	// This is used in route groups.
	namePrefix string
}

// MiddlewareFunc is the signature of a middleware function.
type MiddlewareFunc func(*Context)

// MiddlewareChain is an array of middlewares.
type MiddlewareChain []MiddlewareFunc

// ControllerFunc is the signature of a controller function.
type ControllerFunc func(*Context, *Response)

// NewRouter returns a pointer to an initialized Router struct.
func NewRouter(k *Kallisto) *Router {
	return &Router{
		kallisto:    k,
		httprouter:  httprouter.New(),
		middlewares: make([]MiddlewareFunc, 0),
	}
}

// Use registers middleware for all routes of the router.
func (r *Router) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = middlewares
}

// GET registers HTTP GET request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) GET(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("GET", path, name, controller)
}

// POST registers HTTP POST request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) POST(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("POST", path, name, controller)
}

// PATCH registers HTTP PATCH request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) PATCH(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("PATCH", path, name, controller)
}

// PUT registers HTTP PUT request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) PUT(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("PUT", path, name, controller)
}

// DELETE registers HTTP DELETE request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) DELETE(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("DELETE", path, name, controller)
}

// HEAD registers HTTP HEAD request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) HEAD(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("HEAD", path, name, controller)
}

// OPTIONS registers HTTP OPTIONS request handles for the specified path.
//
// The name parameter is used to identify the route independent of its path.
func (r *Router) OPTIONS(path string, name string, controller ControllerFunc) *Route {
	return r.Handle("OPTIONS", path, name, controller)
}

// NotFound sets a controller as a custom NotFound handler.
func (r *Router) NotFound(c ControllerFunc) {

	route := NewRoute()
	route.Controller = c

	r.httprouter.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(404)
		ctx := newContext(r.kallisto, route)
		ctx.Request = req
		ctx.Response = newResponse(w, ctx)
		ctx.Next()
	})
}

// PanicHandler is a wrapper for the httprouter PanicHandler method to work
// with a ControllerFunc.
func (r *Router) PanicHandler(c ControllerFunc) {
	route := NewRoute()
	route.Controller = c

	r.httprouter.PanicHandler = func(w http.ResponseWriter, req *http.Request, stack interface{}) {
		ctx := newContext(r.kallisto, route)
		ctx.Request = req
		ctx.Response = newResponse(w, ctx)

		ctx.Set("PanicStack", stack)
		ctx.Next()
	}
}

// Group returns a router with the given path and name prefixes and middlewares.
// Routes with common path or name prefixes could be registered via the group method.
func (r *Router) Group(pathPrefix string, namePrefix string, fn func(*Router), middlewares ...MiddlewareFunc) {
	fn(&Router{
		middlewares: append(r.middlewares, middlewares...),
		kallisto:    r.kallisto,
		httprouter:  r.httprouter,
		pathPrefix:  path.Join(r.pathPrefix, pathPrefix),
		namePrefix:  r.namePrefix + namePrefix,
	})
}

// Handle registers handlers for for the supplied path.
//
// Shortcut methods are available the standard HTTP methods GET, POST, PUT, PATCH and DELETE.
func (r *Router) Handle(method string, uri string, name string, controller ControllerFunc) *Route {
	route := &Route{
		Path:       path.Join(r.pathPrefix, uri),
		Before:     r.middlewares,
		After:      make([]MiddlewareFunc, 0),
		Controller: controller,
	}

	r.kallisto.routes[r.namePrefix+name] = route

	r.httprouter.Handle(method, path.Join(r.pathPrefix, uri), func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		ctx := newContext(r.kallisto, route)
		ctx.Params = ps
		ctx.Request = req
		ctx.Response = newResponse(w, ctx)

		ctx.Next()
	})

	return route
}

// ServeHTTP is the necessary method to implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.httprouter.ServeHTTP(w, req)
}

// ServeStatic registers a route to the content which should be served as static files.
func (r *Router) ServeStatic(path string, root http.FileSystem) {
	r.httprouter.ServeFiles(path, root)
}
