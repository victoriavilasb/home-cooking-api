package handlers

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful"

	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/core/ports"
)

const (
	groceriesAPIPath = "/api/groceries"
)

func (h *Handler) RegisterGrocery(request *restful.Request, response *restful.Response) {
	var grocery *domain.Grocery
	if err := request.ReadEntity(&grocery); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if err := h.service.RegisterGrocery(request.Request.Context(), *grocery); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateGrocery(request *restful.Request, response *restful.Response) {
	var grocery *domain.Grocery
	if err := request.ReadEntity(&grocery); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if err := h.service.UpdateGrocery(request.Request.Context(), request.PathParameter("id"), *grocery); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteGrocery(request *restful.Request, response *restful.Response) {
	if err := h.service.DeleteGrocery(request.Request.Context(), request.PathParameter("id")); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteHeader(http.StatusOK)
}

func (h *Handler) RetrieveGroceries(request *restful.Request, response *restful.Response) {
	var filter *ports.GroceryFilter
	request.ReadEntity(&filter)

	groceries, err := h.service.RetrieveGroceries(request.Request.Context(), filter)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if len(groceries) == 0 {
		response.WriteError(http.StatusNotFound, errors.New("no groceries found matching this filter"))
	}

	response.WriteHeaderAndEntity(http.StatusFound, groceries)
}
