// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042ReservationInfo_GetKey(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0042ReservationInfo
		want   object.ObjectKey
	}{
		{
			name:   "empty",
			fields: api.V0042ReservationInfo{},
			want:   "",
		},
		{
			name:   "key",
			fields: api.V0042ReservationInfo{Name: ptr.To("test_0")},
			want:   "test_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfo{V0042ReservationInfo: tt.fields}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfo.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfo_GetType(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0042ReservationInfo
		want   object.ObjectType
	}{
		{
			name:   "type",
			fields: api.V0042ReservationInfo{},
			want:   ObjectTypeV0042ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfo{V0042ReservationInfo: tt.fields}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfo.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfo_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0042ReservationInfo
		want   object.Object
	}{
		{
			name:   "empty",
			fields: api.V0042ReservationInfo{},
			want:   &V0042ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0042ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0042ReservationInfo{api.V0042ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfo{V0042ReservationInfo: tt.fields}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfo.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfo_DeepCopy(t *testing.T) {
	tests := []struct {
		name   string
		fields api.V0042ReservationInfo
		want   *V0042ReservationInfo
	}{
		{
			name:   "empty",
			fields: api.V0042ReservationInfo{},
			want:   &V0042ReservationInfo{},
		},
		{
			name:   "key",
			fields: api.V0042ReservationInfo{Name: ptr.To("test_0")},
			want:   &V0042ReservationInfo{api.V0042ReservationInfo{Name: ptr.To("test_0")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfo{V0042ReservationInfo: tt.fields}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfo.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfoList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042ReservationInfo
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0042ReservationInfo{},
			want:  ObjectTypeV0042ReservationInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfoList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfoList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfoList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042ReservationInfo
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0042ReservationInfo{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0042ReservationInfo{{api.V0042ReservationInfo{Name: ptr.To("test_0")}}},
			want:  []object.Object{&V0042ReservationInfo{api.V0042ReservationInfo{Name: ptr.To("test_0")}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfoList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfoList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042ReservationInfoList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0042ReservationInfo
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0042ReservationInfo{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0042ReservationInfo{},
			arg:        &V0042ReservationInfo{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0042ReservationInfo{{api.V0042ReservationInfo{Name: ptr.To("test_0")}}},
			arg:        &V0042ReservationInfo{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfoList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0042ReservationInfoList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042ReservationInfoList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042ReservationInfo
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0042ReservationInfo{},
			want:  &V0042ReservationInfoList{Items: []V0042ReservationInfo{}},
		},
		{
			name:  "existing",
			items: []V0042ReservationInfo{{api.V0042ReservationInfo{Name: ptr.To("test_0")}}},
			want: &V0042ReservationInfoList{Items: []V0042ReservationInfo{
				{api.V0042ReservationInfo{Name: ptr.To("test_0")}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042ReservationInfoList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042ReservationInfoList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
