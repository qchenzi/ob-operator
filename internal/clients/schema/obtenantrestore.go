/*
Copyright (c) 2023 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package schema

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	OBTenantRestoreKind     = "OBTenantRestore"
	OBTenantRestoreResource = "obtenantrestores"
)

var (
	OBTenantRestoreGVR = schema.GroupVersionResource{
		Group:    Group,
		Version:  Version,
		Resource: OBTenantRestoreResource,
	}
	OBTenantRestoreGVK = schema.GroupVersionKind{
		Group:   Group,
		Version: Version,
		Kind:    OBTenantRestoreKind,
	}
)
