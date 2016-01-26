// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"net/http"
	"strings"
)

// A Response inherits the http.ResponseWriter interface and adds some convenience
// methods to return data.
type Response struct {
	http.ResponseWriter

	// Renderer is responsible for rendering a template.
	Renderer

	ctx *Context
}

// newResponse creates a Response struct and returns a pointer to it.
func newResponse(w http.ResponseWriter, ctx *Context) *Response {
	return &Response{
		ResponseWriter: w,
		ctx:            ctx,
	}
}

// SetRenderer is the setter for a Renderer.
func (r *Response) SetRenderer(renderer Renderer) {
	r.Renderer = renderer
}

// Text cocatenates the given strings and writes them to the response body.
// In addition Text sets the appropriate header for this content type.
func (r *Response) Text(text ...string) {
	r.Header().Set("Content-Type", "text/plain")
	r.Write([]byte(strings.Join(text, "")))
}

// HTML calls the Render method of the Renderer to parse the given templates
// file names and sets the appropriate header for this content type.
func (r *Response) HTML(data Data, fileNames ...string) {
	if data != nil {
		for k, v := range data {
			r.ctx.Data[k] = v
		}
	}
	r.Header().Set("Content-Type", "text/html")
	r.Render(r, r.ctx.Data, fileNames)
}

// Static returns a file identified by the given name and sets the appropriate
// header for this content type.
func (r *Response) Static(fileName string, req *http.Request) {
	http.ServeFile(r, req, fileName)
}

// JSON parses the given data to JSON and sets the appropriate header for
// this content type.
func (r *Response) JSON(data interface{}) {
	r.Header().Set("Content-Type", "application/json")
}

// XML parses the given data to XML and sets the appropriate header for
// this content type.
func (r *Response) XML(data interface{}) {
	r.Header().Set("Content-Type", "text/xml")
}
