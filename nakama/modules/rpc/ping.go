package rpc

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Ping(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	if userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string); ok && userID != "" {
		return "", runtime.NewError("server-to-server key required", 7)
	}

	logger.Info("Private RPC called successfully")

	return "", nil
}
