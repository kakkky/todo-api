package repository

import (
	"context"
	"time"
)

// KVSに対するインターフェース
type KvsCommander interface {
	Save(ctx context.Context, duration time.Duration, userID, jwtID string) error
	Load(ctx context.Context, userID string) (string, error)
	Delete(ctx context.Context, userID string) error
}
