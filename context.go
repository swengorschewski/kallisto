// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// A Context holds information for a HTTP request.
//
// It is passed to every middleware and controller that handles the
// request to provide a state. Furthermore it stores a reference to all
// middlewares and the controller.
type Context struct {

	// middlewareIndex is the index pointer for the MiddlewareChains stored
	// in Route struct. It is initialized with -1 one and for each middleware called
	// it will be incremented by one.
	middlewareIndex int8

	// doneBefore is a switch to signalzie if the before middlewares of the
	// route param is are called. If so the controller will be executed
	// and after that the after middlewares will be executed.
	doneBefore bool

	// kallisto stores a reference to the app.
	kallisto *Kallisto

	// route stores all route information like the requested path and name as well as
	// before and after middlewares and the controller.
	route *Route

	// Params is the httprouter paramter store.
	Params httprouter.Params

	// Request is a pointer to the http request.
	Request *http.Request

	// Response a helper struct to write different responses back to the client.
	Response *Response

	// Session stores a session for this client.
	Session Session

	// data is a key value store to keep data for this request.
	data map[string]interface{}
}

// newContext creates a Context struct with the given parameters and returns a pointer.
func newContext(k *Kallisto, r *Route) *Context {
	return &Context{
		middlewareIndex: -1,
		doneBefore:      false,
		kallisto:        k,
		route:           r,
		data:            make(map[string]interface{}),
	}
}

// Param returns a route param indentified by the given key.
func (c *Context) Param(key string) string {
	return c.Params.ByName(key)
}

// Next calls the next middleware in the middleware stack or if appropriate
// the controler and passes a pointer to itself to the middleware/controller.
func (c *Context) Next() {
	if c.middlewareIndex < int8(len(c.route.Before)-1) && !c.doneBefore {
		c.middlewareIndex++
		c.route.Before[c.middlewareIndex](c)
		c.Next()
	} else if c.middlewareIndex == int8(len(c.route.Before)-1) && !c.doneBefore {
		c.doneBefore = true
		c.middlewareIndex = -1

		c.route.Controller(c, c.Response)

		c.Next()
	} else if c.middlewareIndex < int8(len(c.route.After)-1) && c.doneBefore {
		c.middlewareIndex++
		c.route.After[c.middlewareIndex](c)
		c.Next()
	}
}

// Set stores the given key value pair.
// Anything stored via Set is request scoped.
func (c *Context) Set(key string, value interface{}) {
	c.data[key] = value
}

// Get returns a request scoped value identified by the given key.
func (c *Context) Get(key string) interface{} {
	return c.data[key]
}

// App returns a pointer to the application.
func (c *Context) App() *Kallisto {
	return c.kallisto
}

// SetSession sets the given session and stores it in the current context.
func (c *Context) SetSession(s Session) {
	c.Session = s
}
