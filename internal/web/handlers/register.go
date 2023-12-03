package handlers

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/victoriavilasb/home-cooking-api/internal/core/domain"
	"github.com/victoriavilasb/home-cooking-api/internal/core/ports"
	"go.uber.org/zap"
)

type Handler struct {
	logger  *zap.SugaredLogger
	service ports.HomeCookingServiceProvider
}

func NewHandler(server *restful.WebService, service ports.HomeCookingServiceProvider, logger *zap.SugaredLogger) Handler {
	handler := Handler{
		logger:  logger,
		service: service,
	}

	server.Route(
		server.POST(groceriesAPIPath).
			Consumes(restful.MIME_JSON).
			Reads(domain.Grocery{}).
			To(handler.RegisterGrocery).
			Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	server.Route(
		server.PUT(groceriesAPIPath).
			Param(restful.PathParameter("id", "external identifier of the grocery").DataType("string")).
			Consumes(restful.MIME_JSON).
			Reads(domain.Grocery{}).
			To(handler.UpdateGrocery).
			Returns(http.StatusOK, http.StatusText(http.StatusOK), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	server.Route(
		server.DELETE(groceriesAPIPath).
			Param(restful.PathParameter("id", "external identifier of the grocery").DataType("string")).
			To(handler.DeleteGrocery).
			Returns(http.StatusOK, http.StatusText(http.StatusOK), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	server.Route(
		server.GET(groceriesAPIPath).
			Consumes(restful.MIME_JSON).
			Reads(ports.GroceryFilter{}).
			To(handler.RetrieveGroceries).
			Produces(restful.MIME_JSON).
			Returns(http.StatusFound, http.StatusText(http.StatusFound), []domain.Grocery{}).
			Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	server.Route(
		server.GET(recommendationsAPIPath).
			Consumes(restful.MIME_JSON).
			Reads(ports.RecommendationFilter{}).
			To(handler.RetrieveRecommendations).
			Produces(restful.MIME_JSON).
			Returns(http.StatusFound, http.StatusText(http.StatusFound), []domain.Recommendation{}).
			Returns(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	server.Route(
		server.POST(recipesAPIPath).
			Consumes(restful.MIME_JSON).
			Reads(domain.Recipe{}).
			To(handler.RegisterRecipe).
			Returns(http.StatusCreated, http.StatusText(http.StatusCreated), nil).
			Returns(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil))

	return handler
}
