package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"html/template"
	//"strings"
	"github.com/onspeedhp/golang_app/database"
	"github.com/onspeedhp/golang_app/globals"
	"github.com/onspeedhp/golang_app/middleware"
	"github.com/onspeedhp/golang_app/routes"
)

func main() {
	router := gin.Default()
	database.Connect()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	routes.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("localhost:8080")
}