// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0043"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0043Diag_GetKey(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0043OpenapiDiagResp
		want  object.ObjectKey
	}{
		{
			name:  "key",
			field: api.V0043OpenapiDiagResp{},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Diag{V0043OpenapiDiagResp: tt.field}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Diag.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Diag_GetType(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0043OpenapiDiagResp
		want  object.ObjectType
	}{
		{
			name:  "type",
			field: api.V0043OpenapiDiagResp{},
			want:  ObjectTypeV0043Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Diag{V0043OpenapiDiagResp: tt.field}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Diag.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Diag_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0043OpenapiDiagResp
		want  object.Object
	}{
		{
			name:  "empty",
			field: api.V0043OpenapiDiagResp{},
			want:  &V0043Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Diag{V0043OpenapiDiagResp: tt.field}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Diag.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043Diag_DeepCopy(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0043OpenapiDiagResp
		want  *V0043Diag
	}{
		{
			name:  "empty",
			field: api.V0043OpenapiDiagResp{},
			want:  &V0043Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043Diag{V0043OpenapiDiagResp: tt.field}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043Diag.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043DiagList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043Diag
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0043Diag{},
			want:  ObjectTypeV0043Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043DiagList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043DiagList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043DiagList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043Diag
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0043Diag{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0043Diag{{V0043OpenapiDiagResp: api.V0043OpenapiDiagResp{}}},
			want:  []object.Object{&V0043Diag{api.V0043OpenapiDiagResp{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043DiagList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043DiagList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0043DiagList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0043Diag
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0043Diag{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0043Diag{},
			arg:        &V0043Diag{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0043Diag{{V0043OpenapiDiagResp: api.V0043OpenapiDiagResp{}}},
			arg:        &V0043Diag{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043DiagList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0043DiagList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0043DiagList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0043Diag
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0043Diag{},
			want:  &V0043DiagList{Items: []V0043Diag{}},
		},
		{
			name:  "existing",
			items: []V0043Diag{{V0043OpenapiDiagResp: api.V0043OpenapiDiagResp{}}},
			want:  &V0043DiagList{Items: []V0043Diag{{api.V0043OpenapiDiagResp{}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0043DiagList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0043DiagList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
