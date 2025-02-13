package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	// local repos
	er "products/errors"
	model "products/models"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AsampleService interface{
	CreateASampleByAppId(ctx context.Context,appID string, model *model.Asamplemodel) error 
	GetASampleByAppId(ctx context.Context, appID string, id string) (*model.Asamplemodel,error)
	UpdateASampleByAppId(ctx context.Context, appID string, model *model.Asamplemodel) error
	DeleteASampleByAppId(ctx context.Context, appID string, id string) error
}
type ASampleHandler struct{
	service AsampleService
}

func NewAServiceHandler(service AsampleService) *ASampleHandler{
	return &ASampleHandler{service: service}
}


func (h *ASampleHandler) CreateAsampleByAppId(w http.ResponseWriter, r *http.Request) (response any, status int, err error) {
	appID := r.URL.Query().Get("appID")
	if appID == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "app_id can't be empty")
	}

	var model model.Asamplemodel
	model.ID = uuid.NewString()
	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
	return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "Invalid request body")
	}

	err = h.service.CreateASampleByAppId(context.Background(), appID, &model)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "Created successfully"}, http.StatusOK, nil
}




func (h *ASampleHandler) GetAsampleByAppId(w http.ResponseWriter, r *http.Request) (response any, status int, err error) {
	appID := r.URL.Query().Get("appID")
	id := chi.URLParam(r,"id")
	if id == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "ID can't be empty")
	}

	if appID == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "app_id can't be empty")
	}

	asample, err := h.service.GetASampleByAppId(context.Background(), appID, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return asample, http.StatusOK, nil
}


func (h *ASampleHandler) UpdateAsampleByAppId(w http.ResponseWriter, r *http.Request) (response any, status int, err error) {
	id := chi.URLParam(r,"id")
	if id == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "ID can't be empty")
	}

	appID := r.URL.Query().Get("appID")
	if appID == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "app_id can't be empty")
	}

	var model model.Asamplemodel
	err = json.NewDecoder(r.Body).Decode(&model)
	model.ID = id
	if err != nil {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "Invalid request body")
	}

	err = h.service.UpdateASampleByAppId(context.Background(), appID, &model)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "Updated successfully"}, http.StatusOK, nil
}

func (h *ASampleHandler) DeleteAsampleByAppId(w http.ResponseWriter, r *http.Request) (response any, status int, err error) {
	id := chi.URLParam(r,"id")
	if id == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "ID can't be empty")
	}

	appID := r.URL.Query().Get("appID")
	if appID == "" {
		return nil, http.StatusBadRequest, er.E(er.InvalidParamsErr, "app_id can't be empty")
	}

	err = h.service.DeleteASampleByAppId(context.Background(), appID, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return map[string]string{"message": "Deleted successfully"}, http.StatusOK, nil
}

