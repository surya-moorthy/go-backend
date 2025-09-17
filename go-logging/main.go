package main

import (
	"log/slog"
	"math/rand"
	"os"
)

type User struct {
	ID         int    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
}

func main() {

	LOG_FILE := os.Getenv("LOG_FILE")

	file , err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)

	if err != nil {
	   panic(err)
	}

	defer file.Close()

	handleOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(file,handleOpts))
	slog.SetDefault(logger)

	// userGroup := slog.Group("user", slog.String("user","john Doe"),slog.Int("id",rand.Intn(50)))

	reqGroup := slog.Group(
		"request",
		"method" , "GET",
	)

	requestLogger := logger.With(reqGroup)

	user := &User{rand.Intn(30),"john Doe","password"}
	requestLogger.Info("Oder Succeeded", "user" , user)
	requestLogger.Error("Error occured")
}