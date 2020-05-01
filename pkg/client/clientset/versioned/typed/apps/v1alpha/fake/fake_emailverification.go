/*
Copyright The Kubernetes Authors.

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
	v1alpha "edgenet/pkg/apis/apps/v1alpha"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEmailVerifications implements EmailVerificationInterface
type FakeEmailVerifications struct {
	Fake *FakeAppsV1alpha
	ns   string
}

var emailverificationsResource = schema.GroupVersionResource{Group: "apps.edgenet.io", Version: "v1alpha", Resource: "emailverifications"}

var emailverificationsKind = schema.GroupVersionKind{Group: "apps.edgenet.io", Version: "v1alpha", Kind: "EmailVerification"}

// Get takes name of the emailVerification, and returns the corresponding emailVerification object, and an error if there is any.
func (c *FakeEmailVerifications) Get(name string, options v1.GetOptions) (result *v1alpha.EmailVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(emailverificationsResource, c.ns, name), &v1alpha.EmailVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EmailVerification), err
}

// List takes label and field selectors, and returns the list of EmailVerifications that match those selectors.
func (c *FakeEmailVerifications) List(opts v1.ListOptions) (result *v1alpha.EmailVerificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(emailverificationsResource, emailverificationsKind, c.ns, opts), &v1alpha.EmailVerificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.EmailVerificationList{ListMeta: obj.(*v1alpha.EmailVerificationList).ListMeta}
	for _, item := range obj.(*v1alpha.EmailVerificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested emailVerifications.
func (c *FakeEmailVerifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(emailverificationsResource, c.ns, opts))

}

// Create takes the representation of a emailVerification and creates it.  Returns the server's representation of the emailVerification, and an error, if there is any.
func (c *FakeEmailVerifications) Create(emailVerification *v1alpha.EmailVerification) (result *v1alpha.EmailVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(emailverificationsResource, c.ns, emailVerification), &v1alpha.EmailVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EmailVerification), err
}

// Update takes the representation of a emailVerification and updates it. Returns the server's representation of the emailVerification, and an error, if there is any.
func (c *FakeEmailVerifications) Update(emailVerification *v1alpha.EmailVerification) (result *v1alpha.EmailVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(emailverificationsResource, c.ns, emailVerification), &v1alpha.EmailVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EmailVerification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEmailVerifications) UpdateStatus(emailVerification *v1alpha.EmailVerification) (*v1alpha.EmailVerification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(emailverificationsResource, "status", c.ns, emailVerification), &v1alpha.EmailVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EmailVerification), err
}

// Delete takes name of the emailVerification and deletes it. Returns an error if one occurs.
func (c *FakeEmailVerifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(emailverificationsResource, c.ns, name), &v1alpha.EmailVerification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEmailVerifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(emailverificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.EmailVerificationList{})
	return err
}

// Patch applies the patch and returns the patched emailVerification.
func (c *FakeEmailVerifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.EmailVerification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(emailverificationsResource, c.ns, name, pt, data, subresources...), &v1alpha.EmailVerification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.EmailVerification), err
}
