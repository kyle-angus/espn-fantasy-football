package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/options"
	v1 "github.com/kyle-angus/espn-fantasy-football-client/cmd/espnffd/routes/v1"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
)

func Setup(opt *options.Options) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	api := r.Group("/api/v1", EspnMiddleware(*espnclient.New(opt.LeagueId, opt.EspnS2, opt.SwID)))
	api.GET("/:season/team", v1.GetTeams)
	api.GET("/:season/team/:id", v1.GetTeamById)
	api.GET("/:season/member", v1.GetMembers)
	api.GET("/:season/member/:id/team", v1.GetTeamByMemberId)

	return r
}

func EspnMiddleware(client espnclient.EspnClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("espn", client)
	}
}
