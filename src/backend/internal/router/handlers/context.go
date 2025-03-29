package handlers

import (
	"context"

	"backend/internal/model"
)

type ctxKey string

const (
	userKey ctxKey = "user"
)

func UserToCtx(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func MustUserFromCtx(ctx context.Context) *model.User {
	user, ok := ctx.Value(userKey).(*model.User)
	if !ok {
		panic("user from ctx")
	}

	return user
}
