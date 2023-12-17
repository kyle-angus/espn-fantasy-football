package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kyle-angus/espn-fantasy-football-client/pkg/espnclient"
)

func GetTeams(ctx *gin.Context) {
	season, err := strconv.ParseUint(ctx.Param("season"), 0, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error":   err.Error(),
			"message": "Unable to parse season",
		})
		return
	}

	espnClient := ctx.MustGet("espn").(espnclient.EspnClient)
	teams, err := espnClient.GetTeams(uint(season))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"message": "Unable to get teams",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"season":  season,
		"teams":   teams,
	})
}

func GetTeamById(ctx *gin.Context) {
	season, err := strconv.ParseInt(ctx.Param("season"), 0, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error":   err.Error(),
			"message": "Unable to parse season",
		})
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error":   err.Error(),
			"message": "Unable to parse team id",
		})
		return
	}

	espnClient := ctx.MustGet("espn").(espnclient.EspnClient)
	team, err := espnClient.GetTeam(uint(season), int(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"message": "Unable to get team",
			"error":   err.Error(),
		})
		return
	}

	if team.Id == 0 {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"message": "Unable to find team by id",
			"teamId":  id,
		})
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"season":  season,
		"team":    team,
	})
}
