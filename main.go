package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-boilerplate-organizer/db"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v3"
)

func startServer(app *fiber.App) {
	log.Fatal(app.Listen(":4001"))
}

func main() {
	ctx := context.Background()
	conn, err := db.InitDB(ctx)
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
	startServer(app)

}
