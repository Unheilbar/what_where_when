// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"os/signal"
	"syscall"

	redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/what_where_when"
	"github.com/unheilbar/what_where_when/pkg/handler"
	"github.com/unheilbar/what_where_when/pkg/hub"
	"github.com/unheilbar/what_where_when/pkg/repository"
	"github.com/unheilbar/what_where_when/pkg/service"
)

func main() {
	nhub := hub.NewHub()
	go nhub.Run()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	repo := repository.NewRepository(rdb)

	services := service.NewServices(repo)

	handlers := handler.NewHandler(services)

	// Create & Run HTTP Server
	server := what_where_when.NewServer()
	go func() {
		if err := server.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Errorf("Error occurred while running server: %s\n", err.Error())
		}
	}()

	logrus.Info("Application Started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
}
