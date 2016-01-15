// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUse(t *testing.T) {
	r := &Router{middlewares: make(MiddlewareChain, 0)}

	m1 := func(c *Context) {}
	m2 := func(c *Context) {}
	middlewares := MiddlewareChain{m1, m2}

	r.Use(m1, m2)

	if !compareMiddlewareFunc(r.middlewares, middlewares) {
		t.Error("Middlewares do not match.")
	}
}

func TestHTTPMethods(t *testing.T) {
	k := New()

	s := "test"
	c := func(c *Context, r *Response) { r.Text(s) }
	cp := func(c *Context, r *Response) { r.Text(c.Request.FormValue("test-value")) }

	k.GET("/get", "get", c)
	r, _ := http.NewRequest("GET", "/get", nil)
	check(k, r, t, s)

	k.POST("/post", "post", cp)
	r, _ = http.NewRequest("POST", "/post", nil)
	r.ParseForm()
	r.Form.Add("test-value", s)
	check(k, r, t, s)

	k.PATCH("/patch", "patch", cp)
	r, _ = http.NewRequest("PATCH", "/patch", nil)
	r.ParseForm()
	r.Form.Add("test-value", s)
	check(k, r, t, s)

	k.PUT("/put", "put", cp)
	r, _ = http.NewRequest("PUT", "/put", nil)
	r.ParseForm()
	r.Form.Add("test-value", s)
	check(k, r, t, s)

	k.DELETE("/delete", "delete", c)
	r, _ = http.NewRequest("DELETE", "/delete", nil)
	check(k, r, t, s)

	k.HEAD("/head", "head", c)
	r, _ = http.NewRequest("HEAD", "/head", nil)
	check(k, r, t, s)

	k.OPTIONS("/options", "options", c)
	r, _ = http.NewRequest("OPTIONS", "/options", nil)
	check(k, r, t, s)
}

func TestNotFound(t *testing.T) {
	k := New()
	s := "Not Found"
	c := func(c *Context, r *Response) { r.Text(s) }

	k.NotFound(c)
	r, _ := http.NewRequest("GET", "/notfound", nil)
	w := httptest.NewRecorder()
	k.ServeHTTP(w, r)

	if w.Code != 404 {
		t.Errorf("Status should be 404 but got %d", w.Code)
	}

	if w.Body.String() != s {
		t.Errorf("Body should be %s but got %s", s, w.Body.String())
	}
}

func TestPanicHandler(t *testing.T) {
	k := New()
	s := "Panic!"
	c := func(c *Context, r *Response) {
		panic("stop here")
		r.Text("should not be reached")
	}
	ph := func(c *Context, r *Response) {
		r.Text(s)
	}

	k.PanicHandler(ph)

	k.GET("/", "home", c)
	r, _ := http.NewRequest("GET", "/", nil)
	check(k, r, t, s)
}

func TestGroup(t *testing.T) {
	k := New()
	s := "first"
	m1 := func(c *Context) { s += ":second:" }
	m2 := func(c *Context) { s += "last" }
	c := func(c *Context, r *Response) { r.Text(s) }
	k.Group("/group", "group::", func(r *Router) {
		r.GET("get", "get", c)
	}, m1, m2)

	r, _ := http.NewRequest("GET", "/group/get", nil)
	check(k, r, t, "first:second:last")
}

type mockFS struct {
	opened bool
}

func (m *mockFS) Open(path string) (http.File, error) {
	m.opened = true
	return nil, errors.New("this is just a mock")
}

func TestStatic(t *testing.T) {
	k := New()
	fs := &mockFS{}
	k.ServeStatic("/public/*filepath", fs)
	r, _ := http.NewRequest("GET", "/public/file.extension", nil)
	w := httptest.NewRecorder()

	k.ServeHTTP(w, r)

	if !fs.opened {
		t.Error("Mock file system should have called open.")
	}
}

func check(k *Kallisto, r *http.Request, t *testing.T, s string) {
	w := httptest.NewRecorder()
	k.ServeHTTP(w, r)

	if w.Body.String() != s {
		t.Errorf("Body should be %s but got %s", s, w.Body.String())
	}
}
