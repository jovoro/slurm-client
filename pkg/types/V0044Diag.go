// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0044Diag = "V0044Diag"
)

type V0044Diag struct {
	api.V0044OpenapiDiagResp
}

// GetKey implements Object.
func (o *V0044Diag) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0044Diag) GetType() object.ObjectType {
	return ObjectTypeV0044Diag
}

// DeepCopyObject implements Object.
func (o *V0044Diag) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0044Diag) DeepCopy() *V0044Diag {
	out := new(V0044Diag)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0044DiagList struct {
	Items []V0044Diag
}

// GetType implements ObjectList.
func (o *V0044DiagList) GetType() object.ObjectType {
	return ObjectTypeV0044Diag
}

// GetItems implements ObjectList.
func (o *V0044DiagList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0044DiagList) AppendItem(object object.Object) {
	out, ok := object.(*V0044Diag)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0044DiagList) DeepCopyObjectList() object.ObjectList {
	out := new(V0044DiagList)
	out.Items = make([]V0044Diag, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
