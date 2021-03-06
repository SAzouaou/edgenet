/*
Copyright 2019 Sorbonne Université

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

// This nodelabeler feature includes GeoLite2 data created by MaxMind, available from
// https://www.maxmind.com.

package main

import (
	"edgenet/pkg/authorization"
	"edgenet/pkg/controller/v1/nodelabeler"
)

func main() {
	// Set kubeconfig to be used to create clientsets
	authorization.SetKubeConfig()
	// Start the controller to watch nodes and attach the labels to them
	nodelabeler.Start()
}
