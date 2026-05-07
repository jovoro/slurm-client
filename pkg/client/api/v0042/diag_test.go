// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"k8s.io/utils/ptr"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/fake"
	"github.com/SlinkyProject/slurm-client/pkg/client/api/v0042/interceptor"
	"github.com/SlinkyProject/slurm-client/pkg/types"
)

func TestSlurmClient_GetDiag(t *testing.T) {
	tests := []struct {
		name    string
		client  api.ClientWithResponsesInterface
		want    *types.V0042Diag
		wantErr bool
	}{
		{
			name: "Fetch",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
						return &api.SlurmV0042GetDiagResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200:      &api.V0042OpenapiDiagResp{},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0042Diag{},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
						return &api.SlurmV0042GetDiagResponse{
							HTTPResponse: &http.Response{
								Status:     http.StatusText(http.StatusInternalServerError),
								StatusCode: http.StatusInternalServerError,
							},
							JSONDefault: &api.V0042OpenapiDiagResp{
								Errors: &[]api.V0042OpenapiError{
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
					SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
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
		want    *types.V0042DiagList
		wantErr bool
	}{
		{
			name: "Fetch",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
						return &api.SlurmV0042GetDiagResponse{
							HTTPResponse: &fake.HttpSuccess,
							JSON200:      &api.V0042OpenapiDiagResp{},
						}, nil
					},
				}).
				Build(),
			want:    &types.V0042DiagList{Items: []types.V0042Diag{{}}},
			wantErr: false,
		},
		{
			name: "HTTP Error",
			client: fake.NewFakeClientBuilder().
				WithInterceptorFuncs(interceptor.Funcs{
					SlurmV0042GetDiagWithResponse: func(ctx context.Context, reqEditors ...api.RequestEditorFn) (*api.SlurmV0042GetDiagResponse, error) {
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
