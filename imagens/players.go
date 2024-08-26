package main 

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPlayers(context *gin.Context) {
	var players = []string{
		"Graham Chapman",
		"John Cleese",
		"Michael Palin",
	}

	context.IndentedJSON(http.StatusOK, players)
}