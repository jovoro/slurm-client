// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043ReservationInfo_GetKey(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0043ReservationInfo
		want   object.ObjectKey
	}{
		{
			name:   "empty",
			fields: api.V0043ReservationInfo{},
			want:   "",
		},
		{
			name:   "key",
			fields: api.V0043ReservationInfo{Name: ptr.To("test_0")},
			want:   "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfo{V0043ReservationInfo: tt.fields}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfo_GetType(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0043ReservationInfo
		want   object.ObjectType
	}{
		{
			name:   "type",
			fields: api.V0043ReservationInfo{},
			want:   ObjectTypeV0043ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfo{V0043ReservationInfo: tt.fields}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfo_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0043ReservationInfo
		want   object.Object
	}{
		{
			name:   "empty",
			fields: api.V0043ReservationInfo{},
			want:   &V0043ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0043ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0043ReservationInfo{api.V0043ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfo{V0043ReservationInfo: tt.fields}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfo_DeepCopy(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0043ReservationInfo
		want   *V0043ReservationInfo
	}{
		{
			name:   "empty",
			fields: api.V0043ReservationInfo{},
			want:   &V0043ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0043ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0043ReservationInfo{api.V0043ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfo{V0043ReservationInfo: tt.fields}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfoList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043ReservationInfo
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0043ReservationInfo{},
			want:  ObjectTypeV0043ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfoList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfoList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043ReservationInfo
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0043ReservationInfo{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0043ReservationInfo{{api.V0043ReservationInfo{Name: ptr.To("test_0")}}},
			want:  []object.Object{&V0043ReservationInfo{api.V0043ReservationInfo{Name: ptr.To("test_0")}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfoList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043ReservationInfoList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0043ReservationInfo
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0043ReservationInfo{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0043ReservationInfo{},
			arg:        &V0043ReservationInfo{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0043ReservationInfo{{api.V0043ReservationInfo{Name: ptr.To("test_0")}}},
			arg:        &V0043ReservationInfo{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfoList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0043ReservationInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0043ReservationInfoList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043ReservationInfo
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0043ReservationInfo{},
			want:  &V0043ReservationInfoList{Items: []V0043ReservationInfo{}},
		},
		{
			name:  "existing",
			items: []V0043ReservationInfo{{api.V0043ReservationInfo{Name: ptr.To("test_0")}}},
			want: &V0043ReservationInfoList{Items: []V0043ReservationInfo{
				{api.V0043ReservationInfo{Name: ptr.To("test_0")}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043ReservationInfoList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043ReservationInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
