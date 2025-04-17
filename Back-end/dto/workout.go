package dto

type WorkoutCreateDTO struct {
	Name          string  `json:"name" binding:"required"`
	Day           uint    `json:"day" binding:"required"`
	RepLowerBound *uint   `json:"rep_lower_bound,omitempty"`
	RepUpperBound *uint   `json:"rep_upper_bound,omitempty"`
	Description   *string `json:"description,omitempty"`
	ExerciseID    uint    `json:"exercise_id" binding:"required"`
	ProgramId     uint    `json:"program_id" binding:"required"`
}
