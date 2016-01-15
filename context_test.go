// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestParam(t *testing.T) {
	c := &Context{Params: httprouter.Params{httprouter.Param{Key: "key", Value: "value"}}}

	if c.Param("key") != "value" {
		t.Errorf("Param should be value but got %s", c.Param("key"))
	}
}

func TestNext(t *testing.T) {

	s := "untouched"

	m1 := func(c *Context) {
		s = "first"
	}

	c := func(c *Context, r *Response) {
		s = s + ":second"
	}

	m2 := func(c *Context) {
		s = s + ":last"
	}

	route := &Route{}
	route.SetBefore(m1)
	route.Controller = c
	route.SetAfter(m2)

	ctx := &Context{
		middlewareIndex: -1,
		route:           route,
	}

	ctx.Next()

	if s != "first:second:last" {
		t.Errorf("Result should be first:second:last but got %s", s)
	}
}

func TestApp(t *testing.T) {
	k := New()
	c := &Context{
		kallisto: k,
	}

	if c.App() != k {
		t.Error("k and App() are not the same instance of kallisto.")
	}
}

type session struct{}

func (s session) Get(k string) interface{}         { return nil }
func (s session) Set(k string, v interface{})      {}
func (s session) Flash(k string) interface{}       { return nil }
func (s session) SetFlash(k string, v interface{}) {}

func TestSession(t *testing.T) {
	s := session{}
	c := &Context{}
	c.Session = s

	if s != c.Session {
		t.Error("s and Session() are not the same instance of session.")
	}
}

func TestGetSet(t *testing.T) {
	c := &Context{data: make(map[string]interface{})}
	c.Set("key", "value")

	if c.Get("key") != "value" {
		t.Errorf("Value should be value but got %s", c.Get("key"))
	}
}
