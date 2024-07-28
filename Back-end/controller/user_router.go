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
	logger.InfoLog.Println("Validating the Request")

	errs := utils.Validate(params, validate)
		if len(errs) !=0{
			utils.RespondWithJSON(w,http.StatusBadRequest,errs)
			return
		}
	
	logger.InfoLog.Println("Inserting DB")

	user:=mapper.MapParametersToUser(*params)
	err = user.CreateUser(app.Db)

	if(err!=nil){
		logger.ErrorLog.Printf("Registration Failed Error : %v " , err )
		status:=http.StatusBadRequest
		if err.Error()=="" {
			status=http.StatusInternalServerError
		}
		utils.RespondWithJSON(w, status  , struct {
			Message string `json:"message"`
			Error string `json:"error"`
		}{Message: "Error occurred during user registration" , Error: err.Error()})	
			return
	}
		logger.InfoLog.Println("Finished Inserting")
		utils.RespondWithJSON(w,http.StatusCreated,user)


}
