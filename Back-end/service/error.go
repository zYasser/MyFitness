package service



type ServiceError struct {
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
}

func (err ServiceError) Error() string {
	return err.Message
}
