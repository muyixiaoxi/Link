package middleware

import (
	context2 "context"
	"net/http"
	"strings"
	"tag/common/jwt"
	"tag/common/response"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (jm *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqContext := r.Context()
		// JWT登录验证的中间件
		// 默认 双Token放在请求头Authorization的Bearer中 , 并以空格隔开
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.Response(w, nil, response.CodeTokenIsEmpty)
			return
		}
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Response(w, nil, response.CodeTokenFormat)
			return
		}
		jwtUtil := jwt.InitNewJWTUtils()
		parseToken, isUpd, err := jwtUtil.ParseAccessToken(parts[1])
		if err != nil {
			response.Response(w, nil, response.CodeTokenInvalid)
			return
		}
		if isUpd {
			response.Response(w, nil, response.CodeTokenPast)
			return
		}
		context := context2.WithValue(reqContext, jwt.UserId, parseToken.UserId)
		newReq := r.WithContext(context)
		next(w, newReq)
	}
}
