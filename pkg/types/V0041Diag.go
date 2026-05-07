// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041Diag = "V0041Diag"
)

type V0041Diag struct {
	api.V0041OpenapiDiagResp
}

// GetKey implements Object.
func (o *V0041Diag) GetKey() object.ObjectKey {
	return ""
}

// GetType implements Object.
func (o *V0041Diag) GetType() object.ObjectType {
	return ObjectTypeV0041Diag
}

// DeepCopyObject implements Object.
func (o *V0041Diag) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041Diag) DeepCopy() *V0041Diag {
	out := new(V0041Diag)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041DiagList struct {
	Items []V0041Diag
}

// GetType implements ObjectList.
func (o *V0041DiagList) GetType() object.ObjectType {
	return ObjectTypeV0041Diag
}

// GetItems implements ObjectList.
func (o *V0041DiagList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041DiagList) AppendItem(object object.Object) {
	out, ok := object.(*V0041Diag)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041DiagList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041DiagList)
	out.Items = make([]V0041Diag, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
