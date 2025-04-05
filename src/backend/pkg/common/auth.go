package common

import (
	"context"
)

type ctxKey string

const (
	userIDKey  ctxKey = "userID"
	adminIDKey ctxKey = "adminID"
)

func UserIDToCtx(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func MustUserIDFromCtx(ctx context.Context) int {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		panic("userID from ctx")
	}

	return userID
}

func CtxContainUserID(ctx context.Context) bool {
	_, ok := ctx.Value(userIDKey).(int)
	return ok
}

func AdminIDToCtx(ctx context.Context, adminID int) context.Context {
	return context.WithValue(ctx, adminIDKey, adminID)
}

func MustAdminIDFromCtx(ctx context.Context) int {
	adminID, ok := ctx.Value(adminIDKey).(int)
	if !ok {
		panic("adminID from ctx")
	}

	return adminID
}

func CtxContainAdminID(ctx context.Context) bool {
	_, ok := ctx.Value(adminIDKey).(int)
	return ok
}
