package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTeams(ctx *gin.Context) {
	season, err := strconv.ParseInt(ctx.Param("season"), 0, 8)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Unable to parse season. Must be an int.")
		return
	}

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"season":  season,
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

	ctx.JSON(http.StatusOK, &gin.H{
		"message": "Ok",
		"season":  season,
		"team":    id,
	})
}
