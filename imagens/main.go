package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var playerRepository PlayerRepository = PlayerRepository{[]Player{}}

func getData(context *gin.Context) {
	var data = []string{
		"ABC",
		"DEF",
	}

	context.IndentedJSON(http.StatusOK, data)
}

func sortList(context *gin.Context) {
	var list []int

	if err := context.BindJSON(&list); err != nil {
		return
	}

	context.IndentedJSON(http.StatusAccepted, MergeSort(list))
}

func quickSort(context *gin.Context) {
	var list []int

	if err := context.BindJSON(&list); err != nil {
		return
	}

	context.IndentedJSON(http.StatusAccepted, QuickSort(list, 0, len(list)-1))
}

func getPlayers(context *gin.Context) {

	context.IndentedJSON(http.StatusAccepted, playerRepository.findAll())
}

func savePlayer(context *gin.Context) {
	var player Player

	if err := context.BindJSON(&player); err != nil {
		fmt.Println("Error parsing request for saving player", context)
		return
	}

	context.IndentedJSON(http.StatusAccepted, playerRepository.save(player))
}

func main() {
	router := gin.Default()
	router.GET("/", getData)
	router.POST("/sort", sortList)
	router.POST("/quick-sort", quickSort)

	router.GET("/players", getPlayers)
	router.POST("/players", savePlayer)

	router.Run(":8181")
}
