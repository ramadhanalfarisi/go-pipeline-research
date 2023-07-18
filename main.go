package main

import "github.com/ramadhanalfarisi/go-concurrency-pipeline/app"

func main() {
	a := app.App{}
	a.ConnectDB()
	a.NewRouter()
	a.Run()
}