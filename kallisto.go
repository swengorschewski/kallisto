// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package kallisto is a web framework to help ease the development of web applications.
// The framework contains a router, a request scoped context and a reponse
// that will be dispatched after a request.
package kallisto

import (
	"net/http"
	"sync"
)

// Kallisto is the main struct of the web framework. It holds the router, all registered
// routes and route names and it can hold optional application scoped data.
type Kallisto struct {
	// Router is the router struct.
	*Router

	// routes stores all registered routes identified by its name.
	routes map[string]*Route

	sync.RWMutex // mutex for data
	// data holds all stored application scoped data identified by a key.
	data map[string]interface{}

	// services stores all registered services identieid by a key.
	services map[string]Runner
}

// New is the constructor method for a kallisto application.
func New() *Kallisto {
	k := &Kallisto{routes: make(map[string]*Route)}
	k.Router = NewRouter(k)
	k.data = make(map[string]interface{})
	k.services = make(map[string]Runner)
	return k
}

// Set stores the given key value pair.
// Anything stored via Set is application scoped.
func (k *Kallisto) Set(key string, value interface{}) {
	k.Lock()
	defer k.Unlock()
	k.data[key] = value
}

// Get returns an application scoped value identified by the given key.
func (k *Kallisto) Get(key string) interface{} {
	k.RLock()
	defer k.RUnlock()
	return k.data[key]
}

// SetService stores a given service identified by the given key.
func (k *Kallisto) SetService(key string, s Runner) {
	k.services[key] = s
}

// Service returns a service identified by the given key.
func (k *Kallisto) Service(key string) interface{} {
	return k.services[key]
}

// Route returns the path of the route identified by the given name.
func (k *Kallisto) Route(name string) string {
	if ok := k.routes[name]; ok != nil {
		return k.routes[name].Path
	}
	return ""
}

// Routes returns all registered routes.
func (k *Kallisto) Routes() map[string]*Route {
	return k.routes
}

// StartServices calls the run method in a seperate go routine of every
// registered service.
func (k *Kallisto) StartServices() {
	for _, service := range k.services {
		go service.Run()
	}
}

// ListenAndServe starts the server and listens for requests to the given url.
func (k *Kallisto) ListenAndServe(url string) {
	k.StartServices()

	http.ListenAndServe(url, k)
}
