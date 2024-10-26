package middleware

import (
	"errors"
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

// func AuthorizationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("access_token")
// 				fmt.Println(cookie)
// 				if err != nil {
// 					http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 					return
// 				}
// 		err=utils.VerifyToken(cookie.Value)
// 		if(err!=nil){
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 				}

// 				next.ServeHTTP(w, r)
// 			})
// 		}

func AuthorizationMiddleware(redis *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := FromContext(r.Context())
			logger.InfoLog.Println("Validating JWT Token")
			cookie, err := r.Cookie("access_token")
			if err != nil {
				logger.ErrorLog.Println("No JWT Token")

				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, err := utils.VerifyToken(cookie.Value)
			if(err==nil){
				next.ServeHTTP(w, r)
				return
			}
			if err != nil {
				var e utils.JwtExpireTokenErr
				if !errors.Is(err, e){
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					
					return
				}
			}
			logger.InfoLog.Println("Recreating JWT token via Refresh Token")
			old_refresh, user := claims["refresh"].(string), claims["username"].(string)
			result := service.VerifyRefreshToken(r.Context(), redis, old_refresh, user, logger)
			if result {
				token, refresh, err := utils.CreateToken(user)
				if err != nil {
					logger.ErrorLog.Println("Failed To Create A JWT Token")
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				service.RemoveRefreshToken(r.Context(), redis, old_refresh, user)

				service.CreateRefreshToken(r.Context(), redis, refresh, user)
				cookie := http.Cookie{
					Name:     "access_token",
					Value:    token,
					HttpOnly: true,
					Secure:   true,
					Path:     "/",
				}
				// Use the http.SetCookie() function to send the cookie to the client.
				// Behind the scenes this adds a `Set-Cookie` header to the response
				// containing the necessary cookie data.
				http.SetCookie(w, &cookie)

			} else {
				logger.ErrorLog.Printf("Refresh Token Expired %s" , old_refresh)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return

			}

			next.ServeHTTP(w, r)
		})
	}
}
