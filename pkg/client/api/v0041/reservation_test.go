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

func TestSlurmClient_GetReservationInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0041ReservationInfo
		wantErr bool
	}{
		{
			name: "Not Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Found",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
							return &api.SlurmV0041GetReservationResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0041OpenapiReservationResp{
									Reservations: []api.V0041ReservationInfo{
										{Name: ptr.To("reservation-0")},
									},
								},
							}, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
			},
			want: &types.V0041ReservationInfo{
				V0041ReservationInfo: api.V0041ReservationInfo{
					Name: ptr.To("reservation-0"),
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
							return &api.SlurmV0041GetReservationResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0041OpenapiReservationResp{
									Errors: &[]api.V0041OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationWithResponse: func(ctx context.Context, reservationName string, params *api.SlurmV0041GetReservationParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx:  context.Background(),
				name: "reservation-0",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.GetReservationInfo(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.GetReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.GetReservationInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlurmClient_ListReservationInfo(t *testing.T) {
	type fields struct {
		ClientWithResponsesInterface api.ClientWithResponsesInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.V0041ReservationInfoList
		wantErr bool
	}{
		{
			name: "Empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClient(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0041ReservationInfoList{
				Items: make([]types.V0041ReservationInfo, 0),
			},
			wantErr: false,
		},
		{
			name: "Non-empty list",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
							return &api.SlurmV0041GetReservationsResponse{
								HTTPResponse: &fake.HttpSuccess,
								JSON200: &api.V0041OpenapiReservationResp{
									Reservations: []api.V0041ReservationInfo{
										{Name: ptr.To("reservation-0")},
										{Name: ptr.To("reservation-1")},
										{Name: ptr.To("reservation-2")},
									},
								},
							}, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: &types.V0041ReservationInfoList{
				Items: []types.V0041ReservationInfo{
					{V0041ReservationInfo: api.V0041ReservationInfo{Name: ptr.To("reservation-0")}},
					{V0041ReservationInfo: api.V0041ReservationInfo{Name: ptr.To("reservation-1")}},
					{V0041ReservationInfo: api.V0041ReservationInfo{Name: ptr.To("reservation-2")}},
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP Status != 200",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
							return &api.SlurmV0041GetReservationsResponse{
								HTTPResponse: &http.Response{
									Status:     http.StatusText(http.StatusInternalServerError),
									StatusCode: http.StatusInternalServerError,
								},
								JSONDefault: &api.V0041OpenapiReservationResp{
									Errors: &[]api.V0041OpenapiError{
										{Error: ptr.To("error 1")},
										{Error: ptr.To("error 2")},
									},
								},
							}, nil
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "HTTP Error",
			fields: fields{
				ClientWithResponsesInterface: fake.NewFakeClientBuilder().
					WithInterceptorFuncs(interceptor.Funcs{
						SlurmV0041GetReservationsWithResponse: func(ctx context.Context, params *api.SlurmV0041GetReservationsParams, reqEditors ...api.RequestEditorFn) (*api.SlurmV0041GetReservationsResponse, error) {
							return nil, errors.New(http.StatusText(http.StatusBadGateway))
						},
					}).
					Build(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &SlurmClient{
				ClientWithResponsesInterface: tt.fields.ClientWithResponsesInterface,
			}
			got, err := c.ListReservationInfo(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlurmClient.ListReservationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlurmClient.ListReservationInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
