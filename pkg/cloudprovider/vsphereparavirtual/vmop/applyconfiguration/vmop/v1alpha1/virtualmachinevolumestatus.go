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

// VirtualMachineVolumeStatusApplyConfiguration represents an declarative configuration of the VirtualMachineVolumeStatus type for use
// with apply.
type VirtualMachineVolumeStatusApplyConfiguration struct {
	Name     *string `json:"name,omitempty"`
	Attached *bool   `json:"attached,omitempty"`
	DiskUuid *string `json:"diskUUID,omitempty"`
	Error    *string `json:"error,omitempty"`
}

// VirtualMachineVolumeStatusApplyConfiguration constructs an declarative configuration of the VirtualMachineVolumeStatus type for use with
// apply.
func VirtualMachineVolumeStatus() *VirtualMachineVolumeStatusApplyConfiguration {
	return &VirtualMachineVolumeStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *VirtualMachineVolumeStatusApplyConfiguration) WithName(value string) *VirtualMachineVolumeStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithAttached sets the Attached field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attached field is set to the value of the last call.
func (b *VirtualMachineVolumeStatusApplyConfiguration) WithAttached(value bool) *VirtualMachineVolumeStatusApplyConfiguration {
	b.Attached = &value
	return b
}

// WithDiskUuid sets the DiskUuid field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DiskUuid field is set to the value of the last call.
func (b *VirtualMachineVolumeStatusApplyConfiguration) WithDiskUuid(value string) *VirtualMachineVolumeStatusApplyConfiguration {
	b.DiskUuid = &value
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *VirtualMachineVolumeStatusApplyConfiguration) WithError(value string) *VirtualMachineVolumeStatusApplyConfiguration {
	b.Error = &value
	return b
}
