// Copyright (c) Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause-Clear

package gateway

import (
	"github.com/labstack/echo/v4"

	storage "github.com/foundriesio/dg-satellite/storage/gateway"
)

type handlers struct {
	storage *storage.Storage
}

func RegisterHandlers(e *echo.Echo, storage *storage.Storage) {
	h := handlers{storage: storage}
	e.Use(h.authDevice)
	e.GET("/device", h.deviceGet)
}
