package api

import (
	"duval/internal/pkg/media"
	"duval/internal/pkg/message"
	"duval/internal/pkg/user"
	"github.com/gin-gonic/gin"
)

// routes defines a list of gin function that append the /api group.
var routes = []func(g *gin.RouterGroup){
	userRoutes,
	authRoutes,
	fileRoutes,
	chatRoutes,
}

func authRoutes(g *gin.RouterGroup) {
	g.POST("/login", user.Login)
	g.POST("/password", user.NewPassword)
	g.GET("/password", user.GetUserPasswordHistory)
	return
}

func userRoutes(g *gin.RouterGroup) {
	g.POST("/user", user.NewUser)
	g.PUT("/user", user.UpdateUser)
	g.GET("/user/:id", user.GetUser)
	g.GET("/user", user.MyProfile)
	g.GET("/profileImage", media.ProfileImage)
	return
}

func chatRoutes(g *gin.RouterGroup) {
	g.GET("/chat", message.GetUserDiscussion)
	return
}

func fileRoutes(g *gin.RouterGroup) {
	g.POST("/upload", media.Upload)
	g.Static("/public", "public")
	return
}
