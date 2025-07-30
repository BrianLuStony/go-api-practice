package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example/api/internal/http/handlers"
	"example/api/internal/http/routes"
	"example/api/internal/http/server"
	"example/api/internal/services"
)

func main() {
	fmt.Println("Hello, this is main")

	// Step 1: Create service (dummy for now)
	userService := services.NewUserService() // assuming NewUserService() takes no args

	// Step 2: Create handler
	userHandler := handlers.NewUserHandler(userService)

	// Step 3: Setup router with handler
	router := routes.SetupRouter(userHandler)

	// Step 4: Start HTTP server
	server := server.NewServer(server.ServerConfig{
		Addr:   ":8080",
		Router: router})

	err := server.Start()

	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

	fmt.Println("Server started successfully on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := server.ShutdownGracefully(5 * time.Second); err != nil {
		fmt.Println("Failed to shutdown server gracefully:", err)
	} else {
		fmt.Println("Server shutdown gracefully")
	}
	fmt.Println("Exiting application")
}
