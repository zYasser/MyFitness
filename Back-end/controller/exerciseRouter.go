package controller

import (
	"net/http"

	"github.com/zYasser/MyFitness/dto"
	"github.com/zYasser/MyFitness/mapper"
	"github.com/zYasser/MyFitness/middleware"
	"github.com/zYasser/MyFitness/utils"
)

func (app *Application) createExercise(w http.ResponseWriter , r *http.Request) {
	con:= r.Context()
	logger:=middleware.FromContext(con)
	logger.InfoLog.Println("Received Create Exercise Request")
	var params dto.Exercise
	if err:=utils.FromJSON(&params , r.Body); err!=nil{
		logger.ErrorLog.Printf("Failed to serialize this object:%v\n", r.Body)
		utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return

	}
	if errs := utils.Validate(&params, validate); len(errs) != 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, errs)
		return
	}
	
	exercise:= mapper.DtoToExercise(params)
	exercise.InsertExercise(app.Db , logger)
	utils.RespondWithJSON(w,http.StatusCreated , exercise)

}