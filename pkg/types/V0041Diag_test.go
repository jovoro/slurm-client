// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0041Diag_GetKey(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0041OpenapiDiagResp
		want  object.ObjectKey
	}{
		{
			name:  "key",
			field: api.V0041OpenapiDiagResp{},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Diag{V0041OpenapiDiagResp: tt.field}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Diag.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Diag_GetType(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0041OpenapiDiagResp
		want  object.ObjectType
	}{
		{
			name:  "type",
			field: api.V0041OpenapiDiagResp{},
			want:  ObjectTypeV0041Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Diag{V0041OpenapiDiagResp: tt.field}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Diag.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Diag_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0041OpenapiDiagResp
		want  object.Object
	}{
		{
			name:  "empty",
			field: api.V0041OpenapiDiagResp{},
			want:  &V0041Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Diag{V0041OpenapiDiagResp: tt.field}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Diag.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041Diag_DeepCopy(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0041OpenapiDiagResp
		want  *V0041Diag
	}{
		{
			name:  "empty",
			field: api.V0041OpenapiDiagResp{},
			want:  &V0041Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041Diag{V0041OpenapiDiagResp: tt.field}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041Diag.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041DiagList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041Diag
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0041Diag{},
			want:  ObjectTypeV0041Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041DiagList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041DiagList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041DiagList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041Diag
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0041Diag{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0041Diag{{V0041OpenapiDiagResp: api.V0041OpenapiDiagResp{}}},
			want:  []object.Object{&V0041Diag{api.V0041OpenapiDiagResp{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041DiagList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041DiagList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0041DiagList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0041Diag
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0041Diag{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0041Diag{},
			arg:        &V0041Diag{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0041Diag{{V0041OpenapiDiagResp: api.V0041OpenapiDiagResp{}}},
			arg:        &V0041Diag{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041DiagList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0041DiagList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0041DiagList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0041Diag
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0041Diag{},
			want:  &V0041DiagList{Items: []V0041Diag{}},
		},
		{
			name:  "existing",
			items: []V0041Diag{{V0041OpenapiDiagResp: api.V0041OpenapiDiagResp{}}},
			want:  &V0041DiagList{Items: []V0041Diag{{api.V0041OpenapiDiagResp{}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0041DiagList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0041DiagList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
