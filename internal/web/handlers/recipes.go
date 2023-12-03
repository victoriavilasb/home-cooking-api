package handlers

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
)

const (
	recipesAPIPath = "/api/recipes"
)

func (h *Handler) RegisterRecipe(request *restful.Request, response *restful.Response) {
	var recipe *domain.Recipe
	if err := request.ReadEntity(&recipe); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if err := h.service.RegisterRecipe(request.Request.Context(), *recipe); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	response.WriteHeader(http.StatusCreated)
}
