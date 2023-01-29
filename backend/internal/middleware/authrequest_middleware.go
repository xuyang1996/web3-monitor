package middleware

import (
	"backend/internal/util"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type AuthRequestMiddleware struct {
}

func NewAuthRequestMiddleware() *AuthRequestMiddleware {
	return &AuthRequestMiddleware{}
}

func (m *AuthRequestMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		account := r.Header.Get("account")
		signature := r.Header.Get("signature")
		if len(account) == 0 || len(signature) == 0 {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest,
				util.NewErrorResponseByCode(util.FailToGetSignatureOrAccount))
			return
		}
		if !util.VerifySignature(account, signature) {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusUnauthorized,
				util.NewErrorResponseByCode(util.FailToVerifySignature))
			return
		}
		next(w, r)
	}
}
