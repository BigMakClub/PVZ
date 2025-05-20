package dto

type CreatePVZRequest struct {
	City string `json:"city" validate:"required,oneof=Москва Санкт-Петербург Казань"`
}

type CreatePVZResponse struct {
	Id               string `json:"id"`
	City             string `json:"city"`
	RegistrationDate string `json:"registrationDate"`
}
