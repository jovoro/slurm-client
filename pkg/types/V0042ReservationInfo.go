// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

const (
	ObjectTypeV0042ReservationInfo = "V0042ReservationInfo"
)

type V0042ReservationInfo struct {
	api.V0042ReservationInfo
}

// GetKey implements Object.
func (o *V0042ReservationInfo) GetKey() object.ObjectKey {
	return object.ObjectKey(ptr.Deref(o.Name, ""))
}

// GetType implements Object.
func (o *V0042ReservationInfo) GetType() object.ObjectType {
	return ObjectTypeV0042ReservationInfo
}

// DeepCopyObject implements Object.
func (o *V0042ReservationInfo) DeepCopyObject() object.Object {
	return o.DeepCopy()
}

func (o *V0042ReservationInfo) DeepCopy() *V0042ReservationInfo {
	out := new(V0042ReservationInfo)
	utils.RemarshalOrDie(o, out)
	return out
}

type V0042ReservationInfoList struct {
	Items []V0042ReservationInfo
}

// GetType implements ObjectList.
func (o *V0042ReservationInfoList) GetType() object.ObjectType {
	return ObjectTypeV0042ReservationInfo
}

// GetItems implements ObjectList.
func (o *V0042ReservationInfoList) GetItems() []object.Object {
	list := make([]object.Object, len(o.Items))
	for i, item := range o.Items {
		list[i] = item.DeepCopyObject()
	}
	return list
}

// AppendItem implements ObjectList.
func (o *V0042ReservationInfoList) AppendItem(object object.Object) {
	out, ok := object.(*V0042ReservationInfo)
	if ok {
		utils.RemarshalOrDie(object, out)
		o.Items = append(o.Items, *out)
	}
}

// DeepCopyObjectList implements ObjectList.
func (o *V0042ReservationInfoList) DeepCopyObjectList() object.ObjectList {
	out := new(V0042ReservationInfoList)
	out.Items = make([]V0042ReservationInfo, len(o.Items))
	for i, item := range o.Items {
		out.Items[i] = *item.DeepCopy()
	}
	return out
}
