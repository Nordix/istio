// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package nodeagent

import (
	"os"
	"testing"

	pconstants "istio.io/istio/cni/pkg/constants"
)

func TestMain(m *testing.M) {
	// Override host NS path globally for all tests
	originalHostNetNSPath := pconstants.HostNetNSPath
	defer func() { pconstants.HostNetNSPath = originalHostNetNSPath }()
	pconstants.HostNetNSPath = "/proc/self/ns/net"

	// Run tests
	os.Exit(m.Run())
}
