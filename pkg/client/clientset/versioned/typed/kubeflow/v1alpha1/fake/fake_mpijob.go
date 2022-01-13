// Copyright 2018 The Kubeflow Authors.
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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubeflow/mpi-operator/pkg/apis/kubeflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMPIJobs implements MPIJobInterface
type FakeMPIJobs struct {
	Fake *FakeKubeflowV1alpha1
	ns   string
}

var mpijobsResource = schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1alpha1", Resource: "mpijobs"}

var mpijobsKind = schema.GroupVersionKind{Group: "kubeflow.org", Version: "v1alpha1", Kind: "MPIJob"}

// Get takes name of the mPIJob, and returns the corresponding mPIJob object, and an error if there is any.
func (c *FakeMPIJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.MPIJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mpijobsResource, c.ns, name), &v1alpha1.MPIJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MPIJob), err
}

// List takes label and field selectors, and returns the list of MPIJobs that match those selectors.
func (c *FakeMPIJobs) List(opts v1.ListOptions) (result *v1alpha1.MPIJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mpijobsResource, mpijobsKind, c.ns, opts), &v1alpha1.MPIJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MPIJobList{}
	for _, item := range obj.(*v1alpha1.MPIJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mPIJobs.
func (c *FakeMPIJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mpijobsResource, c.ns, opts))

}

// Create takes the representation of a mPIJob and creates it.  Returns the server's representation of the mPIJob, and an error, if there is any.
func (c *FakeMPIJobs) Create(mPIJob *v1alpha1.MPIJob) (result *v1alpha1.MPIJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mpijobsResource, c.ns, mPIJob), &v1alpha1.MPIJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MPIJob), err
}

// Update takes the representation of a mPIJob and updates it. Returns the server's representation of the mPIJob, and an error, if there is any.
func (c *FakeMPIJobs) Update(mPIJob *v1alpha1.MPIJob) (result *v1alpha1.MPIJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mpijobsResource, c.ns, mPIJob), &v1alpha1.MPIJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MPIJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMPIJobs) UpdateStatus(mPIJob *v1alpha1.MPIJob) (*v1alpha1.MPIJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mpijobsResource, "status", c.ns, mPIJob), &v1alpha1.MPIJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MPIJob), err
}

// Delete takes name of the mPIJob and deletes it. Returns an error if one occurs.
func (c *FakeMPIJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mpijobsResource, c.ns, name), &v1alpha1.MPIJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMPIJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mpijobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MPIJobList{})
	return err
}

// Patch applies the patch and returns the patched mPIJob.
func (c *FakeMPIJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MPIJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mpijobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MPIJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MPIJob), err
}
