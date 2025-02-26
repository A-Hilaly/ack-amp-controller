// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AlertManagerDefinitionSpec defines the desired state of AlertManagerDefinition.
type AlertManagerDefinitionSpec struct {
	// The ID of the workspace in which to create the alert manager definition.
	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceID"`

	// +kubebuilder:validation:Required
	Configuration *string `json:"configuration"`
}

// AlertManagerDefinitionStatus defines the observed state of AlertManagerDefinition
type AlertManagerDefinitionStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Status code of this definition.
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty"`
	// The reason for failure if any.
	// +kubebuilder:validation:Optional
	StatusReason *string `json:"statusReason,omitempty"`
}

// AlertManagerDefinition is the Schema for the AlertManagerDefinitions API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="WORKSPACE-ID",type=string,priority=0,JSONPath=`.spec.workspaceID`
type AlertManagerDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlertManagerDefinitionSpec   `json:"spec,omitempty"`
	Status            AlertManagerDefinitionStatus `json:"status,omitempty"`
}

// AlertManagerDefinitionList contains a list of AlertManagerDefinition
// +kubebuilder:object:root=true
type AlertManagerDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertManagerDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlertManagerDefinition{}, &AlertManagerDefinitionList{})
}
