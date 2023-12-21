package main

import (
	"log"
	"net/http"
	"websocket-chat/controller"
	"websocket-chat/handlers"
	"websocket-chat/utils"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	

	go handlers.HandleMessages()

	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
	e.POST("/migrate", controller.Migrate)
	e.GET("/ws", handlers.HandleWebSocket)

	e.POST("/create-room", func(c echo.Context) error {
		return handlers.CreateRoom(c, db)
	})

	e.POST("/join-room", handlers.JoinRoom)
	e.GET("/room-list", handlers.RoomList)
	e.GET("/health", healthHandler)
	e.GET("/readiness", readinessHandler)

	// Defer closing the database connection
	e.Logger.Fatal(e.Start(":8080"))
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func readinessHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
