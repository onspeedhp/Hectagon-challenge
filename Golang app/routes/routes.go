package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/onspeedhp/golang_app/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	g.GET("/", controllers.IndexGetHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardGetHandler())
	g.GET("/logout", controllers.LogoutGetHandler())
	g.GET("/repository/:repositoryName", controllers.RepositoryGetHandler())
	g.GET("/author/:authorName", controllers.AuthorGetHandler())
}