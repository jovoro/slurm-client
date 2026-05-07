// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0042

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	api "github.com/SlinkyProject/slurm-client/api/v0042"
	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type ReservationInterface interface {
	GetReservationInfo(ctx context.Context, name string) (*types.V0042ReservationInfo, error)
	ListReservationInfo(ctx context.Context) (*types.V0042ReservationInfoList, error)
}

var _ ReservationInterface = &SlurmClient{}

// GetReservationInfo implements ClientInterface
func (c *SlurmClient) GetReservationInfo(ctx context.Context, name string) (*types.V0042ReservationInfo, error) {
	params := &api.SlurmV0042GetReservationParams{}
	res, err := c.SlurmV0042GetReservationWithResponse(ctx, name, params)
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

	out := &types.V0042ReservationInfo{}
	utils.RemarshalOrDie(res.JSON200.Reservations[0], out)
	return out, nil
}

// ListReservationInfo implements ClientInterface
func (c *SlurmClient) ListReservationInfo(ctx context.Context) (*types.V0042ReservationInfoList, error) {
	params := &api.SlurmV0042GetReservationsParams{}
	res, err := c.SlurmV0042GetReservationsWithResponse(ctx, params)
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

	list := &types.V0042ReservationInfoList{
		Items: make([]types.V0042ReservationInfo, len(res.JSON200.Reservations)),
	}
	for i, item := range res.JSON200.Reservations {
		utils.RemarshalOrDie(item, &list.Items[i])
	}
	return list, nil
}
