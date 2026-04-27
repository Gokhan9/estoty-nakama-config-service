package rpc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func UpdateMetadata(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok || userID == "" {
		return "", runtime.NewError("authentication required", 7)
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &metadata); err != nil {
		return "", runtime.NewError("payload must be a JSON object", 3)
	}
	if metadata == nil {
		return "", runtime.NewError("payload must be a JSON object", 3)
	}

	account, err := nk.AccountGetId(ctx, userID)
	if err != nil {
		return "", err
	}

	merged := make(map[string]interface{})
	if account != nil && account.GetUser() != nil && account.GetUser().GetMetadata() != "" {
		if err := json.Unmarshal([]byte(account.GetUser().GetMetadata()), &merged); err != nil {
			logger.Warn("Could not parse existing user metadata, replacing it", "user_id", userID, "error", err.Error())
			merged = make(map[string]interface{})
		}
	}
	for key, value := range metadata {
		merged[key] = value
	}

	err = nk.AccountUpdateId(ctx, userID, "", merged, "", "", "", "", "")
	if err != nil {
		return "", err
	}

	response, err := json.Marshal(map[string]interface{}{
		"status":   "ok",
		"metadata": merged,
	})
	if err != nil {
		return "", err
	}

	return string(response), nil
}
