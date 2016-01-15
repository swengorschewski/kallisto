// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

// A Session is a store to preserve the state of a HTTP connection.
type Session interface {
	// Set stores the given key value pair.
	Set(k string, v interface{})

	// Get returns a value indentified by the given key.
	Get(k string) interface{}

	// SetFlash stores the given key value pair.
	// The stored key value pair is only available for the next request.
	SetFlash(k string, v interface{})

	// Flash returns stored value identified by the given key.
	Flash(k string) interface{}
}
