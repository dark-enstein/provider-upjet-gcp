// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DomainMappingInitParameters struct {

	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is STRICT.
	// Possible values are: STRICT, OVERRIDE.
	OverrideStrategy *string `json:"overrideStrategy,omitempty" tf:"override_strategy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SSLSettings *SSLSettingsInitParameters `json:"sslSettings,omitempty" tf:"ssl_settings,omitempty"`
}

type DomainMappingObservation struct {

	// an identifier for the resource with format apps/{{project}}/domainMappings/{{domain_name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is STRICT.
	// Possible values are: STRICT, OVERRIDE.
	OverrideStrategy *string `json:"overrideStrategy,omitempty" tf:"override_strategy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The resource records required to configure this domain mapping. These records must be added to the domain's DNS
	// configuration in order to serve the application via this domain mapping.
	// Structure is documented below.
	ResourceRecords []ResourceRecordsObservation `json:"resourceRecords,omitempty" tf:"resource_records,omitempty"`

	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	SSLSettings *SSLSettingsObservation `json:"sslSettings,omitempty" tf:"ssl_settings,omitempty"`
}

type DomainMappingParameters struct {

	// Whether the domain creation should override any existing mappings for this domain.
	// By default, overrides are rejected.
	// Default value is STRICT.
	// Possible values are: STRICT, OVERRIDE.
	// +kubebuilder:validation:Optional
	OverrideStrategy *string `json:"overrideStrategy,omitempty" tf:"override_strategy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLSettings *SSLSettingsParameters `json:"sslSettings,omitempty" tf:"ssl_settings,omitempty"`
}

type ResourceRecordsInitParameters struct {
}

type ResourceRecordsObservation struct {

	// Relative name of the object affected by this record. Only applicable for CNAME records. Example: 'www'.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Data for this record. Values vary by record type, as defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1).
	Rrdata *string `json:"rrdata,omitempty" tf:"rrdata,omitempty"`

	// Resource record type. Example: AAAA.
	// Possible values are: A, AAAA, CNAME.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ResourceRecordsParameters struct {
}

type SSLSettingsInitParameters struct {

	// ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will
	// remove SSL support.
	// By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
	// or to configure SSL manually, specify SslManagementType.MANUAL on a CREATE or UPDATE request. You must be
	// authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource.
	// Example: 12345.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned.
	// If MANUAL, certificateId must be manually specified in order to configure SSL for this domain.
	// Possible values are: AUTOMATIC, MANUAL.
	SSLManagementType *string `json:"sslManagementType,omitempty" tf:"ssl_management_type,omitempty"`
}

type SSLSettingsObservation struct {

	// ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will
	// remove SSL support.
	// By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
	// or to configure SSL manually, specify SslManagementType.MANUAL on a CREATE or UPDATE request. You must be
	// authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource.
	// Example: 12345.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// (Output)
	// ID of the managed AuthorizedCertificate resource currently being provisioned, if applicable. Until the new
	// managed certificate has been successfully provisioned, the previous SSL state will be preserved. Once the
	// provisioning process completes, the certificateId field will reflect the new managed certificate and this
	// field will be left empty. To remove SSL support while there is still a pending managed certificate, clear the
	// certificateId field with an update request.
	PendingManagedCertificateID *string `json:"pendingManagedCertificateId,omitempty" tf:"pending_managed_certificate_id,omitempty"`

	// SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned.
	// If MANUAL, certificateId must be manually specified in order to configure SSL for this domain.
	// Possible values are: AUTOMATIC, MANUAL.
	SSLManagementType *string `json:"sslManagementType,omitempty" tf:"ssl_management_type,omitempty"`
}

type SSLSettingsParameters struct {

	// ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will
	// remove SSL support.
	// By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
	// or to configure SSL manually, specify SslManagementType.MANUAL on a CREATE or UPDATE request. You must be
	// authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource.
	// Example: 12345.
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned.
	// If MANUAL, certificateId must be manually specified in order to configure SSL for this domain.
	// Possible values are: AUTOMATIC, MANUAL.
	// +kubebuilder:validation:Optional
	SSLManagementType *string `json:"sslManagementType" tf:"ssl_management_type,omitempty"`
}

// DomainMappingSpec defines the desired state of DomainMapping
type DomainMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMappingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DomainMappingInitParameters `json:"initProvider,omitempty"`
}

// DomainMappingStatus defines the observed state of DomainMapping.
type DomainMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DomainMapping is the Schema for the DomainMappings API. A domain serving an App Engine application.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainMappingSpec   `json:"spec"`
	Status            DomainMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMappingList contains a list of DomainMappings
type DomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMapping `json:"items"`
}

// Repository type metadata.
var (
	DomainMapping_Kind             = "DomainMapping"
	DomainMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMapping_Kind}.String()
	DomainMapping_KindAPIVersion   = DomainMapping_Kind + "." + CRDGroupVersion.String()
	DomainMapping_GroupVersionKind = CRDGroupVersion.WithKind(DomainMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMapping{}, &DomainMappingList{})
}
