package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getData(context *gin.Context) {
	var data = []string{
		"ABC",
		"DEF",
	}

	context.IndentedJSON(http.StatusOK, data)
}

func main() {
	router := gin.Default()
	router.GET("/", getData)

	router.Run(":8181")
}
