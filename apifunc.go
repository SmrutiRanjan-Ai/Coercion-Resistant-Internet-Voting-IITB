package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func getCandidates(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, candidateList)
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
	/*psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user1, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	err = db.Ping()
	CheckError(err)
	insertDynStmt := `insert into voters ("id", "votertoken","confirmationtoken", "ballotindex") values($1, $2, $3, $4)`
	_, er := db.Exec(insertDynStmt, user.ID, user.VoterID, user.ConfirmationToken, user.BallotIndex)
	CheckError(er)*/
	fmt.Println(voters)
	voterBallot := createBallot(&user)
	c.IndentedJSON(http.StatusCreated, *voterBallot)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getBallot(voterID string, candidates []candidate) {
	fmt.Println("hello")
}

func createBallot(user *voter) *ballot {
	serial := uuidGen()
	now := time.Now()
	ballot1Temp := createBallot1(serial)
	ballot2Temp := createBallot2(serial)
	var voterBallot *ballot
	*voterBallot = ballot{Ballot_1: *ballot1Temp, Ballot_2: *ballot2Temp, Serial: serial, Voterid: user.VoterID}
	ballotVoterPair[serial] = user.VoterID
	ballotListTimestamp[serial] = now.Unix()
	createPairing(serial)
	return voterBallot

}

func createBallot1(serial string) *ballot1 {
	var b1 *ballot1
	c1List, ch1 := createCandidate1(serial)
	*b1 = ballot1{Candidate1List: c1List, Options1: ch1, Serial: serial, Nonce: 0}
	return b1
}

func createBallot2(serial string) *ballot2 {
	var b2 *ballot2
	c2List, ch2 := createCandidate2(serial)
	*b2 = ballot2{Candidate2List: c2List, Options: ch2, Serial: serial}
	return b2
}

func createPairing(serial string) {
	var set []string
	var pair = make(map[int]string)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < candidatesNum; i++ {
		set = append(set, string('A'+i))
	}
	for i := range set {
		j := rand.Intn(i + 1)
		set[i], set[j] = set[j], set[i]
	}
	for idx, val := range set {
		pair[idx+1] = val
	}
	pairingList[serial] = pair
	for i := range set {
		j := rand.Intn(i + 1)
		set[i], set[j] = set[j], set[i]
	}
	i := 0
	var cPair = make(map[string]string)
	for key, _ := range candidates {
		cPair[key] = set[i]
		i++
	}
	cPairingList[serial] = cPair
}

func createCandidate1(serial string) ([]candidate1, []choice1) {
	var c1List []candidate1
	var ch1 []choice1
	cPair := cPairingList[serial]
	idx := 1
	for key, val := range candidates {
		c1List = append(c1List, candidate1{option: cPair[key], candidateName: val})
		ch1 = append(ch1, choice1{Index: idx, Choice: false})
		idx++
	}
	return c1List, ch1
}

func createCandidate2(serial string) ([]candidate2, []choice2) {
	var c2List []candidate2
	var ch2 []choice2
	pair := pairingList[serial]
	idx := 1
	for key, _ := range candidates {
		c2List = append(c2List, candidate2{candidateSerial: key})
		ch2 = append(ch2, choice2{Index: idx, Option: pair[idx], Choice: false})
		idx++
	}
	return c2List, ch2
}

func postBallot1(c *gin.Context) {
	var b1 *ballot1

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&b1); err != nil {
		return
	}
	serial := b1.Serial
	tallyballot1[serial] = *b1
}

func postBallot2(c *gin.Context) {
	var b2 *ballot2

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&b2); err != nil {
		return
	}
	serial := b2.Serial
	tallyballot2[serial] = *b2
	tallyballots[serial] = true
	ts := time.Now().Unix()
	sT := serialTimestamp{serial: serial, timestamp: ts}
	_, status := tallyID[b2.Pk]
	if status {

	}
}

func createCandidates(num int, name string) {
	for i := 0; i < num; i++ {
		concatenated := fmt.Sprintf("%s %d", name, i+1)
		serial := uuidGen()
		candidates[serial] = concatenated
	}

}

func verifyBallot() {
	/*lengthBallotsReceived:=len(tallyballots)
	lengthBallotsIssued:=len(ballotListTimestamp)*/

}
