// Copyright 2016 Swen Gorschewski. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package kallisto

import (
	"reflect"
	"testing"
)

// Runner mock to test the Service and StartService methods.
type service struct {
	running bool
	c       chan bool
}

func (s *service) Run() {
	s.running = true
	s.c <- true
}

func TestGetAndSet(t *testing.T) {
	k := New()

	testData := New()

	k.Set("test", testData)

	savedTestData := k.Get("test")

	if testData != savedTestData {
		t.Error("savedTestData is not the same as testData")
	}
}

func TestService(t *testing.T) {
	k := New()
	s := &service{}

	k.SetService("service", s)

	service := k.Service("service")

	if s != service {
		t.Error("r and testRunner are not equal")
	}
}

func TestRoutes(t *testing.T) {
	routes := make(map[string]*Route)
	routes["test"] = &Route{Path: "/test"}
	routes["test2"] = &Route{Path: "/test2"}

	k := &Kallisto{routes: routes}

	if k.Route("test") != "/test" {
		t.Errorf("Route path should be /test but got %s", k.Route("test"))
	}

	if k.Route("not present") != "" {
		t.Errorf("Route path should be empty string but got %s", k.Route("not present"))
	}

	if !reflect.DeepEqual(k.Routes(), routes) {
		t.Error("Routes are not equal.")
	}
}

func TestStartServices(t *testing.T) {
	k := New()
	c := make(chan bool)
	s := &service{c: c}
	k.SetService("test", s)
	k.StartServices()

	<-c

	if !s.running {
		t.Errorf("Running should be true but got %t", s.running)
	}
}

func TestListenAndServe(t *testing.T) {

}
