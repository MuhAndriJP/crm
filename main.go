package main

import (
	"log"
	"os"

	postgre "github.com/MuhAndriJP/crm/config"
	"github.com/MuhAndriJP/crm/helper"
	"github.com/MuhAndriJP/crm/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	helper.InitEnv()
	postgre.Init()

	app := fiber.New()
	route.New(app)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
