package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/model"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/pipeline"
)

type InsertController struct {
	Model model.PersonModel
}

func (c *InsertController) InsertController(ctx *fiber.Ctx) error {
	for {
		select {
		case <-ctx.Context().Done():
			return ctx.Context().Err()
		default:
			persons := pipeline.CreatePerson(100)
			err1 := c.Model.InsertPerson(persons)
			err2 := c.Model.InsertPerson(persons)
			err3 := c.Model.InsertPerson(persons)
			err4 := c.Model.InsertPerson(persons)
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				log.Fatal(err1)
				log.Fatal(err2)
				log.Fatal(err3)
				log.Fatal(err4)
			}
			ctx.SendStatus(200)
			return nil
		}
	}
}
