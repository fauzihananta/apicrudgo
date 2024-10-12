package middleware

import (
	"context"
	"log"
	"net/http"
)

type keyHeader string

const (
	USER_ID keyHeader = "x-user-id"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			ctxmap := map[keyHeader]string{}

			parseHeader(ctxmap, r.Header)
			ctx = submitSessionCtx(ctx, ctxmap)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func parseHeader(cm map[keyHeader]string, h http.Header) {
	userID := h.Get(string(USER_ID))
	if userID != "" {
		cm[USER_ID] = userID
	}
}

func submitSessionCtx(ctx context.Context, cm map[keyHeader]string) context.Context {
	userID, exist := cm[USER_ID]
	if exist {
		ctx = context.WithValue(ctx, USER_ID, userID)
	}

	return ctx
}

func GetUserID(ctx context.Context) (string, bool) {
	return getDataContext(ctx, USER_ID)
}

func getDataContext(ctx context.Context, header keyHeader) (string, bool) {
	switch value := ctx.Value(header).(type) {
	case string:
		return value, true
	default:
		log.Println(value)
		return "", false
	}
}
