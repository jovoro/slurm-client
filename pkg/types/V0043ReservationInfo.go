// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0043ReservationInfo = "V0043ReservationInfo"
)

type V0043ReservationInfo struct {
	api.V0043ReservationInfo
}

// GetKey implements Object.
func (o *V0043ReservationInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0043ReservationInfo) GetType() object.ObjectType {
	return ObjectTypeV0043ReservationInfo
}

// DeepCopyObject implements Object.
func (o *V0043ReservationInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0043ReservationInfo) DeepCopy() *V0043ReservationInfo {
	out := new(V0043ReservationInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0043ReservationInfoList struct {
	Items []V0043ReservationInfo
}

// GetType implements ObjectList.
func (o *V0043ReservationInfoList) GetType() object.ObjectType {
	return ObjectTypeV0043ReservationInfo
}

// GetItems implements ObjectList.
func (o *V0043ReservationInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0043ReservationInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0043ReservationInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0043ReservationInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0043ReservationInfoList)
	out.Items = make([]V0043ReservationInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
