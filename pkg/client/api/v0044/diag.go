// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package v0044

import (
	"context"
	"errors"
	"net/http"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/SlinkyProject/slurm-client/pkg/types"
	"github.com/SlinkyProject/slurm-client/pkg/utils"
)

type DiagInterface interface {
	GetDiag(ctx context.Context) (*types.V0044Diag, error)
	ListDiag(ctx context.Context) (*types.V0044DiagList, error)
}

var _ DiagInterface = &SlurmClient{}

// GetDiag implements ClientInterface
func (c *SlurmClient) GetDiag(ctx context.Context) (*types.V0044Diag, error) {
	res, err := c.SlurmV0044GetDiagWithResponse(ctx)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != 200 {
		errs := []error{errors.New(http.StatusText(res.StatusCode()))}
		if res.JSONDefault != nil {
			errs = append(errs, getOpenapiErrors(res.JSONDefault.Errors)...)
		}
		return nil, utilerrors.NewAggregate(errs)
	}
	out := &types.V0044Diag{}
	utils.RemarshalOrDie(res.JSON200, out)
	return out, nil
}

// ListDiag implements ClientInterface
func (c *SlurmClient) ListDiag(ctx context.Context) (*types.V0044DiagList, error) {
	res, err := c.GetDiag(ctx)
	if err != nil {
		return nil, err
	}
	list := &types.V0044DiagList{
		Items: []types.V0044Diag{
			*res,
		},
	}
	return list, nil
}
