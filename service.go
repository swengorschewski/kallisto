// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

// A Runner is a singleton that will be started when the application starts.
//
// For communication it should use channels.
type Runner interface {
	Run()
}
