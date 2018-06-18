/*
Copyright 2018 The Kubernetes Authors.

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
	akash_io_v1 "github.com/ovrclk/akash/pkg/apis/akash.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManifestCRDs implements ManifestCRDInterface
type FakeManifestCRDs struct {
	Fake *FakeExampleV1
	ns   string
}

var manifestcrdsResource = schema.GroupVersionResource{Group: "example.com", Version: "v1", Resource: "manifestcrds"}

var manifestcrdsKind = schema.GroupVersionKind{Group: "example.com", Version: "v1", Kind: "ManifestCRD"}

// Get takes name of the manifestCRD, and returns the corresponding manifestCRD object, and an error if there is any.
func (c *FakeManifestCRDs) Get(name string, options v1.GetOptions) (result *akash_io_v1.ManifestCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manifestcrdsResource, c.ns, name), &akash_io_v1.ManifestCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*akash_io_v1.ManifestCRD), err
}

// List takes label and field selectors, and returns the list of ManifestCRDs that match those selectors.
func (c *FakeManifestCRDs) List(opts v1.ListOptions) (result *akash_io_v1.ManifestCRDList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manifestcrdsResource, manifestcrdsKind, c.ns, opts), &akash_io_v1.ManifestCRDList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &akash_io_v1.ManifestCRDList{}
	for _, item := range obj.(*akash_io_v1.ManifestCRDList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifestCRDs.
func (c *FakeManifestCRDs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manifestcrdsResource, c.ns, opts))

}

// Create takes the representation of a manifestCRD and creates it.  Returns the server's representation of the manifestCRD, and an error, if there is any.
func (c *FakeManifestCRDs) Create(manifestCRD *akash_io_v1.ManifestCRD) (result *akash_io_v1.ManifestCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manifestcrdsResource, c.ns, manifestCRD), &akash_io_v1.ManifestCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*akash_io_v1.ManifestCRD), err
}

// Update takes the representation of a manifestCRD and updates it. Returns the server's representation of the manifestCRD, and an error, if there is any.
func (c *FakeManifestCRDs) Update(manifestCRD *akash_io_v1.ManifestCRD) (result *akash_io_v1.ManifestCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manifestcrdsResource, c.ns, manifestCRD), &akash_io_v1.ManifestCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*akash_io_v1.ManifestCRD), err
}

// Delete takes name of the manifestCRD and deletes it. Returns an error if one occurs.
func (c *FakeManifestCRDs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(manifestcrdsResource, c.ns, name), &akash_io_v1.ManifestCRD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifestCRDs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manifestcrdsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &akash_io_v1.ManifestCRDList{})
	return err
}

// Patch applies the patch and returns the patched manifestCRD.
func (c *FakeManifestCRDs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *akash_io_v1.ManifestCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manifestcrdsResource, c.ns, name, data, subresources...), &akash_io_v1.ManifestCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*akash_io_v1.ManifestCRD), err
}
