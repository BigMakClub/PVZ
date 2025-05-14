package dto

import "time"

type AddProductRequest struct {
	Type  string `json:"type"`
	PVZID string `json:"pvzId"`
}

type ProductResponse struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	DateTime time.Time `json:"dateTime"`
}
