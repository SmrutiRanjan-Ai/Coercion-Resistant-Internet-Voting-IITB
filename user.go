package main

import (
	"github.com/gin-gonic/gin"
)

func user() {

	router := gin.Default()

	router.GET("/candidates", getCandidates)

	router.GET("/vote", getVotes)
	router.POST("/vote", postVoterID)

	err := router.Run("localhost:8081")
	if err != nil {
		return
	}
}
