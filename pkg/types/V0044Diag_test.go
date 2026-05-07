// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"testing"

	apiequality "k8s.io/apimachinery/pkg/api/equality"

	api "github.com/SlinkyProject/slurm-client/api/v0044"
	"github.com/SlinkyProject/slurm-client/pkg/object"
)

func TestV0044Diag_GetKey(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0044OpenapiDiagResp
		want  object.ObjectKey
	}{
		{
			name:  "key",
			field: api.V0044OpenapiDiagResp{},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Diag{V0044OpenapiDiagResp: tt.field}
			if got := o.GetKey(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Diag.GetKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Diag_GetType(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0044OpenapiDiagResp
		want  object.ObjectType
	}{
		{
			name:  "type",
			field: api.V0044OpenapiDiagResp{},
			want:  ObjectTypeV0044Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Diag{V0044OpenapiDiagResp: tt.field}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Diag.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Diag_DeepCopyObject(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0044OpenapiDiagResp
		want  object.Object
	}{
		{
			name:  "empty",
			field: api.V0044OpenapiDiagResp{},
			want:  &V0044Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Diag{V0044OpenapiDiagResp: tt.field}
			if got := o.DeepCopyObject(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Diag.DeepCopyObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044Diag_DeepCopy(t *testing.T) {
	tests := []struct {
		name  string
		field api.V0044OpenapiDiagResp
		want  *V0044Diag
	}{
		{
			name:  "empty",
			field: api.V0044OpenapiDiagResp{},
			want:  &V0044Diag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044Diag{V0044OpenapiDiagResp: tt.field}
			if got := o.DeepCopy(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044Diag.DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044DiagList_GetType(t *testing.T) {
	tests := []struct {
		name  string
		items []V0044Diag
		want  object.ObjectType
	}{
		{
			name:  "type",
			items: []V0044Diag{},
			want:  ObjectTypeV0044Diag,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044DiagList{Items: tt.items}
			if got := o.GetType(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044DiagList.GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044DiagList_GetItems(t *testing.T) {
	tests := []struct {
		name  string
		items []V0044Diag
		want  []object.Object
	}{
		{
			name:  "empty",
			items: []V0044Diag{},
			want:  []object.Object{},
		},
		{
			name:  "items",
			items: []V0044Diag{{V0044OpenapiDiagResp: api.V0044OpenapiDiagResp{}}},
			want:  []object.Object{&V0044Diag{api.V0044OpenapiDiagResp{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044DiagList{Items: tt.items}
			if got := o.GetItems(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044DiagList.GetItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestV0044DiagList_AppendItem(t *testing.T) {
	tests := []struct {
		name       string
		items      []V0044Diag
		arg        object.Object
		wantAppend bool
	}{
		{
			name:       "nil",
			items:      []V0044Diag{},
			arg:        nil,
			wantAppend: false,
		},
		{
			name:       "empty",
			items:      []V0044Diag{},
			arg:        &V0044Diag{},
			wantAppend: true,
		},
		{
			name:       "existing",
			items:      []V0044Diag{{V0044OpenapiDiagResp: api.V0044OpenapiDiagResp{}}},
			arg:        &V0044Diag{},
			wantAppend: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044DiagList{Items: tt.items}
			want := len(o.GetItems())
			if tt.wantAppend {
				want++
			}
			o.AppendItem(tt.arg)
			if got := len(o.GetItems()); want != got {
				t.Errorf("V0044DiagList.AppendItem() = %v, want %v", got, want)
			}
		})
	}
}

func TestV0044DiagList_DeepCopyObjectList(t *testing.T) {
	tests := []struct {
		name  string
		items []V0044Diag
		want  object.ObjectList
	}{
		{
			name:  "empty",
			items: []V0044Diag{},
			want:  &V0044DiagList{Items: []V0044Diag{}},
		},
		{
			name:  "existing",
			items: []V0044Diag{{V0044OpenapiDiagResp: api.V0044OpenapiDiagResp{}}},
			want:  &V0044DiagList{Items: []V0044Diag{{api.V0044OpenapiDiagResp{}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &V0044DiagList{Items: tt.items}
			if got := o.DeepCopyObjectList(); !apiequality.Semantic.DeepEqual(got, tt.want) {
				t.Errorf("V0044DiagList.DeepCopyObjectList() = %v, want %v", got, tt.want)
			}
		})
	}
}
