package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

var validate = validator.New()

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	con:= r.Context()
	logger:=middleware.FromContext(con)
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
	if err := user.CreateUser(app.Db, logger); err != nil {
		logger.ErrorLog.Printf("Registration Failed: %v", err)
		status := http.StatusBadRequest
		if err.Error() == "" {
			status = http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status, map[string]string{
			"error":   err.Error(),
		})
		return
	}

	logger.InfoLog.Println("Finished Inserting")
	utils.RespondWithJSON(w, http.StatusCreated, user)
}
func (app *Application) login(w http.ResponseWriter, r *http.Request){
	con:= r.Context()
	logger :=middleware.FromContext(con)

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
	
	
	err := service.ValidateUser(app.Db , params ,logger)
	if(err !=nil){
		status := http.StatusUnauthorized
		if err.Error() == "" {
			status = http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status, map[string]string{
			"error":   err.Error(),
		})
		return

	}
	token,err :=utils.CreateToken(params.Username, logger)
	if(err !=nil){
		logger.ErrorLog.Println("Failed To Create A JWT Token")
	}
	if(err==nil){
		cookie := http.Cookie{
			Name:     "access_token",
			Value:    fmt.Sprintf("bearer %s" , token),
			HttpOnly: true,
			Secure:   true,
		}
		// Use the http.SetCookie() function to send the cookie to the client.
		// Behind the scenes this adds a `Set-Cookie` header to the response
		// containing the necessary cookie data.
		http.SetCookie(w, &cookie)

	
	}


	utils.RespondWithJSON(w, http.StatusOK,nil )
	

}

