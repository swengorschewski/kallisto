// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

// ServeStatic returns a file.
func ServeStatic(path string) ControllerFunc {
	return func(c *Context, r *Response) {
		r.Static(path, c.Request)
	}
}
