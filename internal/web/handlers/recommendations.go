package handlers

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/victoriavilasb/home-cooking-api/internal/core/ports"
)

const (
	recommendationsAPIPath = "/api/recommendations"
)

func (h *Handler) RetrieveRecommendations(request *restful.Request, response *restful.Response) {
	var filter *ports.RecommendationFilter
	if err := request.ReadEntity(&filter); err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	recommendations, err := h.service.RetrieveRecommendations(request.Request.Context(), filter)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if len(recommendations) == 0 {
		response.WriteError(http.StatusNotFound, errors.New("no recommendations found matching this filter"))
	}

	response.WriteHeaderAndEntity(http.StatusFound, recommendations)
}
