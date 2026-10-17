package middleware

import "context"

type contextKey struct{}

var userIdKey contextKey

func UserId(ctx context.Context) (uint64, bool) {
	userId, ok := ctx.Value(userIdKey).(uint64)
	return userId, ok
}
