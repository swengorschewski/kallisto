// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import "io"

// Renderer is an interface to make it easier to replace the templating engine.
type Renderer interface {
	// Render turns the templates and the given data to valid HTML.
	Render(io.Writer, interface{}, []string)
}
