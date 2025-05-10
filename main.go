package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-template-boilerplate/cmd/routes"
	"github.com/go-template-boilerplate/db"
	_ "github.com/go-template-boilerplate/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func startServer(app *fiber.App) {
	log.Fatal(app.Listen(":4001"))
}

// @title			Order Api
// @version		1.0
// @description	This is an Boilerplate for Backend
// @termsOfService	http://swagger.io/terms/
func main() {
	ctx := context.Background()
	conn, queries, err := db.InitDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	fmt.Println(`
	 $$$$$$\   $$$$$$\         $$$$$$\  $$$$$$$\ $$$$$$\ 
	$$  __$$\ $$  __$$\       $$  __$$\ $$  __$$\\_$$  _|
	$$ /  \__|$$ /  $$ |      $$ /  $$ |$$ |  $$ | $$ |  
	$$ |$$$$\ $$ |  $$ |      $$$$$$$$ |$$$$$$$  | $$ |  
	$$ |\_$$ |$$ |  $$ |      $$  __$$ |$$  ____/  $$ |  
	$$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |       $$ |  
	\$$$$$$  | $$$$$$  |      $$ |  $$ |$$ |     $$$$$$\ 
	\______/  \______/       \__|  \__|\__|     \______|
	`)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))
	routes.Routes(app, ctx, queries)
	startServer(app)

}
