/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// VirtualMachineClassPoliciesApplyConfiguration represents an declarative configuration of the VirtualMachineClassPolicies type for use
// with apply.
type VirtualMachineClassPoliciesApplyConfiguration struct {
	Resources *VirtualMachineClassResourcesApplyConfiguration `json:"resources,omitempty"`
}

// VirtualMachineClassPoliciesApplyConfiguration constructs an declarative configuration of the VirtualMachineClassPolicies type for use with
// apply.
func VirtualMachineClassPolicies() *VirtualMachineClassPoliciesApplyConfiguration {
	return &VirtualMachineClassPoliciesApplyConfiguration{}
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *VirtualMachineClassPoliciesApplyConfiguration) WithResources(value *VirtualMachineClassResourcesApplyConfiguration) *VirtualMachineClassPoliciesApplyConfiguration {
	b.Resources = value
	return b
}
