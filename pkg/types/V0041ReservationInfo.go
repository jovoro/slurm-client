// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0041ReservationInfo = "V0041ReservationInfo"
)

type V0041ReservationInfo struct {
	api.V0041ReservationInfo
}

// GetKey implements Object.
func (o *V0041ReservationInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0041ReservationInfo) GetType() object.ObjectType {
	return ObjectTypeV0041ReservationInfo
}

// DeepCopyObject implements Object.
func (o *V0041ReservationInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0041ReservationInfo) DeepCopy() *V0041ReservationInfo {
	out := new(V0041ReservationInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0041ReservationInfoList struct {
	Items []V0041ReservationInfo
}

// GetType implements ObjectList.
func (o *V0041ReservationInfoList) GetType() object.ObjectType {
	return ObjectTypeV0041ReservationInfo
}

// GetItems implements ObjectList.
func (o *V0041ReservationInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0041ReservationInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0041ReservationInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0041ReservationInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0041ReservationInfoList)
	out.Items = make([]V0041ReservationInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
