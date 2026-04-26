package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	initializer.RegisterRpc("update_metadata", rpcUpdateMetadata)
	initializer.RegisterRpc("get_config", rpcGetConfig)
	initializer.RegisterRpc("private_ping", rpcPrivatePing)

	logger.Info("RPCs registered")
	return nil
}
