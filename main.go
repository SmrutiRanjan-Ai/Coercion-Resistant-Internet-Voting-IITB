package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

import "github.com/gin-gonic/gin"

func main() {
	go user()
	publicKey, privateKey = publicInit()
	fmt.Println(publicKey, privateKey)

	fmt.Println("Connected!")

	router := gin.Default()

	router.GET("/candidates", getCandidates)
	router.GET("/vote", getVotes)
	router.POST("/vote", postVoterID)
	router.GET("/getballot", getBallot)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
