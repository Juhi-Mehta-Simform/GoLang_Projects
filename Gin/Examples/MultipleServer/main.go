package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var g errgroup.Group

func Server01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/server1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to server1",
		})
	})
	return e
}

func Server02() http.Handler {
	e := gin.Default()
	e.GET("/server2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to server2",
		})
	})
	return e
}

func main() {
	server1 := http.Server{
		Addr:         ":8080",
		Handler:      Server01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server2 := http.Server{
		Addr:         ":8081",
		Handler:      Server02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server1.ListenAndServe()
	})
	g.Go(func() error {
		return server2.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
