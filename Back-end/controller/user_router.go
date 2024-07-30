package controller

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/service"
	"github.com/zYasser/MyFitness/utils"
)

var validate = validator.New()

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println(user.Password)
	if err := user.CreateUser(app.Db); err != nil {
		logger.ErrorLog.Printf("Registration Failed: %v", err)
		status := http.StatusBadRequest
		if err.Error() == "" {
			status = http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status, map[string]string{
			"message": "Error occurred during user registration",
			"error":   err.Error(),
		})
		return
	}

	logger.InfoLog.Println("Finished Inserting")
	utils.RespondWithJSON(w, http.StatusCreated, user)
}
func (app *Application) login(w http.ResponseWriter, r *http.Request){
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
	
	
	err :=service.ValidateUser(app.Db , params)
	if(err !=nil){
		status := http.StatusUnauthorized
		if err.Error() == "" {
			status = http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status, map[string]string{
			"message": "Error occurred during user registration",
			"error":   err.Error(),
		})
		return

	}
	utils.RespondWithJSON(w, http.StatusOK,nil )
	

}