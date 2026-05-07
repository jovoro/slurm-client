// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0042Diag_GetKey(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0042OpenapiDiagResp
		want  object.ObjectKey
	}{
		{
			name:  "key",
			field: api.V0042OpenapiDiagResp{},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Diag{V0042OpenapiDiagResp: tt.field}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Diag.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Diag_GetType(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0042OpenapiDiagResp
		want  object.ObjectType
	}{
		{
			name:  "type",
			field: api.V0042OpenapiDiagResp{},
			want:  ObjectTypeV0042Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Diag{V0042OpenapiDiagResp: tt.field}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Diag.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Diag_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0042OpenapiDiagResp
		want  object.Object
	}{
		{
			name:  "empty",
			field: api.V0042OpenapiDiagResp{},
			want:  &V0042Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Diag{V0042OpenapiDiagResp: tt.field}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Diag.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042Diag_DeepCopy(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0042OpenapiDiagResp
		want  *V0042Diag
	}{
		{
			name:  "empty",
			field: api.V0042OpenapiDiagResp{},
			want:  &V0042Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042Diag{V0042OpenapiDiagResp: tt.field}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042Diag.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042DiagList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042Diag
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0042Diag{},
			want:  ObjectTypeV0042Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042DiagList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042DiagList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042DiagList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042Diag
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0042Diag{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0042Diag{{V0042OpenapiDiagResp: api.V0042OpenapiDiagResp{}}},
			want:  []object.Object{&V0042Diag{api.V0042OpenapiDiagResp{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042DiagList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042DiagList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0042DiagList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0042Diag
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0042Diag{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0042Diag{},
			arg:        &V0042Diag{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0042Diag{{V0042OpenapiDiagResp: api.V0042OpenapiDiagResp{}}},
			arg:        &V0042Diag{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042DiagList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0042DiagList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0042DiagList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0042Diag
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0042Diag{},
			want:  &V0042DiagList{Items: []V0042Diag{}},
		},
		{
			name:  "existing",
			items: []V0042Diag{{V0042OpenapiDiagResp: api.V0042OpenapiDiagResp{}}},
			want:  &V0042DiagList{Items: []V0042Diag{{api.V0042OpenapiDiagResp{}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0042DiagList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0042DiagList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
