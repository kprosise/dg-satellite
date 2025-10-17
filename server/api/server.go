// Copyright (c) Qualcomm Technologies, Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-3-Clause-Clear

package api

import (
	"fmt"
	"time"

	"github.com/foundriesio/dg-satellite/auth"
	"github.com/foundriesio/dg-satellite/server"
	"github.com/foundriesio/dg-satellite/storage"
	"github.com/foundriesio/dg-satellite/storage/api"
)

const serverName = "rest-api"

func NewServer(ctx Context, db *storage.DbHandle, fs *storage.FsHandle, port uint16, authFunc auth.AuthUserFunc) (server.Server, error) {
	strg, err := api.NewStorage(db, fs)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s storage: %w", serverName, err)
	}
	e := server.NewEchoServer()
	srv := server.NewServer(ctx, e, serverName, port, nil)
	RegisterHandlers(e, strg, authFunc)
	return &apiServer{server: srv, daemons: NewDaemons(ctx, strg)}, nil
}

type apiServer struct {
	server  server.Server
	daemons *daemons
}

func (s apiServer) Start(quit chan error) {
	s.daemons.Start()
	s.server.Start(quit)
}

func (s apiServer) Shutdown(timeout time.Duration) {
	s.daemons.Shutdown()
	s.server.Shutdown(timeout)
}

func (s apiServer) GetAddress() string {
	return s.server.GetAddress()
}

func (s apiServer) GetDnsName() string {
	return s.server.GetDnsName()
}
