package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
)

func GetMembers(ctx *gin.Context) {
	season, err := strconv.ParseUint(ctx.Param("season"), 0, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error":   err.Error(),
			"message": "Unable to parse season",
		})
		return
	}

	espnclient := ctx.MustGet("espn").(espnclient.EspnClient)
	members, err := espnclient.GetMembers(uint(season))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"error":   err.Error(),
			"message": "Unable to fetch members",
		})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"season":  season,
		"members": members,
	})
}

func GetTeamByMemberId(ctx *gin.Context) {
	season, err := strconv.ParseInt(ctx.Param("season"), 0, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error":   err.Error(),
			"message": "Unable to parse season",
		})

		return
	}
	id := ctx.Param("id")
	espnclient := ctx.MustGet("espn").(espnclient.EspnClient)
	team, err := espnclient.GetTeamByMemberId(uint(season), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"error":   err.Error(),
			"message": "Unable to fetch team",
		})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"team":    team,
	})
}
