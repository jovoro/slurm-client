// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0041/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0041/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_GetDiag(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0041Diag
		wantErr bool
	}{
		{
			name: "Fetch",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
						return &api.SlurmV0041GetDiagResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200:      &api.V0041OpenapiDiagResp{},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0041Diag{},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
						return &api.SlurmV0041GetDiagResponse{
							HTTPResponse: &http.Response{
								Status:     http.StatusText(http.StatusInternalServerError),
								StatusCode: http.StatusInternalServerError,
							},
							JSONDefault: &api.V0041OpenapiDiagResp{
								Errors: &[]api.V0041OpenapiError{
									{Error: ptr.To("error 1")},
								},
							},
						}, nil
					},
				}).
				Build(),
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
						return nil, errors.New(http.StatusText(http.StatusBadGateway))
					},
				}).
				Build(),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			got, err := c.GetDiag(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetDiag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetDiag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListDiag(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0041DiagList
		wantErr bool
	}{
		{
			name: "Fetch",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
						return &api.SlurmV0041GetDiagResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200:      &api.V0041OpenapiDiagResp{},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0041DiagList{Items: []types.V0041Diag{{}}},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0041GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetDiagResponse, error) {
						return nil, errors.New(http.StatusText(http.StatusBadGateway))
					},
				}).
				Build(),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{ClientWithResponsesInterface: tt.client}
			got, err := c.ListDiag(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListDiag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListDiag() = %v, want %v", got, tt.want)
			}
		})
	}
}
