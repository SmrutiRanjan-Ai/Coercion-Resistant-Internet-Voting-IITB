package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

import "github.com/gin-gonic/gin"

func main() {

	createCandidates(candidatesNum, "Narendra Modi")
	publicKey, privateKey = publicInit()
	fmt.Println(publicKey, privateKey)
	fmt.Println("Connected!")
	router := gin.Default()
	router.GET("/candidates", getCandidates)
	router.GET("/vote", getVotes)
	router.GET("/tallyvotes", calculateVotes)
	router.POST("/vote", postVoterID)
	router.POST("/hedgehog", postHedgehog)
	router.POST("/secretcode", postSecretCode)
	router.GET("/secretcodes", getSecretCodes)
	router.POST("/querysecretcode", querySecretCode)
	router.POST("/nullvote", postNullVote)
	router.POST("/ballot1", postBallot1)
	router.POST("/ballot2", postBallot2)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
