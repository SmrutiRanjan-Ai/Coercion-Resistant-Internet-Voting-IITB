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
	c.IndentedJSON(http.StatusCreated, voterBallot)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func getBallot(voterID string, candidates []candidate) {
	fmt.Println("hello")
}

func createBallot(user *voter) ballot {
	serial := uuidGen()
	now := time.Now()
	createPairing(serial)
	ballot1Temp := createBallot1(serial, user.VoterID)
	ballot2Temp := createBallot2(serial, user.VoterID)
	var voterBallot ballot
	voterBallot = ballot{Ballot_1: ballot1Temp, Ballot_2: ballot2Temp, Serial: serial, Voterid: user.VoterID}
	ballotVoterPair[serial] = user.VoterID
	ballotListTimestamp[serial] = now.Unix()

	return voterBallot

}

func createBallot1(serial string, pk string) ballot1 {
	var b1 ballot1
	c1List, ch1 := createCandidate1(serial)
	b1 = ballot1{Candidate1List: c1List, Options1: ch1, Serial: serial, Nonce: 0, Pk: pk}
	return b1
}

func createBallot2(serial string, pk string) ballot2 {
	var b2 ballot2
	c2List, ch2 := createCandidate2(serial)
	b2 = ballot2{Candidate2List: c2List, Options: ch2, Serial: serial, Pk: pk}
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
		c1List = append(c1List, candidate1{Option: cPair[key], CandidateSerial: key, CandidateName: val})
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
	for key, val := range candidates {
		c2List = append(c2List, candidate2{CandidateSerial: key, CandidateName: val})
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
	tallyballots[serial] = true
	ts := time.Now().Unix()
	sT := serialTimestamp{serial: serial, timestamp: ts}
	_, status := tallyID[b1.Pk]
	if status {
		tallyID[b1.Pk] = append(tallyID[b1.Pk], sT)
	} else {
		tallyID[b1.Pk] = []serialTimestamp{sT}
	}

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
		tallyID[b2.Pk] = append(tallyID[b2.Pk], sT)
	} else {
		tallyID[b2.Pk] = []serialTimestamp{sT}
	}

}

func createCandidates(num int, name string) {
	for i := 0; i < num; i++ {
		concatenated := fmt.Sprintf("%s %d", name, i+1)
		serial := uuidGen()
		candidates[serial] = concatenated
		c := candidate{ID: serial, Name: concatenated}
		candidateList = append(candidateList, c)
	}

}

func verifyBallot() {
	lengthBallotsReceived := len(tallyballots)
	lengthBallotsIssued := len(ballotListTimestamp)
	if lengthBallotsReceived <= lengthBallotsIssued {
		fmt.Println("Ballots in control")
	} else {
		fmt.Println("Fake Ballots introduced")
	}

}

func calculateVotes(c *gin.Context) {
	verifyBallot()
	var choiceIndex int
	var finalCandidate string
	for key, value := range tallyballot1 {
		serial := key
		_, status := tallyID[value.Pk]
		if status {
			_, status := calculatedVotes[key]
			if status {
				continue
			} else {
				calculatedVotes[key] = true
				for _, val := range value.Options1 {
					if val.Choice == true {
						choiceIndex = val.Index
						break
					}
				}
				stringChoice := pairingList[serial][choiceIndex]

				for _, candidate := range value.Candidate1List {
					if stringChoice == candidate.Option {
						finalCandidate = candidate.CandidateSerial
					}
				}
				_, status := candidateVotes[finalCandidate]
				if status {
					candidateVotes[finalCandidate]++
				} else {
					candidateVotes[finalCandidate] = 1
				}

			}
		}
	}
	for key, value := range tallyballot2 {
		_, status := tallyID[value.Pk]
		if status {
			_, status := calculatedVotes[key]
			if status {
				continue
			} else {
				for _, val := range value.Options {
					if val.Choice {
						pair := cPairingList[key]
						for cand, opt := range pair {
							if opt == val.Option {
								finalCandidate = cand
								break
							}
						}
					}
					break
				}
				_, status := candidateVotes[finalCandidate]
				if status {
					candidateVotes[finalCandidate]++
				} else {
					candidateVotes[finalCandidate] = 1
				}

			}
		}

	}
	var results []candidateVotesResults
	var result candidateVotesResults
	for i, j := range candidateVotes {
		result.CandidateSerial = i
		result.CandidateName = candidates[i]
		result.CandidateCount = j
		fmt.Println(i, candidates[i], j)
		results = append(results, result)
	}
	c.IndentedJSON(http.StatusOK, results)

}

func postHedgehog(c *gin.Context) {
	var hedgehog *hedgehog
	if err := c.BindJSON(&hedgehog); err != nil {
		return
	}
	hedgehogNonceList[hedgehog.VoterPk] = *hedgehog
	c.IndentedJSON(http.StatusOK, hedgehog)
}

func postSecretCode(c *gin.Context) {
	var nonce *secretCode
	if err := c.BindJSON(&nonce); err != nil {
		return
	}
	nonceList[nonce.Nonce] = true
	c.IndentedJSON(http.StatusOK, nonce)
}

func querySecretCode(c *gin.Context) {
	var nonce *secretCode
	if err := c.BindJSON(&nonce); err != nil {
		return
	}
	_, status := nonceList[nonce.Nonce]
	if status {
		c.IndentedJSON(http.StatusOK, true)
	} else {
		c.IndentedJSON(http.StatusOK, false)
	}
}

func postNullVote(c *gin.Context) {
	var h *hedgehog
	if err := c.BindJSON(&h); err != nil {
		return
	}
	_, status := hedgehogNonceList[h.VoterPk]
	if status {
		delete(tallyID, h.VoterPk)
		c.IndentedJSON(http.StatusOK, *h)
	} else {
		c.IndentedJSON(http.StatusOK, false)
	}
}

func getSecretCodes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nonceList)
}
