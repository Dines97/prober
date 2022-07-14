/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"prober.io/prober/probes"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Severity string

// ProbeSpec defines the desired state of Probe
type ProbeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Optional probe description field
	// +optional
	Description string `json:"description,omitempty"`

	// Tags to filter probes
	// +optional
	Tags []string `json:"tags,omitempty"`

	// Template of probe
	RabbitMQ probes.RabbitMQ `json:"rabbitmq,omitempty"`

	Resolution probes.Resolution `json:"resolution,omitempty"`

	//Probe probes.Probe `json:"probe"`
}

// ProbeStatus defines the observed state of Probe
type ProbeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:default:=Unknown
	RunResult probes.RunResult `json:"runResult"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Probe is the Schema for the probes API
// +kubebuilder:printcolumn:name="Last run",type=string,JSONPath=`.status.runResult`
type Probe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProbeSpec `json:"spec,omitempty"`

	// +kubebuilder:default={runResult: Unknown}
	Status ProbeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProbeList contains a list of Probe
type ProbeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Probe `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Probe{}, &ProbeList{})
}
