package main

import (
	"github.com/imperiustx/qna_service/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	service := app.NewApplication()
	service.Start()
}
