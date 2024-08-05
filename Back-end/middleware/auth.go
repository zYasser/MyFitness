package middleware

import (
	"net/http"

	"github.com/zYasser/MyFitness/utils"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("access_token")
		if(err!=nil){
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		err=utils.VerifyToken(cookie.Value)
		if(err!=nil){
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return

		}
		
		next.ServeHTTP(w,r)
	})
}
