// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043Diag = "V0043Diag"
)

type V0043Diag struct {
	api.V0043OpenapiDiagResp
}

// GetKey implements Object.
func (o *V0043Diag) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0043Diag) GetType() object.ObjectType {
	return ObjectTypeV0043Diag
}

// DeepCopyObject implements Object.
func (o *V0043Diag) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043Diag) DeepCopy() *V0043Diag {
	out := new(V0043Diag)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043DiagList struct {
	Items []V0043Diag
}

// GetType implements ObjectList.
func (o *V0043DiagList) GetType() object.ObjectType {
	return ObjectTypeV0043Diag
}

// GetItems implements ObjectList.
func (o *V0043DiagList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043DiagList) AppendItem(object object.Object) {
	out, ok := object.(*V0043Diag)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043DiagList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043DiagList)
	out.Items = make([]V0043Diag, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
