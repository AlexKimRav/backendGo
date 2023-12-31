package main

import (
	"backendgo/config"
	"backendgo/controller"
	_ "backendgo/docs"
	"backendgo/helper"
	"backendgo/model"
	"backendgo/repository"
	"backendgo/router"
	"backendgo/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title           Tag Example API
// @version         1.0
// @description     This is a sample server By go gin gorm

// @host      localhost:8888
// @BasePath  /api
func main() {
	log.Info().Msg("start server")
	db := config.ConnectToDb()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	//adding rep,service,ctrl constructors to main
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
