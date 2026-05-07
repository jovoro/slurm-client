// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0041

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0041"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type ReservationInterface interface {
	GetReservationInfo(ctx context.Context, name string) (*types.V0041ReservationInfo, error)
	ListReservationInfo(ctx context.Context) (*types.V0041ReservationInfoList, error)
}

var _ ReservationInterface = &SlurmClient{}

// GetReservationInfo implements ClientInterface
func (c *SlurmClient) GetReservationInfo(ctx context.Context, name string) (*types.V0041ReservationInfo, error) {
	params := &api.SlurmV0041GetReservationParams{}
	res, err := c.SlurmV0041GetReservationWithResponse(ctx, name, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	if len(res.JSON200.Reservations) == 0 {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	out := &types.V0041ReservationInfo{}
	utils.RemarshalOrDie(res.JSON200.Reservations[0], out)
	return out, nil
}

// ListReservationInfo implements ClientInterface
func (c *SlurmClient) ListReservationInfo(ctx context.Context) (*types.V0041ReservationInfoList, error) {
	params := &api.SlurmV0041GetReservationsParams{}
	res, err := c.SlurmV0041GetReservationsWithResponse(ctx, params)
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}

	list := &types.V0041ReservationInfoList{
		Items: make([]types.V0041ReservationInfo, len(res.JSON200.Reservations)),
	}
	for i, item := range res.JSON200.Reservations {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
