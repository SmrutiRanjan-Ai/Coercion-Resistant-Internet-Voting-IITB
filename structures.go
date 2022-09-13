package main

import (
	"crypto/ecdsa"
)

const (
	host     = "localhost"
	port     = 5432
	user1    = "postgres"
	password = "213050077"
	dbname   = "MyDB"
)

var voters []voter
var publicKey *ecdsa.PublicKey
var privateKey *ecdsa.PrivateKey

type candidate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type voter struct {
	ID                string `json:"id"`
	VoterID           string
	BallotIndex       int
	ConfirmationToken string
}
type ballot struct {
	Ballot_1 ballot1 `json:"ballot1"`
	Ballot_2 ballot2 `json:"ballot2"`
	Serial   string  `json:"serial"`
	Voterid  string  `json:"voterid"`
}
type candidate1 struct {
	option        string `json:"candidate_serial"`
	candidateName string `json:"candidate_serial"`
}

type candidate2 struct {
	candidateSerial string `json:"candidate_serial"`
}

type choice1 struct {
	Index  int  `json:"index"`
	Choice bool `json:"choice"`
}

type choice2 struct {
	Index  int    `json:"index"`
	Option string `json:"option"`
	Choice bool   `json:"choice"`
}

type ballot1 struct {
	Candidate1List []candidate1 `json:"candidate1List"`
	Options1       []choice1    `json:"options1"`
	Serial         string       `json:"serial"`
	Nonce          int64        `json:"nonce"`
	Hash           string       `json:"hash"`
	Pk             string       `json:"pk"`
}

type ballot2 struct {
	Candidate2List []candidate2 `json:"candidate2List"`
	Options        []choice2    `json:"options2"`
	Serial         string       `json:"serial"`
	Nonce          int64        `json:"nonce"`
	Hash           string       `json:"hash"`
	Pk             string       `json:"pk"`
}

type vote struct {
	ID        string    `json:"id"`
	Candidate candidate `json:"candidate"`
}

var ballotVoterPair = make(map[string]string)
var ballotListTimestamp = make(map[string]int64)
var pairingList = make(map[string]map[int]string)
var cPairingList = make(map[string]map[string]string)
var votes []vote
var candidatesNum = 5
var candidates = make(map[string]string)
var tallyballot1 = make(map[string]ballot1)
var tallyballot2 = make(map[string]ballot2)
var tallyballots = make(map[string]bool)

type serialTimestamp struct {
	serial    string
	timestamp int64
}

var tallyID = make(map[string][]serialTimestamp)
