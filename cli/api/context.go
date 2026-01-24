// Copyright (c) Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause-Clear

package api

import "context"

type apiContextKey int

const (
	contextKey apiContextKey = iota
)

func CtxGetApi(ctx context.Context) *Api {
	return ctx.Value(contextKey).(*Api)
}

func CtxWithApi(ctx context.Context, api *Api) context.Context {
	return context.WithValue(ctx, contextKey, api)
}
