package rpc

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

func Config(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {

	config := map[string]interface{}{
		"welcome_message": "Welcome to the game!",
		"xp_rate":         1.5,
		"rarity_options":  []string{"common", "rare", "epic", "legendary"},
	}

	bytes, _ := json.Marshal(config)
	return string(bytes), nil
}
