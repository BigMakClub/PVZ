package dto

import "time"

type OpenReceptionRequest struct {
	PVZID string `json:"pvzId"`
}

type ReceptionResponse struct {
	ID       string    `json:"id"`
	PVZID    string    `json:"pvzId"`
	Status   string    `json:"status"`
	DateTime time.Time `json:"dateTime"`
}
