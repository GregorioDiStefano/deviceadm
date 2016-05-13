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

type DeviceID string

// Device wrapper
type Device struct {
	//system-generated device ID
	ID DeviceID `json:"id"`

	//blob of encrypted identity attributes
	DeviceIdentity string `json:"device_identity"`

	//device's public key
	Key string `json:"key"`

	//admission status('accepted', 'rejected', 'pending')
	Status string `json:"status"`

	//decoded, human-readable identity attribute set
	Attributes map[string]string `json:"attributes"`
}

func (did DeviceID) String() string {
	return string(did)
}