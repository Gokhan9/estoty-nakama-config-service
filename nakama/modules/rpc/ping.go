package rpc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Ping(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	var req map[string]string
	if err := json.Unmarshal([]byte(payload), &req); err != nil {
		return "", runtime.NewError("Invalid payload", 3)
	}

	if req["secret"] != "my-super-secret-key" {
		return "", runtime.NewError("Forbidden", 7)
	}

	logger.Info("Private RPC called successfully")

	return "", nil
}
