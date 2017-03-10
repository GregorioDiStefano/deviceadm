// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupApi(t *testing.T) {
	// expecting an error
	api, err := SetupAPI("foo")
	assert.Nil(t, api)
	assert.Error(t, err)

	api, err = SetupAPI(EnvDev)
	assert.NotNil(t, api)
	assert.Nil(t, err)
}

func TestSetupDataStore(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestSetupDataStore in short mode.")
	}

	d, err := SetupDataStore("::invalid-url::")
	assert.Nil(t, d)
	assert.Error(t, err)
}
