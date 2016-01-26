# kallisto
 Kallisto a web framework based on Julien Schmidts [httprouter](https://github.com/julienschmidt/httprouter) package and inspired by [martini](https://github.com/go-martini/martini), [revel](https://github.com/revel/revel) and [gin](https://github.com/gin-gonic/gin).
 
 It was written for a university project and is not (and propablly never will be) ready for use in production environments.
 
 [Documentation] (https://godoc.org/github.com/swengorschewski/kallisto)
 
 Example:
 ```go
package main

import "gitlab.com/swen/kallisto"

func main() {
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
```
