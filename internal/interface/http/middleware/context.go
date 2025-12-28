package middleware

import "context"

type UserIDKey struct{}

func FromUserIDContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey{}).(int64)
	return userID, ok
}
