package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

var validate = validator.New()

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	app.Logger.InfoLog.Println("Received Register Request")

	var params dto.User
	if err := utils.FromJSON(&params, r.Body); err != nil {
		app.Logger.ErrorLog.Printf("Failed to decode JSON: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	app.Logger.InfoLog.Println("Validating the Request")

	// Validate user DTO
	if errs := utils.Validate(&params, validate); len(errs) != 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, errs)
		return
	}

	app.Logger.InfoLog.Println("Inserting into DB")

	// Map DTO to user model and create user
	user := mapper.MapUserDtoToUser(params)
	if err := service.CreateUser(user, app.Db, app.Logger); err != nil {
		app.Logger.ErrorLog.Printf("Registration Failed: %v", err)
		status := http.StatusBadRequest
		if err.Error() == "" {
			status = http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status, map[string]string{
			"error": err.Error(),
		})
		return
	}

	app.Logger.InfoLog.Println("Finished Inserting")
	utils.RespondWithJSON(w, http.StatusCreated, user)
}
func (app *Application) login(w http.ResponseWriter, r *http.Request) {

	app.Logger.InfoLog.Println("Received log in request")
	var params dto.UserLogin
	if err := utils.FromJSON(&params, r.Body); err != nil {
		app.Logger.ErrorLog.Printf("Failed to decode JSON: %v", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}
	// Validate user DTO
	if errs := utils.Validate(&params, validate); len(errs) != 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, errs)
		return
	}

	err := service.ValidateUser(app.Db, params, app.Logger)
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
		app.Logger.ErrorLog.Println("Failed To Create A JWT Token")
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Something Went Wrong Please Try Again ",
		})
		return

	}

	service.CreateRefreshToken(r.Context(), app.Redis, refresh, params.Username)
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
