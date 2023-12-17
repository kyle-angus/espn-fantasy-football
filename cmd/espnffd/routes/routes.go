package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/routes/v1"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	api := r.Group("/api/v1")
	api.GET("/:season/team", v1.GetTeams)
	api.GET("/:season/team/:id", v1.GetTeamById)

	return r
}
