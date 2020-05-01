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

package v1alpha

import (
	v1alpha "edgenet/pkg/apis/apps/v1alpha"
	scheme "edgenet/pkg/client/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoginsGetter has a method to return a LoginInterface.
// A group's client should implement this interface.
type LoginsGetter interface {
	Logins(namespace string) LoginInterface
}

// LoginInterface has methods to work with Login resources.
type LoginInterface interface {
	Create(*v1alpha.Login) (*v1alpha.Login, error)
	Update(*v1alpha.Login) (*v1alpha.Login, error)
	UpdateStatus(*v1alpha.Login) (*v1alpha.Login, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha.Login, error)
	List(opts v1.ListOptions) (*v1alpha.LoginList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.Login, err error)
	LoginExpansion
}

// logins implements LoginInterface
type logins struct {
	client rest.Interface
	ns     string
}

// newLogins returns a Logins
func newLogins(c *AppsV1alphaClient, namespace string) *logins {
	return &logins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the login, and returns the corresponding login object, and an error if there is any.
func (c *logins) Get(name string, options v1.GetOptions) (result *v1alpha.Login, err error) {
	result = &v1alpha.Login{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("logins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Logins that match those selectors.
func (c *logins) List(opts v1.ListOptions) (result *v1alpha.LoginList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha.LoginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("logins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logins.
func (c *logins) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("logins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a login and creates it.  Returns the server's representation of the login, and an error, if there is any.
func (c *logins) Create(login *v1alpha.Login) (result *v1alpha.Login, err error) {
	result = &v1alpha.Login{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("logins").
		Body(login).
		Do().
		Into(result)
	return
}

// Update takes the representation of a login and updates it. Returns the server's representation of the login, and an error, if there is any.
func (c *logins) Update(login *v1alpha.Login) (result *v1alpha.Login, err error) {
	result = &v1alpha.Login{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("logins").
		Name(login.Name).
		Body(login).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *logins) UpdateStatus(login *v1alpha.Login) (result *v1alpha.Login, err error) {
	result = &v1alpha.Login{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("logins").
		Name(login.Name).
		SubResource("status").
		Body(login).
		Do().
		Into(result)
	return
}

// Delete takes name of the login and deletes it. Returns an error if one occurs.
func (c *logins) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("logins").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *logins) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("logins").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched login.
func (c *logins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.Login, err error) {
	result = &v1alpha.Login{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("logins").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
