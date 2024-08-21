/*
Copyright 2024 The Kanister Authors.

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

import (
	v1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
)

// ActionSetStatusApplyConfiguration represents an declarative configuration of the ActionSetStatus type for use
// with apply.
type ActionSetStatusApplyConfiguration struct {
	State    *v1alpha1.State                   `json:"state,omitempty"`
	Actions  []ActionStatusApplyConfiguration  `json:"actions,omitempty"`
	Error    *ErrorApplyConfiguration          `json:"error,omitempty"`
	Progress *ActionProgressApplyConfiguration `json:"progress,omitempty"`
}

// ActionSetStatusApplyConfiguration constructs an declarative configuration of the ActionSetStatus type for use with
// apply.
func ActionSetStatus() *ActionSetStatusApplyConfiguration {
	return &ActionSetStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *ActionSetStatusApplyConfiguration) WithState(value v1alpha1.State) *ActionSetStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithActions adds the given value to the Actions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Actions field.
func (b *ActionSetStatusApplyConfiguration) WithActions(values ...*ActionStatusApplyConfiguration) *ActionSetStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithActions")
		}
		b.Actions = append(b.Actions, *values[i])
	}
	return b
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *ActionSetStatusApplyConfiguration) WithError(value *ErrorApplyConfiguration) *ActionSetStatusApplyConfiguration {
	b.Error = value
	return b
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *ActionSetStatusApplyConfiguration) WithProgress(value *ActionProgressApplyConfiguration) *ActionSetStatusApplyConfiguration {
	b.Progress = value
	return b
}