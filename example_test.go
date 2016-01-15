// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto_test

import "gitlab.com/swen/kallisto"

func ExampleSimpleServer() {
	// create a kallisto mux
	k := kallisto.New()

	// register a route for a HTTP GET request
	// the first parameter represents the request url
	// the second parameter is a name an internal name for this route
	// and the last parameter is a controller function that will be called
	// if a request for this route comes in.
	k.GET("/", "index", func(ctx *kallisto.Context, res *kallisto.Response) {
		res.Text("Hello world!")
	})

	// starts a webserver listening for requests to localhost on port 8080
	k.ListenAndServe("localhost:8080")
}
