package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type TestMiddleware1Middleware struct {
}

func NewTestMiddleware1Middleware() *TestMiddleware1Middleware {
	return &TestMiddleware1Middleware{}
}

func (m *TestMiddleware1Middleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.WithContext(r.Context()).Infow("查看中间件 testmiddle1", logx.Field("战争", "框架"))
		next(w, r)
	}
}
