// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0041ReservationInfo_GetKey(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0041ReservationInfo
		want   object.ObjectKey
	}{
		{
			name:   "empty",
			fields: api.V0041ReservationInfo{},
			want:   "",
		},
		{
			name:   "key",
			fields: api.V0041ReservationInfo{Name: ptr.To("test_0")},
			want:   "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfo{V0041ReservationInfo: tt.fields}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfo_GetType(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0041ReservationInfo
		want   object.ObjectType
	}{
		{
			name:   "type",
			fields: api.V0041ReservationInfo{},
			want:   ObjectTypeV0041ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfo{V0041ReservationInfo: tt.fields}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfo_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0041ReservationInfo
		want   object.Object
	}{
		{
			name:   "empty",
			fields: api.V0041ReservationInfo{},
			want:   &V0041ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0041ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0041ReservationInfo{api.V0041ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfo{V0041ReservationInfo: tt.fields}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfo_DeepCopy(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0041ReservationInfo
		want   *V0041ReservationInfo
	}{
		{
			name:   "empty",
			fields: api.V0041ReservationInfo{},
			want:   &V0041ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0041ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0041ReservationInfo{api.V0041ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfo{V0041ReservationInfo: tt.fields}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfoList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041ReservationInfo
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0041ReservationInfo{},
			want:  ObjectTypeV0041ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfoList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfoList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041ReservationInfo
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0041ReservationInfo{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0041ReservationInfo{{api.V0041ReservationInfo{Name: ptr.To("test_0")}}},
			want:  []object.Object{&V0041ReservationInfo{api.V0041ReservationInfo{Name: ptr.To("test_0")}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfoList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041ReservationInfoList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0041ReservationInfo
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0041ReservationInfo{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0041ReservationInfo{},
			arg:        &V0041ReservationInfo{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0041ReservationInfo{{api.V0041ReservationInfo{Name: ptr.To("test_0")}}},
			arg:        &V0041ReservationInfo{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfoList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0041ReservationInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041ReservationInfoList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041ReservationInfo
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0041ReservationInfo{},
			want:  &V0041ReservationInfoList{Items: []V0041ReservationInfo{}},
		},
		{
			name:  "existing",
			items: []V0041ReservationInfo{{api.V0041ReservationInfo{Name: ptr.To("test_0")}}},
			want: &V0041ReservationInfoList{Items: []V0041ReservationInfo{
				{api.V0041ReservationInfo{Name: ptr.To("test_0")}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041ReservationInfoList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041ReservationInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
