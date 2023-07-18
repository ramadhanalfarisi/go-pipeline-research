package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/controllers"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/model"
)

type App struct {
	Fapp *fiber.App
	DB *sql.DB
}

func(a *App) NewRouter() {
	router := fiber.New()
	modelPerson := model.PersonModel{DB: a.DB}
	insertController := controllers.InsertController{Model: modelPerson}
	router.Post("/upload", controllers.UploadController)
	router.Post("/insert", insertController.InsertController)
	a.Fapp = router
}

func(a *App) Run(){
	a.Fapp.Listen(":8080")
}