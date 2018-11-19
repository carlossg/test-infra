/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"k8s.io/client-go/rest"
	"testing"

	"k8s.io/test-infra/prow/plugins"
)

// Make sure that our plugins are valid.
func TestPlugins(t *testing.T) {
	pa := &plugins.PluginAgent{}
	if err := pa.Load("../../plugins.yaml"); err != nil {
		t.Fatalf("Could not load plugins: %v.", err)
	}
}
func TestFoo(t *testing.T) {
	configs := map[string]rest.Config{}
	//var defCtx *string
	//defCtx = new(string)

	configs["default"] = rest.Config{Username: "foo"}
	context := "default"

	_, ok := configs[context]
	if !ok {
		t.Errorf("context not found: [%s]", context)
	}

	//logrus.Info("ok")
	//configs := map[string]rest.Config{}
	//var defCtx *string
	//
	//defCtx = new(string)
	//
	//configs[*defCtx] = *localCfg
}
