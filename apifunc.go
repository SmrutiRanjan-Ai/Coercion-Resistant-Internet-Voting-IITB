package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getCandidates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, candidates)
}
func getVotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, voters)
}

func postVoterID(c *gin.Context) {
	var user voter

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&user); err != nil {
		return
	}
	var found = false
	for i, a := range voters {
		if a.ID == user.ID {
			voters[i].BallotIndex += 1
			found = true
			fmt.Println("hello", a)
			fmt.Println("helloa", voters[i])
			user = voters[i]
			break

		}
	}
	if found {
		fmt.Println("found")

	} else {
		user.VoterID = vid(user.ID)
		user.BallotIndex = 0
		user.ConfirmationToken = "Confirmed"
		voters = append(voters, user)
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user1, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	insertDynStmt := `insert into voters ("id", "votertoken","confirmationtoken", "ballotindex") values($1, $2, $3, $4)`
	_, er := db.Exec(insertDynStmt, user.ID, user.VoterID, user.ConfirmationToken, user.BallotIndex)
	CheckError(er)
	fmt.Println(voters)
	c.IndentedJSON(http.StatusCreated, user)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getBallot(voterID string, candidates []candidate) {
	
}
