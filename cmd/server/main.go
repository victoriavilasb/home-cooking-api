package main

import (
	"database/sql"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/victoriavilasb/home-cooking-api/internal/core/service"
	"github.com/victoriavilasb/home-cooking-api/internal/repository"
	"github.com/victoriavilasb/home-cooking-api/internal/web/handlers"
	"go.uber.org/zap"
)

func main() {
	logger := zap.L().Sugar()

	connection := "postgresql://postgres:postgres@localhost:5432/home-cooking?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		db.Close()
		panic(err)
	}
	defer db.Close()

	repo := repository.NewRepositoryClient(db, logger)

	svc, err := service.NewHomeCookingService(repo, logger)
	if err != nil {
		panic(err)
	}

	ws := new(restful.WebService)
	_ = handlers.NewHandler(ws, svc, logger)

	container := restful.NewContainer()
	container.Add(ws)
	if err := http.ListenAndServe(":9251", container); err != nil {
		panic(err)
	}
}
