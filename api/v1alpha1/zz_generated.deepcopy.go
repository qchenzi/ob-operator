//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupVolumeSpec) DeepCopyInto(out *BackupVolumeSpec) {
	*out = *in
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(v1.Volume)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupVolumeSpec.
func (in *BackupVolumeSpec) DeepCopy() *BackupVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(BackupVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorTemplate) DeepCopyInto(out *MonitorTemplate) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(ObagentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorTemplate.
func (in *MonitorTemplate) DeepCopy() *MonitorTemplate {
	if in == nil {
		return nil
	}
	out := new(MonitorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBCluster) DeepCopyInto(out *OBCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBCluster.
func (in *OBCluster) DeepCopy() *OBCluster {
	if in == nil {
		return nil
	}
	out := new(OBCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterBackup) DeepCopyInto(out *OBClusterBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterBackup.
func (in *OBClusterBackup) DeepCopy() *OBClusterBackup {
	if in == nil {
		return nil
	}
	out := new(OBClusterBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterBackupList) DeepCopyInto(out *OBClusterBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBClusterBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterBackupList.
func (in *OBClusterBackupList) DeepCopy() *OBClusterBackupList {
	if in == nil {
		return nil
	}
	out := new(OBClusterBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterBackupSpec) DeepCopyInto(out *OBClusterBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterBackupSpec.
func (in *OBClusterBackupSpec) DeepCopy() *OBClusterBackupSpec {
	if in == nil {
		return nil
	}
	out := new(OBClusterBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterBackupStatus) DeepCopyInto(out *OBClusterBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterBackupStatus.
func (in *OBClusterBackupStatus) DeepCopy() *OBClusterBackupStatus {
	if in == nil {
		return nil
	}
	out := new(OBClusterBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterList) DeepCopyInto(out *OBClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterList.
func (in *OBClusterList) DeepCopy() *OBClusterList {
	if in == nil {
		return nil
	}
	out := new(OBClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterRestore) DeepCopyInto(out *OBClusterRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterRestore.
func (in *OBClusterRestore) DeepCopy() *OBClusterRestore {
	if in == nil {
		return nil
	}
	out := new(OBClusterRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterRestoreList) DeepCopyInto(out *OBClusterRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBClusterRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterRestoreList.
func (in *OBClusterRestoreList) DeepCopy() *OBClusterRestoreList {
	if in == nil {
		return nil
	}
	out := new(OBClusterRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBClusterRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterRestoreSpec) DeepCopyInto(out *OBClusterRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterRestoreSpec.
func (in *OBClusterRestoreSpec) DeepCopy() *OBClusterRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(OBClusterRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterRestoreStatus) DeepCopyInto(out *OBClusterRestoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterRestoreStatus.
func (in *OBClusterRestoreStatus) DeepCopy() *OBClusterRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(OBClusterRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterSpec) DeepCopyInto(out *OBClusterSpec) {
	*out = *in
	if in.OBServerTemplate != nil {
		in, out := &in.OBServerTemplate, &out.OBServerTemplate
		*out = new(OBServerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitorTemplate != nil {
		in, out := &in.MonitorTemplate, &out.MonitorTemplate
		*out = new(MonitorTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupVolume != nil {
		in, out := &in.BackupVolume, &out.BackupVolume
		*out = new(BackupVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]Parameter, len(*in))
		copy(*out, *in)
	}
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = make([]OBZoneTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserSecrets != nil {
		in, out := &in.UserSecrets, &out.UserSecrets
		*out = new(OBUserSecrets)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterSpec.
func (in *OBClusterSpec) DeepCopy() *OBClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OBClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBClusterStatus) DeepCopyInto(out *OBClusterStatus) {
	*out = *in
	if in.OperationContext != nil {
		in, out := &in.OperationContext, &out.OperationContext
		*out = new(OperationContext)
		(*in).DeepCopyInto(*out)
	}
	if in.OBZoneStatus != nil {
		in, out := &in.OBZoneStatus, &out.OBZoneStatus
		*out = make([]OBZoneReplicaStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBClusterStatus.
func (in *OBClusterStatus) DeepCopy() *OBClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OBClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBParameter) DeepCopyInto(out *OBParameter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBParameter.
func (in *OBParameter) DeepCopy() *OBParameter {
	if in == nil {
		return nil
	}
	out := new(OBParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBParameter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBParameterList) DeepCopyInto(out *OBParameterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBParameter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBParameterList.
func (in *OBParameterList) DeepCopy() *OBParameterList {
	if in == nil {
		return nil
	}
	out := new(OBParameterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBParameterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBParameterSpec) DeepCopyInto(out *OBParameterSpec) {
	*out = *in
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = new(Parameter)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBParameterSpec.
func (in *OBParameterSpec) DeepCopy() *OBParameterSpec {
	if in == nil {
		return nil
	}
	out := new(OBParameterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBParameterStatus) DeepCopyInto(out *OBParameterStatus) {
	*out = *in
	if in.OperationContext != nil {
		in, out := &in.OperationContext, &out.OperationContext
		*out = new(OperationContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = new([]ParameterValue)
		if **in != nil {
			in, out := *in, *out
			*out = make([]ParameterValue, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBParameterStatus.
func (in *OBParameterStatus) DeepCopy() *OBParameterStatus {
	if in == nil {
		return nil
	}
	out := new(OBParameterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServer) DeepCopyInto(out *OBServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServer.
func (in *OBServer) DeepCopy() *OBServer {
	if in == nil {
		return nil
	}
	out := new(OBServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServerList) DeepCopyInto(out *OBServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServerList.
func (in *OBServerList) DeepCopy() *OBServerList {
	if in == nil {
		return nil
	}
	out := new(OBServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServerReplicaStatus) DeepCopyInto(out *OBServerReplicaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServerReplicaStatus.
func (in *OBServerReplicaStatus) DeepCopy() *OBServerReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(OBServerReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServerSpec) DeepCopyInto(out *OBServerSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OBServerTemplate != nil {
		in, out := &in.OBServerTemplate, &out.OBServerTemplate
		*out = new(OBServerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitorTemplate != nil {
		in, out := &in.MonitorTemplate, &out.MonitorTemplate
		*out = new(MonitorTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupVolume != nil {
		in, out := &in.BackupVolume, &out.BackupVolume
		*out = new(BackupVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServerSpec.
func (in *OBServerSpec) DeepCopy() *OBServerSpec {
	if in == nil {
		return nil
	}
	out := new(OBServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServerStatus) DeepCopyInto(out *OBServerStatus) {
	*out = *in
	if in.OperationContext != nil {
		in, out := &in.OperationContext, &out.OperationContext
		*out = new(OperationContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServerStatus.
func (in *OBServerStatus) DeepCopy() *OBServerStatus {
	if in == nil {
		return nil
	}
	out := new(OBServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBServerTemplate) DeepCopyInto(out *OBServerTemplate) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(ResourceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(OceanbaseStorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBServerTemplate.
func (in *OBServerTemplate) DeepCopy() *OBServerTemplate {
	if in == nil {
		return nil
	}
	out := new(OBServerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenant) DeepCopyInto(out *OBTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenant.
func (in *OBTenant) DeepCopy() *OBTenant {
	if in == nil {
		return nil
	}
	out := new(OBTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantBackup) DeepCopyInto(out *OBTenantBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantBackup.
func (in *OBTenantBackup) DeepCopy() *OBTenantBackup {
	if in == nil {
		return nil
	}
	out := new(OBTenantBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenantBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantBackupList) DeepCopyInto(out *OBTenantBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBTenantBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantBackupList.
func (in *OBTenantBackupList) DeepCopy() *OBTenantBackupList {
	if in == nil {
		return nil
	}
	out := new(OBTenantBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenantBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantBackupSpec) DeepCopyInto(out *OBTenantBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantBackupSpec.
func (in *OBTenantBackupSpec) DeepCopy() *OBTenantBackupSpec {
	if in == nil {
		return nil
	}
	out := new(OBTenantBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantBackupStatus) DeepCopyInto(out *OBTenantBackupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantBackupStatus.
func (in *OBTenantBackupStatus) DeepCopy() *OBTenantBackupStatus {
	if in == nil {
		return nil
	}
	out := new(OBTenantBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantList) DeepCopyInto(out *OBTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantList.
func (in *OBTenantList) DeepCopy() *OBTenantList {
	if in == nil {
		return nil
	}
	out := new(OBTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantRestore) DeepCopyInto(out *OBTenantRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantRestore.
func (in *OBTenantRestore) DeepCopy() *OBTenantRestore {
	if in == nil {
		return nil
	}
	out := new(OBTenantRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenantRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantRestoreList) DeepCopyInto(out *OBTenantRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBTenantRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantRestoreList.
func (in *OBTenantRestoreList) DeepCopy() *OBTenantRestoreList {
	if in == nil {
		return nil
	}
	out := new(OBTenantRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBTenantRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantRestoreSpec) DeepCopyInto(out *OBTenantRestoreSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantRestoreSpec.
func (in *OBTenantRestoreSpec) DeepCopy() *OBTenantRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(OBTenantRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantRestoreStatus) DeepCopyInto(out *OBTenantRestoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantRestoreStatus.
func (in *OBTenantRestoreStatus) DeepCopy() *OBTenantRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(OBTenantRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantSpec) DeepCopyInto(out *OBTenantSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantSpec.
func (in *OBTenantSpec) DeepCopy() *OBTenantSpec {
	if in == nil {
		return nil
	}
	out := new(OBTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBTenantStatus) DeepCopyInto(out *OBTenantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBTenantStatus.
func (in *OBTenantStatus) DeepCopy() *OBTenantStatus {
	if in == nil {
		return nil
	}
	out := new(OBTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBUnit) DeepCopyInto(out *OBUnit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBUnit.
func (in *OBUnit) DeepCopy() *OBUnit {
	if in == nil {
		return nil
	}
	out := new(OBUnit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBUnit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBUnitList) DeepCopyInto(out *OBUnitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBUnit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBUnitList.
func (in *OBUnitList) DeepCopy() *OBUnitList {
	if in == nil {
		return nil
	}
	out := new(OBUnitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBUnitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBUnitSpec) DeepCopyInto(out *OBUnitSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBUnitSpec.
func (in *OBUnitSpec) DeepCopy() *OBUnitSpec {
	if in == nil {
		return nil
	}
	out := new(OBUnitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBUnitStatus) DeepCopyInto(out *OBUnitStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBUnitStatus.
func (in *OBUnitStatus) DeepCopy() *OBUnitStatus {
	if in == nil {
		return nil
	}
	out := new(OBUnitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBUserSecrets) DeepCopyInto(out *OBUserSecrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBUserSecrets.
func (in *OBUserSecrets) DeepCopy() *OBUserSecrets {
	if in == nil {
		return nil
	}
	out := new(OBUserSecrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZone) DeepCopyInto(out *OBZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZone.
func (in *OBZone) DeepCopy() *OBZone {
	if in == nil {
		return nil
	}
	out := new(OBZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneList) DeepCopyInto(out *OBZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OBZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneList.
func (in *OBZoneList) DeepCopy() *OBZoneList {
	if in == nil {
		return nil
	}
	out := new(OBZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OBZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneReplicaStatus) DeepCopyInto(out *OBZoneReplicaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneReplicaStatus.
func (in *OBZoneReplicaStatus) DeepCopy() *OBZoneReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(OBZoneReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneSpec) DeepCopyInto(out *OBZoneSpec) {
	*out = *in
	in.Topology.DeepCopyInto(&out.Topology)
	if in.OBServerTemplate != nil {
		in, out := &in.OBServerTemplate, &out.OBServerTemplate
		*out = new(OBServerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.MonitorTemplate != nil {
		in, out := &in.MonitorTemplate, &out.MonitorTemplate
		*out = new(MonitorTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupVolume != nil {
		in, out := &in.BackupVolume, &out.BackupVolume
		*out = new(BackupVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneSpec.
func (in *OBZoneSpec) DeepCopy() *OBZoneSpec {
	if in == nil {
		return nil
	}
	out := new(OBZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneStatus) DeepCopyInto(out *OBZoneStatus) {
	*out = *in
	if in.OperationContext != nil {
		in, out := &in.OperationContext, &out.OperationContext
		*out = new(OperationContext)
		(*in).DeepCopyInto(*out)
	}
	if in.OBServerStatus != nil {
		in, out := &in.OBServerStatus, &out.OBServerStatus
		*out = make([]OBServerReplicaStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneStatus.
func (in *OBZoneStatus) DeepCopy() *OBZoneStatus {
	if in == nil {
		return nil
	}
	out := new(OBZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OBZoneTopology) DeepCopyInto(out *OBZoneTopology) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OBZoneTopology.
func (in *OBZoneTopology) DeepCopy() *OBZoneTopology {
	if in == nil {
		return nil
	}
	out := new(OBZoneTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObagentStorageSpec) DeepCopyInto(out *ObagentStorageSpec) {
	*out = *in
	if in.ConfigStorage != nil {
		in, out := &in.ConfigStorage, &out.ConfigStorage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObagentStorageSpec.
func (in *ObagentStorageSpec) DeepCopy() *ObagentStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ObagentStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OceanbaseStorageSpec) DeepCopyInto(out *OceanbaseStorageSpec) {
	*out = *in
	if in.DataStorage != nil {
		in, out := &in.DataStorage, &out.DataStorage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RedoLogStorage != nil {
		in, out := &in.RedoLogStorage, &out.RedoLogStorage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LogStorage != nil {
		in, out := &in.LogStorage, &out.LogStorage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceanbaseStorageSpec.
func (in *OceanbaseStorageSpec) DeepCopy() *OceanbaseStorageSpec {
	if in == nil {
		return nil
	}
	out := new(OceanbaseStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationContext) DeepCopyInto(out *OperationContext) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationContext.
func (in *OperationContext) DeepCopy() *OperationContext {
	if in == nil {
		return nil
	}
	out := new(OperationContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVCStatus) DeepCopyInto(out *PVCStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVCStatus.
func (in *PVCStatus) DeepCopy() *PVCStatus {
	if in == nil {
		return nil
	}
	out := new(PVCStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Parameter) DeepCopyInto(out *Parameter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Parameter.
func (in *Parameter) DeepCopy() *Parameter {
	if in == nil {
		return nil
	}
	out := new(Parameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterValue) DeepCopyInto(out *ParameterValue) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterValue.
func (in *ParameterValue) DeepCopy() *ParameterValue {
	if in == nil {
		return nil
	}
	out := new(ParameterValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	out.Cpu = in.Cpu.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	out.Size = in.Size.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}