// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"reflect"
	"testing"
)

var m1 = func(c *Context) {}
var m2 = func(c *Context) {}

func TestBefore(t *testing.T) {
	before := MiddlewareChain{m1, m2}

	r := &Route{}
	r.SetBefore(m1, m2)

	if !compareMiddlewareFunc(r.Before, before) {
		t.Error("Before middlewares do not match.")
	}
}

func TestAfter(t *testing.T) {
	after := MiddlewareChain{m1, m2}

	r := &Route{}
	r.SetAfter(m1, m2)

	if !compareMiddlewareFunc(r.After, after) {
		t.Error("After middlewares do not match.")
	}
}

func compareMiddlewareFunc(m1 MiddlewareChain, m2 MiddlewareChain) bool {
	if m1 == nil && m2 == nil {
		return true
	}

	if m1 == nil || m2 == nil {
		return false
	}

	if len(m1) != len(m2) {
		return false
	}

	for i := range m1 {
		if reflect.ValueOf(m1[i]).Pointer() != reflect.ValueOf(m2[i]).Pointer() {
			return false
		}
	}

	return true
}
