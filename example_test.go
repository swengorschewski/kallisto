// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto_test

import "github.com/swengorschewski/kallisto"

func Example() {
	// create a kallisto mux
	k := kallisto.New()

	// register a route for a HTTP GET request
	// the first parameter represents the request url,
	// the second parameter is an internal name for the route,
	// and the last parameter is a controller function that will be called
	// if a request for this route comes in.
	k.GET("/", "index", func(ctx *kallisto.Context, res *kallisto.Response) {
		res.Text("Hello world!")
	})

	// starts a webserver listening for requests to localhost on port 8080
	k.ListenAndServe("localhost:8080")
}
