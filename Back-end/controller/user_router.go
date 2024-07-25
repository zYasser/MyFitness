package controller

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/utils"
)

var validate = validator.New()

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	logger.InfoLog.Println("Received Register Request")
	params:=&dto.User{}
	err:=utils.FromJSON(params,r.Body)
	if(err!=nil){
		logger.ErrorLog.Printf("Failed to decode to json error: %v \n" ,err )	
	}

	errs := utils.Validate(params, validate)
		if errs != nil {
			utils.RespondWithJSON(w,http.StatusBadRequest,errs)
			return
		}
	
	user:=mapper.MapParametersToUser(*params)
	err = user.CreateUser(app.Db)
	if(err!=nil){
		logger.ErrorLog.Printf("Registration Failed Error : %v " , err )
		utils.RespondWithJSON(w, http.StatusBadRequest, struct {
			Message string `json:"message"`
			Error string `json:"error"`
		}{Message: "Error occurred during user registration" , Error: err.Error()})	}

	utils.RespondWithJSON(w,http.StatusCreated,user)


}
