package controller

import (
	"github.com/redis/go-redis/v9"
	"github.com/zYasser/MyFitness/internal/user/dto"
	"github.com/zYasser/MyFitness/internal/user/mapper"
	"github.com/zYasser/MyFitness/internal/user/service"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/utils"
	"gorm.io/gorm"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func register(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		con := r.Context()
		logger := middleware.FromContext(con)
		logger.InfoLog.Println("Received Register Request")

		// Decode request body to user DTO
		var params dto.User
		if err := utils.FromJSON(&params, r.Body); err != nil {
			logger.ErrorLog.Printf("Failed to decode JSON: %v", err)
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		logger.InfoLog.Println("Validating the Request")

		// Validate user DTO
		if errs := utils.Validate(&params, validate); len(errs) != 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		logger.InfoLog.Println("Inserting into DB")

		// Map DTO to user model and create user
		user := mapper.MapUserDtoToUser(params)
		if err := user.CreateUser(db, logger); err != nil {
			logger.ErrorLog.Printf("Registration Failed: %v", err)
			status := http.StatusBadRequest
			if err.Error() == "" {
				status = http.StatusInternalServerError
			}
			utils.RespondWithJSON(w, status, map[string]string{
				"error": err.Error(),
			})
			return
		}

		logger.InfoLog.Println("Finished Inserting")
		utils.RespondWithJSON(w, http.StatusCreated, user)

	}
}
func login(db *gorm.DB, redis *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		con := r.Context()
		logger := middleware.FromContext(con)

		logger.InfoLog.Println("Received log in request")
		var params dto.UserLogin
		if err := utils.FromJSON(&params, r.Body); err != nil {
			logger.ErrorLog.Printf("Failed to decode JSON: %v", err)
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}
		// Validate user DTO
		if errs := utils.Validate(&params, validate); len(errs) != 0 {
			utils.RespondWithJSON(w, http.StatusBadRequest, errs)
			return
		}

		err := service.Validate(db, params, logger)
		if err != nil {
			status := http.StatusUnauthorized
			if err.Error() == "" {
				status = http.StatusInternalServerError
			}
			utils.RespondWithJSON(w, status, map[string]string{
				"error": err.Error(),
			})
			return

		}
		token, refresh, err := utils.CreateToken(params.Username)
		if err != nil {
			logger.ErrorLog.Println("Failed To Create A JWT Token")
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "Something Went Wrong Please Try Again ",
			})
			return

		}

		service.CreateRefreshToken(con, redis, refresh, params.Username)
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

		utils.RespondWithJSON(w, http.StatusOK, nil)

	}
}
