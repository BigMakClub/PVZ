package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pvz/iternal/adapter/http/dto"
	"pvz/iternal/adapter/transport"
	"pvz/iternal/domain"
	"pvz/iternal/usecase"
)

type PVZHandler struct {
	svc usecase.PVZService
}

func NewPVZHandler(svc usecase.PVZService) *PVZHandler {
	return &PVZHandler{svc: svc}
}

func (h *PVZHandler) Register(r chi.Router) {
	r.Post("/pvz", h.create)
	r.Get("/pvz", h.list)
}

func (h *PVZHandler) create(w http.ResponseWriter, r *http.Request) {
	var in dto.CreatePVZRequest
	if err := transport.DecodeAndValidate(*r, &in); err != nil {
		respondError(w, http.StatusBadRequest, err)
		return
	}

	pvz, err := h.svc.Create(r.Context(), domain.City(in.City))

	if err != nil {
		respondError(w, http.StatusBadRequest, err)
	}

	out := dto.CreatePVZResponse{
		Id:               pvz.Id.String(),
		City:             in.City,
		RegistrationDate: pvz.RegistrationDate.String(),
	}
	respondJSON(w, http.StatusCreated, out)
}

func (h *PVZHandler) list(w http.ResponseWriter, r *http.Request) {

}
