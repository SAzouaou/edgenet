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

// FakeUsers implements UserInterface
type FakeUsers struct {
	Fake *FakeAppsV1alpha
	ns   string
}

var usersResource = schema.GroupVersionResource{Group: "apps.edgenet.io", Version: "v1alpha", Resource: "users"}

var usersKind = schema.GroupVersionKind{Group: "apps.edgenet.io", Version: "v1alpha", Kind: "User"}

// Get takes name of the user, and returns the corresponding user object, and an error if there is any.
func (c *FakeUsers) Get(name string, options v1.GetOptions) (result *v1alpha.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usersResource, c.ns, name), &v1alpha.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.User), err
}

// List takes label and field selectors, and returns the list of Users that match those selectors.
func (c *FakeUsers) List(opts v1.ListOptions) (result *v1alpha.UserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usersResource, usersKind, c.ns, opts), &v1alpha.UserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.UserList{ListMeta: obj.(*v1alpha.UserList).ListMeta}
	for _, item := range obj.(*v1alpha.UserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested users.
func (c *FakeUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usersResource, c.ns, opts))

}

// Create takes the representation of a user and creates it.  Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Create(user *v1alpha.User) (result *v1alpha.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usersResource, c.ns, user), &v1alpha.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.User), err
}

// Update takes the representation of a user and updates it. Returns the server's representation of the user, and an error, if there is any.
func (c *FakeUsers) Update(user *v1alpha.User) (result *v1alpha.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usersResource, c.ns, user), &v1alpha.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.User), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUsers) UpdateStatus(user *v1alpha.User) (*v1alpha.User, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(usersResource, "status", c.ns, user), &v1alpha.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.User), err
}

// Delete takes name of the user and deletes it. Returns an error if one occurs.
func (c *FakeUsers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(usersResource, c.ns, name), &v1alpha.User{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.UserList{})
	return err
}

// Patch applies the patch and returns the patched user.
func (c *FakeUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.User, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usersResource, c.ns, name, pt, data, subresources...), &v1alpha.User{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.User), err
}
