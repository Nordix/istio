// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import "istio.io/istio/pkg/cluster"

// Origin of a resource. This is source-implementation dependent.
type Origin interface {
	FriendlyName() string
	Namespace() Namespace
	Reference() Reference

	// FieldMap returns the flat map containing paths of the fields in the resource as keys,
	// and their corresponding line numbers as values
	FieldMap() map[string]int
	Comparator() string

	// ClusterName returns the cluster name where the resource is located
	ClusterName() cluster.ID
}

// Reference provides more information about an Origin. This is also source-implementation dependent.
type Reference interface {
	String() string
}
