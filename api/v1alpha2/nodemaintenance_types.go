/*
Copyright 2021.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodemaintenanceSpec defines the desired state of Nodemaintenance
type NodemaintenanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Hostname is the hostname of the nm
	Hostname string `json:"hostname,omitempty"`
	// MacAddress is the mac address of the nm
	MacAddress string `json:"macAddress,omitempty"`
	// UniqueId  is the uniqueid of the nm
	UniqueId string `json:"uniqueId,omitempty"`
	// services is the service that the nm have
	Services []Service `json:"services,omitempty"`
	// gateway is the endpoint to access the service on the node  can be fqdn or ip address
	Gateway string `json:"gateway"`
}

// NodemaintenanceStatus defines the observed state of Nodemaintenance
type NodemaintenanceStatus struct {
	// nodeStatus reflects the current node status
	NodeStatus    NodeStatus      `json:"nodeStatus,omitempty"`
	ServiceStatus []ServiceStatus `json:"serviceStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster

// Nodemaintenance is the Schema for the nodemaintenances API
type Nodemaintenance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodemaintenanceSpec   `json:"spec,omitempty"`
	Status NodemaintenanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodemaintenanceList contains a list of Nodemaintenance
type NodemaintenanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nodemaintenance `json:"items"`
}

type Service struct {
	// name  is the name of the nm service
	Name string `json:"name"`
	// type is the type of the service: tcp udp  http https stcp sudp
	Type string `json:"type"`
	// meta contains the metadata of the service uploaded from the edgeside
	Meta map[string]string `json:"meta,omitempty"`
}
type ServiceStatus struct {
	// name  is the name of the nm service
	Name string `json:"name"`
	// type is the type of the service: tcp udp  http https stcp sudp
	Type string `json:"type"`
	// port indicates the port exposes on the gateway
	Port int `json:"port,omitempty"`
	// // lastUpdateTimestamp indicates the last time  heart beat sent from the edge
	// LastUpdateTimestamp string `json:"lastUpdateTimestamp,omitempty"`
	// // Active indicates if the service is useable
	// Active bool `json:"active,omitempty"`
}

type NodeStatus struct {
	// maintainable indicates if the node sends heartbeat to the cloud in time
	Maintainable bool `json:"maintainable"`
	// lastUpdateTimestamp indicates the last time  heart beat sent from the edge
	LastUpdateTimestamp string `json:"lastUpdateTimestamp,omitempty"`
	// firstLoginTimestamp indicates the first time the node is online
	FirstLoginTimestamp string `json:"firstLoginTimestamp,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Nodemaintenance{}, &NodemaintenanceList{})
}
