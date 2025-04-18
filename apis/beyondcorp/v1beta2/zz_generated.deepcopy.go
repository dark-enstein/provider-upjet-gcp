//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnection) DeepCopyInto(out *AppConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnection.
func (in *AppConnection) DeepCopy() *AppConnection {
	if in == nil {
		return nil
	}
	out := new(AppConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionInitParameters) DeepCopyInto(out *AppConnectionInitParameters) {
	*out = *in
	if in.ApplicationEndpoint != nil {
		in, out := &in.ApplicationEndpoint, &out.ApplicationEndpoint
		*out = new(ApplicationEndpointInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Connectors != nil {
		in, out := &in.Connectors, &out.Connectors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConnectorsRefs != nil {
		in, out := &in.ConnectorsRefs, &out.ConnectorsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectorsSelector != nil {
		in, out := &in.ConnectorsSelector, &out.ConnectorsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionInitParameters.
func (in *AppConnectionInitParameters) DeepCopy() *AppConnectionInitParameters {
	if in == nil {
		return nil
	}
	out := new(AppConnectionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionList) DeepCopyInto(out *AppConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionList.
func (in *AppConnectionList) DeepCopy() *AppConnectionList {
	if in == nil {
		return nil
	}
	out := new(AppConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionObservation) DeepCopyInto(out *AppConnectionObservation) {
	*out = *in
	if in.ApplicationEndpoint != nil {
		in, out := &in.ApplicationEndpoint, &out.ApplicationEndpoint
		*out = new(ApplicationEndpointObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Connectors != nil {
		in, out := &in.Connectors, &out.Connectors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionObservation.
func (in *AppConnectionObservation) DeepCopy() *AppConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(AppConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionParameters) DeepCopyInto(out *AppConnectionParameters) {
	*out = *in
	if in.ApplicationEndpoint != nil {
		in, out := &in.ApplicationEndpoint, &out.ApplicationEndpoint
		*out = new(ApplicationEndpointParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Connectors != nil {
		in, out := &in.Connectors, &out.Connectors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ConnectorsRefs != nil {
		in, out := &in.ConnectorsRefs, &out.ConnectorsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectorsSelector != nil {
		in, out := &in.ConnectorsSelector, &out.ConnectorsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionParameters.
func (in *AppConnectionParameters) DeepCopy() *AppConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(AppConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionSpec) DeepCopyInto(out *AppConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionSpec.
func (in *AppConnectionSpec) DeepCopy() *AppConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(AppConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectionStatus) DeepCopyInto(out *AppConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectionStatus.
func (in *AppConnectionStatus) DeepCopy() *AppConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(AppConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnector) DeepCopyInto(out *AppConnector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnector.
func (in *AppConnector) DeepCopy() *AppConnector {
	if in == nil {
		return nil
	}
	out := new(AppConnector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppConnector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorInitParameters) DeepCopyInto(out *AppConnectorInitParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PrincipalInfo != nil {
		in, out := &in.PrincipalInfo, &out.PrincipalInfo
		*out = new(PrincipalInfoInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorInitParameters.
func (in *AppConnectorInitParameters) DeepCopy() *AppConnectorInitParameters {
	if in == nil {
		return nil
	}
	out := new(AppConnectorInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorList) DeepCopyInto(out *AppConnectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppConnector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorList.
func (in *AppConnectorList) DeepCopy() *AppConnectorList {
	if in == nil {
		return nil
	}
	out := new(AppConnectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppConnectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorObservation) DeepCopyInto(out *AppConnectorObservation) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PrincipalInfo != nil {
		in, out := &in.PrincipalInfo, &out.PrincipalInfo
		*out = new(PrincipalInfoObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorObservation.
func (in *AppConnectorObservation) DeepCopy() *AppConnectorObservation {
	if in == nil {
		return nil
	}
	out := new(AppConnectorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorParameters) DeepCopyInto(out *AppConnectorParameters) {
	*out = *in
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PrincipalInfo != nil {
		in, out := &in.PrincipalInfo, &out.PrincipalInfo
		*out = new(PrincipalInfoParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorParameters.
func (in *AppConnectorParameters) DeepCopy() *AppConnectorParameters {
	if in == nil {
		return nil
	}
	out := new(AppConnectorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorSpec) DeepCopyInto(out *AppConnectorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorSpec.
func (in *AppConnectorSpec) DeepCopy() *AppConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(AppConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppConnectorStatus) DeepCopyInto(out *AppConnectorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppConnectorStatus.
func (in *AppConnectorStatus) DeepCopy() *AppConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(AppConnectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationEndpointInitParameters) DeepCopyInto(out *ApplicationEndpointInitParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationEndpointInitParameters.
func (in *ApplicationEndpointInitParameters) DeepCopy() *ApplicationEndpointInitParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationEndpointInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationEndpointObservation) DeepCopyInto(out *ApplicationEndpointObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationEndpointObservation.
func (in *ApplicationEndpointObservation) DeepCopy() *ApplicationEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationEndpointParameters) DeepCopyInto(out *ApplicationEndpointParameters) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationEndpointParameters.
func (in *ApplicationEndpointParameters) DeepCopy() *ApplicationEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayInitParameters) DeepCopyInto(out *GatewayInitParameters) {
	*out = *in
	if in.AppGateway != nil {
		in, out := &in.AppGateway, &out.AppGateway
		*out = new(string)
		**out = **in
	}
	if in.AppGatewayRef != nil {
		in, out := &in.AppGatewayRef, &out.AppGatewayRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AppGatewaySelector != nil {
		in, out := &in.AppGatewaySelector, &out.AppGatewaySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayInitParameters.
func (in *GatewayInitParameters) DeepCopy() *GatewayInitParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayObservation) DeepCopyInto(out *GatewayObservation) {
	*out = *in
	if in.AppGateway != nil {
		in, out := &in.AppGateway, &out.AppGateway
		*out = new(string)
		**out = **in
	}
	if in.IngressPort != nil {
		in, out := &in.IngressPort, &out.IngressPort
		*out = new(float64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayObservation.
func (in *GatewayObservation) DeepCopy() *GatewayObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParameters) DeepCopyInto(out *GatewayParameters) {
	*out = *in
	if in.AppGateway != nil {
		in, out := &in.AppGateway, &out.AppGateway
		*out = new(string)
		**out = **in
	}
	if in.AppGatewayRef != nil {
		in, out := &in.AppGatewayRef, &out.AppGatewayRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.AppGatewaySelector != nil {
		in, out := &in.AppGatewaySelector, &out.AppGatewaySelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParameters.
func (in *GatewayParameters) DeepCopy() *GatewayParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalInfoInitParameters) DeepCopyInto(out *PrincipalInfoInitParameters) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountInitParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalInfoInitParameters.
func (in *PrincipalInfoInitParameters) DeepCopy() *PrincipalInfoInitParameters {
	if in == nil {
		return nil
	}
	out := new(PrincipalInfoInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalInfoObservation) DeepCopyInto(out *PrincipalInfoObservation) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountObservation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalInfoObservation.
func (in *PrincipalInfoObservation) DeepCopy() *PrincipalInfoObservation {
	if in == nil {
		return nil
	}
	out := new(PrincipalInfoObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalInfoParameters) DeepCopyInto(out *PrincipalInfoParameters) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountParameters)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalInfoParameters.
func (in *PrincipalInfoParameters) DeepCopy() *PrincipalInfoParameters {
	if in == nil {
		return nil
	}
	out := new(PrincipalInfoParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountInitParameters) DeepCopyInto(out *ServiceAccountInitParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.EmailRef != nil {
		in, out := &in.EmailRef, &out.EmailRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EmailSelector != nil {
		in, out := &in.EmailSelector, &out.EmailSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountInitParameters.
func (in *ServiceAccountInitParameters) DeepCopy() *ServiceAccountInitParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountObservation) DeepCopyInto(out *ServiceAccountObservation) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountObservation.
func (in *ServiceAccountObservation) DeepCopy() *ServiceAccountObservation {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountParameters) DeepCopyInto(out *ServiceAccountParameters) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.EmailRef != nil {
		in, out := &in.EmailRef, &out.EmailRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.EmailSelector != nil {
		in, out := &in.EmailSelector, &out.EmailSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountParameters.
func (in *ServiceAccountParameters) DeepCopy() *ServiceAccountParameters {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountParameters)
	in.DeepCopyInto(out)
	return out
}
