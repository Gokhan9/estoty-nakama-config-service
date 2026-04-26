package rpc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func UpdateMetadata(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return "", runtime.NewError("No user ID", 3)
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &metadata); err != nil {
		return "", runtime.NewError("Invalid JSON", 3)
	}

	err := nk.AccountUpdateId(ctx, userID, "", metadata, "", "", "", "")
	if err != nil {
		return "", err
	}

	return `{"status":"ok"}`, nil
}
