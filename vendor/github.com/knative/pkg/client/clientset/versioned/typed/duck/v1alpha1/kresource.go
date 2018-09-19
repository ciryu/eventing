/*
Copyright 2018 The Knative Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	scheme "github.com/knative/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KResourcesGetter has a method to return a KResourceInterface.
// A group's client should implement this interface.
type KResourcesGetter interface {
	KResources(namespace string) KResourceInterface
}

// KResourceInterface has methods to work with KResource resources.
type KResourceInterface interface {
	Create(*v1alpha1.KResource) (*v1alpha1.KResource, error)
	Update(*v1alpha1.KResource) (*v1alpha1.KResource, error)
	UpdateStatus(*v1alpha1.KResource) (*v1alpha1.KResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KResource, error)
	List(opts v1.ListOptions) (*v1alpha1.KResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KResource, err error)
	KResourceExpansion
}

// kResources implements KResourceInterface
type kResources struct {
	client rest.Interface
	ns     string
}

// newKResources returns a KResources
func newKResources(c *DuckV1alpha1Client, namespace string) *kResources {
	return &kResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kResource, and returns the corresponding kResource object, and an error if there is any.
func (c *kResources) Get(name string, options v1.GetOptions) (result *v1alpha1.KResource, err error) {
	result = &v1alpha1.KResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KResources that match those selectors.
func (c *kResources) List(opts v1.ListOptions) (result *v1alpha1.KResourceList, err error) {
	result = &v1alpha1.KResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kResources.
func (c *kResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kResource and creates it.  Returns the server's representation of the kResource, and an error, if there is any.
func (c *kResources) Create(kResource *v1alpha1.KResource) (result *v1alpha1.KResource, err error) {
	result = &v1alpha1.KResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kresources").
		Body(kResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kResource and updates it. Returns the server's representation of the kResource, and an error, if there is any.
func (c *kResources) Update(kResource *v1alpha1.KResource) (result *v1alpha1.KResource, err error) {
	result = &v1alpha1.KResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kresources").
		Name(kResource.Name).
		Body(kResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kResources) UpdateStatus(kResource *v1alpha1.KResource) (result *v1alpha1.KResource, err error) {
	result = &v1alpha1.KResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kresources").
		Name(kResource.Name).
		SubResource("status").
		Body(kResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the kResource and deletes it. Returns an error if one occurs.
func (c *kResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kResource.
func (c *kResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KResource, err error) {
	result = &v1alpha1.KResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}