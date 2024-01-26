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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "github.com/vmware-tanzu/vm-operator-api/api/v1alpha1"
	vmopv1alpha1 "k8s.io/cloud-provider-vsphere/pkg/cloudprovider/vsphereparavirtual/vmop/applyconfiguration/vmop/v1alpha1"
)

// FakeVirtualMachineImages implements VirtualMachineImageInterface
type FakeVirtualMachineImages struct {
	Fake *FakeVmoperatorV1alpha1
}

var virtualmachineimagesResource = v1alpha1.SchemeGroupVersion.WithResource("virtualmachineimages")

var virtualmachineimagesKind = v1alpha1.SchemeGroupVersion.WithKind("VirtualMachineImage")

// Get takes name of the virtualMachineImage, and returns the corresponding virtualMachineImage object, and an error if there is any.
func (c *FakeVirtualMachineImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(virtualmachineimagesResource, name), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// List takes label and field selectors, and returns the list of VirtualMachineImages that match those selectors.
func (c *FakeVirtualMachineImages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualMachineImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(virtualmachineimagesResource, virtualmachineimagesKind, opts), &v1alpha1.VirtualMachineImageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineImageList{ListMeta: obj.(*v1alpha1.VirtualMachineImageList).ListMeta}
	for _, item := range obj.(*v1alpha1.VirtualMachineImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineImages.
func (c *FakeVirtualMachineImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(virtualmachineimagesResource, opts))
}

// Create takes the representation of a virtualMachineImage and creates it.  Returns the server's representation of the virtualMachineImage, and an error, if there is any.
func (c *FakeVirtualMachineImages) Create(ctx context.Context, virtualMachineImage *v1alpha1.VirtualMachineImage, opts v1.CreateOptions) (result *v1alpha1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(virtualmachineimagesResource, virtualMachineImage), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// Update takes the representation of a virtualMachineImage and updates it. Returns the server's representation of the virtualMachineImage, and an error, if there is any.
func (c *FakeVirtualMachineImages) Update(ctx context.Context, virtualMachineImage *v1alpha1.VirtualMachineImage, opts v1.UpdateOptions) (result *v1alpha1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(virtualmachineimagesResource, virtualMachineImage), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineImages) UpdateStatus(ctx context.Context, virtualMachineImage *v1alpha1.VirtualMachineImage, opts v1.UpdateOptions) (*v1alpha1.VirtualMachineImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(virtualmachineimagesResource, "status", virtualMachineImage), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// Delete takes name of the virtualMachineImage and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(virtualmachineimagesResource, name, opts), &v1alpha1.VirtualMachineImage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(virtualmachineimagesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineImageList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineImage.
func (c *FakeVirtualMachineImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachineimagesResource, name, pt, data, subresources...), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied virtualMachineImage.
func (c *FakeVirtualMachineImages) Apply(ctx context.Context, virtualMachineImage *vmopv1alpha1.VirtualMachineImageApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VirtualMachineImage, err error) {
	if virtualMachineImage == nil {
		return nil, fmt.Errorf("virtualMachineImage provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualMachineImage)
	if err != nil {
		return nil, err
	}
	name := virtualMachineImage.Name
	if name == nil {
		return nil, fmt.Errorf("virtualMachineImage.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachineimagesResource, *name, types.ApplyPatchType, data), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeVirtualMachineImages) ApplyStatus(ctx context.Context, virtualMachineImage *vmopv1alpha1.VirtualMachineImageApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VirtualMachineImage, err error) {
	if virtualMachineImage == nil {
		return nil, fmt.Errorf("virtualMachineImage provided to Apply must not be nil")
	}
	data, err := json.Marshal(virtualMachineImage)
	if err != nil {
		return nil, err
	}
	name := virtualMachineImage.Name
	if name == nil {
		return nil, fmt.Errorf("virtualMachineImage.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachineimagesResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.VirtualMachineImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachineImage), err
}
