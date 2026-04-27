package rpc

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Config(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	path := os.Getenv("GAME_CONFIG_PATH")
	if path == "" {
		path = "/nakama/data/config/game_config.json"
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		logger.Error("Could not read game config file", "path", path, "error", err.Error())
		return "", runtime.NewError("game config unavailable", 13)
	}

	if !json.Valid(bytes) {
		logger.Error("Game config file is not valid JSON", "path", path)
		return "", runtime.NewError("game config is invalid", 13)
	}

	return string(bytes), nil
}
