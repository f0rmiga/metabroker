/*
Copyright 2020 SUSE

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

	v1alpha1 "github.com/SUSE/metabroker/apis/servicebroker/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCredentials implements CredentialInterface
type FakeCredentials struct {
	Fake *FakeServicebrokerV1alpha1
	ns   string
}

var credentialsResource = schema.GroupVersionResource{Group: "servicebroker", Version: "v1alpha1", Resource: "credentials"}

var credentialsKind = schema.GroupVersionKind{Group: "servicebroker", Version: "v1alpha1", Kind: "Credential"}

// Get takes name of the credential, and returns the corresponding credential object, and an error if there is any.
func (c *FakeCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(credentialsResource, c.ns, name), &v1alpha1.Credential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// List takes label and field selectors, and returns the list of Credentials that match those selectors.
func (c *FakeCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(credentialsResource, credentialsKind, c.ns, opts), &v1alpha1.CredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CredentialList{ListMeta: obj.(*v1alpha1.CredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.CredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested credentials.
func (c *FakeCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(credentialsResource, c.ns, opts))

}

// Create takes the representation of a credential and creates it.  Returns the server's representation of the credential, and an error, if there is any.
func (c *FakeCredentials) Create(ctx context.Context, credential *v1alpha1.Credential, opts v1.CreateOptions) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(credentialsResource, c.ns, credential), &v1alpha1.Credential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// Update takes the representation of a credential and updates it. Returns the server's representation of the credential, and an error, if there is any.
func (c *FakeCredentials) Update(ctx context.Context, credential *v1alpha1.Credential, opts v1.UpdateOptions) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(credentialsResource, c.ns, credential), &v1alpha1.Credential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCredentials) UpdateStatus(ctx context.Context, credential *v1alpha1.Credential, opts v1.UpdateOptions) (*v1alpha1.Credential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(credentialsResource, "status", c.ns, credential), &v1alpha1.Credential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}

// Delete takes name of the credential and deletes it. Returns an error if one occurs.
func (c *FakeCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(credentialsResource, c.ns, name), &v1alpha1.Credential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(credentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CredentialList{})
	return err
}

// Patch applies the patch and returns the patched credential.
func (c *FakeCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Credential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(credentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Credential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Credential), err
}
