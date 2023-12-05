package handlers

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful"
)

const (
	recommendationsAPIPath = "/api/recommendations"
)

func (h *Handler) RetrieveRecommendations(request *restful.Request, response *restful.Response) {
	recommendations, err := h.service.RetrieveRecommendations(request.Request.Context(), nil)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}

	if len(recommendations) == 0 {
		response.WriteError(http.StatusNotFound, errors.New("no recommendations found matching this filter"))
	}

	response.WriteHeaderAndEntity(http.StatusOK, recommendations)
}
