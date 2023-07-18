package controllers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/model"
	"github.com/ramadhanalfarisi/go-concurrency-pipeline/pipeline"
)

func CreatePayload(count int, chanin chan model.Payload) {
	go func() {
		for i := 0; i < count; i++ {
			c := strconv.Itoa(i + 1)
			payload := model.Payload{Title: "Title " + c, Filepath: "logs/file" + c + ".txt", Content: "Content " + c}
			chanin <- payload
		}
		close(chanin)
	}()
}

func UploadController(ctx *fiber.Ctx) error {
	for {
		select {
		case <-ctx.Context().Done():
			return ctx.Context().Err()
		default:
			chanin := make(chan model.Payload, 4)
			CreatePayload(1000, chanin)
			write1 := pipeline.WriteText(chanin)
			write2 := pipeline.WriteText(chanin)
			write3 := pipeline.WriteText(chanin)
			write4 := pipeline.WriteText(chanin)
			resultPath := pipeline.CollectPath(write1, write2, write3, write4)
			json, err := json.Marshal(resultPath)
			if err != nil {
				log.Fatal(err)
			}
			ctx.Set("Content-Type", "application/json")
			ctx.SendStatus(200)
			ctx.SendString(string(json))
			return nil
		}
	}
}
