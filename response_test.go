// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
)

var b = "test"
var ctx = newContext(nil, nil)

type renderer struct{}

func (r renderer) Render(w io.Writer, v interface{}, t []string) {
	fmt.Fprintf(w, "%s", b)
}

func TestText(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)
	r.Text("first", ":second:", "last")

	if w.Body.String() != "first:second:last" {
		t.Errorf("Body should be first:second:last but got %s", w.Body.String())
	}

	if r.Header().Get("Content-Type") != "text/plain" {
		t.Errorf("Content-Type should be text/plain but got %s", r.Header().Get("Content-Type"))
	}
}

func TestHTML(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)

	r.SetRenderer(renderer{})
	r.HTML(nil, "path/to/file")

	if w.Body.String() != b {
		t.Errorf("Body should be %s but got %s", b, w.Body.String())
	}

	if r.Header().Get("Content-Type") != "text/html" {
		t.Errorf("Content-Type should be text/html but got %s", r.Header().Get("Content-Type"))
	}
}

func TestHeaderAndStatusCode(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)

	r.WriteHeader(404)
	r.Header().Set("Custom-Attr", "test")

	if w.Code != 404 {
		t.Errorf("Status code should be 404 but got %d", w.Code)
	}

	if w.Header().Get("Custom-Attr") != "test" {
		t.Errorf("Header Custom-Attr should be test but got %s", w.Header().Get("Custom-Attr"))
	}
}

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)
	r.JSON(nil)

	if r.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type should be application/json but got %s", r.Header().Get("Content-Type"))
	}
}

func TestXML(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)
	r.XML(nil)

	if r.Header().Get("Content-Type") != "text/xml" {
		t.Errorf("Content-Type should be text/xml but got %s", r.Header().Get("Content-Type"))
	}
}

func TestSend(t *testing.T) {
	w := httptest.NewRecorder()
	r := newResponse(w, ctx)

	r.Text("test message")

	if w.Body.String() != "test message" {
		t.Errorf("Body should be test message but got %s", w.Body.String())
	}
}
