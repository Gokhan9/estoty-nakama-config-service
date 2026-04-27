package main

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"nakama-config-service/nakama/modules/rpc"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Main(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {

	initializer.RegisterRpc("update_metadata", rpc.UpdateMetadata)
	initializer.RegisterRpc("config", rpc.Config)
	initializer.RegisterRpc("private_ping", rpc.Ping)

	logger.Info("RPCs registered")
	return nil
}
