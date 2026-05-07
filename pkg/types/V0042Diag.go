// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042Diag = "V0042Diag"
)

type V0042Diag struct {
	api.V0042OpenapiDiagResp
}

// GetKey implements Object.
func (o *V0042Diag) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0042Diag) GetType() object.ObjectType {
	return ObjectTypeV0042Diag
}

// DeepCopyObject implements Object.
func (o *V0042Diag) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042Diag) DeepCopy() *V0042Diag {
	out := new(V0042Diag)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042DiagList struct {
	Items []V0042Diag
}

// GetType implements ObjectList.
func (o *V0042DiagList) GetType() object.ObjectType {
	return ObjectTypeV0042Diag
}

// GetItems implements ObjectList.
func (o *V0042DiagList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042DiagList) AppendItem(object object.Object) {
	out, ok := object.(*V0042Diag)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042DiagList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042DiagList)
	out.Items = make([]V0042Diag, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
