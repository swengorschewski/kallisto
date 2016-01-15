// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"net/http"
	"strings"
)

// A Response inherits the http.ResponseWriter interface and adds some convinience
// methods to return data.
type Response struct {
	http.ResponseWriter

	// Renderer is responsebile for rendering a template.
	Renderer
}

// Data stores the data passed to the templates.
type Data map[string]interface{}

// SetRenderer is the setter for a Renderer.
func (r *Response) SetRenderer(renderer Renderer) {
	r.Renderer = renderer
}

// Text cocantenates the given strings and writes them to the response body.
// In addition Text sets the approriate header for this content type.
func (r *Response) Text(text ...string) {
	r.Header().Set("Content-Type", "text/plain")
	r.Write([]byte(strings.Join(text, "")))
}

// HTML calls the Render method of the Renderer to parse templates
// and sets the approriate header for this content type.
func (r *Response) HTML(data Data, templateNames ...string) {
	r.Header().Set("Content-Type", "text/html")
	r.Render(r, data, templateNames)
}

// Static returns a file identified by the given name and sets the approriate
// header for this content type.
func (r *Response) Static(fileName string, req *http.Request) {
	http.ServeFile(r, req, fileName)
}

// JSON parses the given data to JSON and sets the approriate header for
// this content type.
func (r *Response) JSON(data interface{}) {
	r.Header().Set("Content-Type", "application/json")
}

// XML parses the given data to XML and sets the approriate header for
// this content type.
func (r *Response) XML(data interface{}) {
	r.Header().Set("Content-Type", "text/xml")
}
